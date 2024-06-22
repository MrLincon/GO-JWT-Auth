package migrate

import (
	"GO-JWT-Auth/initializers"
	"GO-JWT-Auth/models"
)

func Migrate() {
	err := initializers.DB.AutoMigrate(&models.Auth{})
	if err != nil {
		return
	}
}
