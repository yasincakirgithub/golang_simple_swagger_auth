package config

import "GOLANG/api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
