package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ylmzemre/FileManager-API/models"
	"gorm.io/gorm"
)

type authReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(db *gorm.DB, secret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req authReq
		if c.BindJSON(&req) != nil || req.Username == "" || req.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"err": "invalid"})
			return
		}
		user := models.User{Username: req.Username, Password: req.Password}
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusConflict, gin.H{"err": "user exists"})
			return
		}
		token := signJWT(user.ID, secret)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func Login(db *gorm.DB, secret []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req authReq
		if c.BindJSON(&req) != nil {
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		var u models.User
		if err := db.Where("username = ?", req.Username).First(&u).Error; err != nil {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}
		if !models.CheckPassword(u.Password, req.Password) {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": signJWT(u.ID, secret)})
	}
}

func signJWT(uid uint, secret []byte) string {
	claims := jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, _ := t.SignedString(secret)
	return s
}
