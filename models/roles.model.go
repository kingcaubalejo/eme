package models

import (
	_"api/utils"
	"errors"
	"net/http"

	"gorm.io/gorm"
	_"time"
	"fmt"
)

type Roles struct {
	RoleId  		uint `json:"-" gorm:"primaryKey" gorm:"autoIncrement" gorm:"unique"`
	RoleName  		string `json:"role_name" gorm:"type:varchar(255)"`
	RoleActions   	string `json:"role_actions" gorm:"type:varchar(255)"`
	IsActive 		bool `json:"is_active" gorm:"type:bool"`
	IsDeleted 		bool `json:"-" gorm:"type:bool"`
	DateModel
}

func (c *Roles) BeforeCreate(tx *gorm.DB) (err error) {

	ctx := DB.Where("role_name = ?", c.RoleName).Find(&c)

	if ctx.RowsAffected != 0 {
		return errors.New("Role already exists.")
	}

	return nil
}

func (c *Roles) BeforeUpdate(tx *gorm.DB) (err error) {

	var role Roles
	ctx := DB.Where("role_name = ?", c.RoleName).Find(&role)

	if ctx.RowsAffected != 0 && role.RoleName != c.RoleName {
		return errors.New("Role already exists.")
	}

	return nil
}

func (c *Roles) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Print(c.RoleName)
	if c.RoleName == "admin" {
		return errors.New("Administrator cannot be deleted.")
	}
    return nil
}

func (c *Roles) Create() error {
	ctx := DB.Create(&c)
	return ctx.Error
}

func (c *Roles) Update(roleId int) error {
	ctx := DB.Model(&c).Where("role_id = ?", roleId).Updates(Roles{
		RoleName: c.RoleName,
		RoleActions: c.RoleActions,
		IsActive: c.IsActive,
	})

	return ctx.Error
}

func (c *Roles) Delete(roleId int) error {
	ctx := DB.Delete(&Roles{}, roleId)
	return ctx.Error
}

func (c *Roles) Get(r *http.Request) ([]Roles, int64, error) {
	var roles []Roles
	var rolesCount int64

	DB.Model(&Roles{}).Count(&rolesCount)

	ctx := DB.Scopes(paginate(r), order(r, []string{
		"role_id", 
		"role_name",
		"role_actions",
		"is_active",
		"created_at",
		"updated_at",
	})).Find(&roles)

	return roles, rolesCount, ctx.Error

}

func (c *Roles) GetInfo(roleId int) error {
	ctx := DB.Find(&c, roleId)
	if ctx.RowsAffected == 0 {
		return errors.New("Unable to find role.")
	}

	return ctx.Error
}
