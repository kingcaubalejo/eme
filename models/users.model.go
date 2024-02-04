package models

import (
	"api/utils"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type Users struct {
	AccountId  uint `json:"-" gorm:"primaryKey" gorm:"autoIncrement" gorm:"unique"`
	FirstName  string `json:"first_name" gorm:"type:varchar(255)" validate:"required,min=3"`
	LastName   string `json:"last_name" gorm:"type:varchar(255)" validate:"required"`
	Email      string `json:"email" gorm:"type:varchar(255)" validate:"required,email"`
	Password   string `json:"password" gorm:"type:varchar(255)" validate:"required"`
	RoleId uint `json:"role_id" gorm:"index"`
	Roles Roles `gorm:"foreignKey:RoleId;references:RoleId"`
	IsActive bool `json:"-" gorm:"type:bool"`
	IsDeleted bool `json:"-" gorm:"type:bool"`
	DateModel
}

func (c *Users) BeforeCreate(tx *gorm.DB) (err error) {

	ctx := DB.Where("email = ?", c.Email).Find(&c)

	if ctx.RowsAffected != 0 {
		return errors.New("User already exists.")
	}

	return nil
}

func (c *Users) BeforeUpdate(tx *gorm.DB) (err error) {

	var user Users
	ctx := DB.Where("email = ?", c.Email).Find(&user)

	if ctx.RowsAffected != 0 && user.AccountId != c.AccountId {
		return errors.New("User already exists.")
	}

	return nil
}

func (c *Users) Create() error {
	ctx := DB.Create(&c)
	return ctx.Error
}

func (c *Users) Update() error {
	ctx := DB.Model(&c).Where("account_id = ?", c.AccountId).Updates(Users{
		FirstName:  c.FirstName,
		LastName:   c.LastName,
		Email:      c.Email,
	})

	return ctx.Error
}

// func (c *TblUserDelete) Delete() error {
// 	ctx := DB.Delete(&Users{}, c.IDS)
// 	if ctx.RowsAffected == 0 {
// 		return errors.New("No user deleted. User not found.")
// 	}
// 	return nil
// }

func (c *Users) Get(r *http.Request) ([]Users, int64, error) {
	var users []Users
	var userCount int64

	DB.Model(&Users{}).Count(&userCount)

	ctx := DB.Scopes(paginate(r), order(r, []string{"id", "name"})).Find(&users)

	return users, userCount, ctx.Error

}

func (c *Users) GetInfo() error {
	ctx := DB.Find(&c)
	if ctx.RowsAffected == 0 {
		return errors.New("Unable to find user.")
	}

	return ctx.Error
}

func (c *Users) FindUser() error {

	ctx := DB.Where("email=? AND password=?", c.Email, utils.MakePassword(c.Password)).Find(&c)

	if ctx.RowsAffected == 0 {
		return errors.New("Wrong username or password.")
	}
	return ctx.Error
}
