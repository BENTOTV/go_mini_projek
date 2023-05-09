package config

import (
	"gorm.io/gorm"
)

const (
	PORT = ":8000"
)

const (
	JWT_SECRET = "ThisIsSecret"
)

const (
	DB_USERNAME = "root"
	DB_PASSWORD = ""
	DB_HOST     = "localhost"
	DB_PORT     = "3306"
	DB_NAME     = "alta_mn"
)

var DB *gorm.DB
