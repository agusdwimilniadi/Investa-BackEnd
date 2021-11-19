package configs

import (
	models "investaBackend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:123edsaq1@tcp(127.0.0.1:3306)/investa_backend?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database not connected")
	}
	Migration()
}
func Migration() {
	DB.AutoMigrate(&models.BankCode{})
	DB.AutoMigrate(&models.Sektor{})
	DB.AutoMigrate(&models.UserProyek{})
	DB.AutoMigrate(&models.UserInvestasi{})
	DB.AutoMigrate(&models.ProyekMitra{})
	DB.AutoMigrate(&models.Investasi{})
}
