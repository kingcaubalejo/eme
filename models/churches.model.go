package models

import (
	_"api/utils"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type Churches struct {
	ChurchId  			uint `json:"church_id" gorm:"primaryKey" gorm:"autoIncrement" gorm:"unique"`
	Name  				string `json:"name" gorm:"type:varchar(255)"`
	Address 			string `json:"address" gorm:"type:varchar(255)"`
	PastorId   			uint `json:"pastor_id" gorm:"index"`
	Pastors Pastors 	`gorm:"foreignKey:PastorId;references:PastorId"`
	IsDeleted 			bool `json:"-" gorm:"type:bool"`
	DateModel
}

func (c *Churches) BeforeCreate(tx *gorm.DB) (err error) {

	ctx := DB.Where("name = ?", c.Name).Find(&c)

	if ctx.RowsAffected != 0 {
		return errors.New("Church already exists.")
	}

	return nil
}

func (c *Churches) BeforeUpdate(tx *gorm.DB) (err error) {

	var churches Churches
	ctx := DB.Where("name = ?", c.Name).Find(&churches)

	if ctx.RowsAffected != 0 && c.Name != churches.Name {
		return errors.New("Churches already exists.")
	}

	return nil
}

func (c *Churches) Create() error {
	ctx := DB.Create(&c)
	return ctx.Error
}

func (c *Churches) Update(pastorId int) error {
	ctx := DB.Model(&c).Where("pastor_id = ?", pastorId).Updates(Churches{
		Name: c.Name,
		Address: c.Address,
		PastorId: c.PastorId,
		IsDeleted: c.IsDeleted,
	})

	return ctx.Error
}

func (c *Churches) Delete(churchId int) error {
	ctx := DB.Delete(&Churches{}, churchId)
	return ctx.Error
}

func (c *Churches) Get(r *http.Request) ([]Churches, int64, error) {
	var churches []Churches
	var churchesCount int64

	DB.Model(&Churches{}).Count(&churchesCount)

	ctx := DB.Scopes(paginate(r), order(r, []string{
		"church_id", 
		"name",
		"address",
		"created_at",
		"updated_at",
	})).Find(&churches)

	return churches, churchesCount, ctx.Error

}

func (c *Churches) GetInfo(churchId int) error {
	ctx := DB.Find(&c, churchId)
	if ctx.RowsAffected == 0 {
		return errors.New("Unable to find the Church.")
	}

	return ctx.Error
}

func (c *Churches) GetPastor(churchId int) error {
	ctx := DB.Model(&Churches{}).Preload("Pastors").Find(&c, churchId)
	if ctx.RowsAffected == 0 {
		return errors.New("Unable to find the church of the pastor.")
	}

	return ctx.Error
}
