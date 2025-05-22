package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ylmzemre/FileManager-API/config"
	"github.com/ylmzemre/FileManager-API/handlers"
	"github.com/ylmzemre/FileManager-API/middleware"
	"github.com/ylmzemre/FileManager-API/models"
)

func main() {
	cfg := config.LoadConfig()
	cfg.DB.AutoMigrate(&models.User{}, &models.File{})

	r := gin.Default()
	r.Static("/", "./static") // index.html
	api := r.Group("/api")

	// auth
	api.POST("/auth/register", handlers.Register(cfg.DB, cfg.JWTSecret))
	api.POST("/auth/login", handlers.Login(cfg.DB, cfg.JWTSecret))

	// protected
	sec := api.Group("/")
	sec.Use(middleware.JWT(cfg.JWTSecret))
	sec.GET("/files", handlers.List(cfg.DB))
	sec.POST("/files", handlers.Upload(cfg.DB, cfg.UploadPath))
	sec.DELETE("/files/:id", handlers.Delete(cfg.DB))

	r.Run(":8080")
}
