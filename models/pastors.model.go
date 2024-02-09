package models

import (
	_"api/utils"
	"errors"
	"net/http"

	"gorm.io/gorm"
	_"time"
	_"fmt"
)

type Pastors struct {
	PastorId  		uint `json:"pastor_id" gorm:"primaryKey" gorm:"autoIncrement" gorm:"unique"`
	LastName  		string `json:"last_name" gorm:"type:varchar(255)"`
	FirstName   	string `json:"first_name" gorm:"type:varchar(255)"`
	IsActive 		bool `json:"is_active" gorm:"type:bool"`
	IsDeleted 		bool `json:"-" gorm:"type:bool"`
	DateModel
}

func (c *Pastors) BeforeCreate(tx *gorm.DB) (err error) {

	ctx := DB.Where("first_name = ?", c.FirstName).Find(&c)

	if ctx.RowsAffected != 0 {
		return errors.New("Pastor already exists.")
	}

	return nil
}

func (c *Pastors) BeforeUpdate(tx *gorm.DB) (err error) {

	var Pastors Pastors
	ctx := DB.Where("first_name = ?", c.FirstName).Find(&Pastors)

	if ctx.RowsAffected != 0 && c.LastName != c.LastName {
		return errors.New("Pastor already exists.")
	}

	return nil
}

func (c *Pastors) Create() error {
	ctx := DB.Create(&c)
	return ctx.Error
}

func (c *Pastors) Update(pastorId int) error {
	ctx := DB.Model(&c).Where("pastor_id = ?", pastorId).Updates(Pastors{
		LastName: c.LastName,
		FirstName: c.FirstName,
		IsActive: c.IsActive,
	})

	return ctx.Error
}

func (c *Pastors) Delete(pastorId int) error {
	ctx := DB.Delete(&Pastors{}, pastorId)
	return ctx.Error
}

func (c *Pastors) Get(r *http.Request) ([]Pastors, int64, error) {
	var pastors []Pastors
	var pastorsCount int64

	DB.Model(&Pastors{}).Count(&pastorsCount)

	ctx := DB.Scopes(paginate(r), order(r, []string{
		"pastor_id", 
		"last_name",
		"first_name",
		"is_active",
		"created_at",
		"updated_at",
	})).Find(&pastors)

	return pastors, pastorsCount, ctx.Error

}

func (c *Pastors) GetInfo(pastorId int) error {
	ctx := DB.Find(&c, pastorId)
	if ctx.RowsAffected == 0 {
		return errors.New("Unable to find the Pastor.")
	}

	return ctx.Error
}
