package config

import (
	"github.com/ylmzemre/FileManager-API/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type AppCfg struct {
	DB         *gorm.DB
	JWTSecret  []byte
	UploadPath string
}

func LoadConfig() *AppCfg {
	dsn := env("DB_DSN", "host=localhost user=postgres password=postgres dbname=file_db port=5432 sslmode=disable")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("db error: %v", err)
	}

	upload := env("UPLOAD_DIR", "./uploads")
	if err := os.MkdirAll(upload, 0755); err != nil {
		log.Fatalf("upload dir: %v", err)
	}

	return &AppCfg{
		DB:         db,
		JWTSecret:  []byte(env("JWT_SECRET", "CHANGE_ME")),
		UploadPath: upload,
	}
}

func env(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
