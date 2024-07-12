package migrate

import (
	"GO-JWT-Auth/initializers"
	"GO-JWT-Auth/models"
)

func Migrate() {
	errAuth := initializers.DB.AutoMigrate(&models.Auth{})
	if errAuth != nil {
		return
	}

	errOtp := initializers.DB.AutoMigrate(&models.Otp{})
	if errOtp != nil {
		return
	}
}
