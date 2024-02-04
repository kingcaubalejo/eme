package main

import (
	"api/database"
	"api/models"
	"fmt"
	"api/utils"

	"gorm.io/gorm"
)

func initDatabase() {

	if err := database.Open(); err != nil {
		panic("Fail to connect to database")
	}
}

func main() {
	initDatabase()

	var db = database.Connect
	models.DB = database.Connect

	fmt.Println("Connected! Starting migration...")

	defer database.Close()
	TblUsers(db)
	TblRoles(db)

	fmt.Println("Migration Finished...")
}

func TblUsers(db *gorm.DB) {

	table := &models.Users{}
	columns := models.Users{
		FirstName:  "Marlon",
		LastName:   "Kamandag",
		Email:      "marlon@admin.com",
		Password:   utils.MakePassword("GodisGoodAllTheTimes"),
		IsActive: true,
	}
	if exist := db.Migrator().HasTable(table); !exist {
		utils.ErrorChecker(0, db.AutoMigrate(table))

		utils.ErrorChecker(0, columns.Create())
	}

}

func TblRoles(db *gorm.DB) {
	table := &models.Roles{}
	columns := models.Roles{
		RoleName:  "root",
		RoleActions:   "*",
		IsActive: true,
	}
	if exist := db.Migrator().HasTable(table); !exist {
		utils.ErrorChecker(0, db.AutoMigrate(table))

		utils.ErrorChecker(0, columns.Create())
	}
}