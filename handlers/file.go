package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ylmzemre/FileManager-API/models"
	"gorm.io/gorm"
)

func Upload(db *gorm.DB, uploadDir string) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetUint("uid")
		fh, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		// uzantı kontrolü
		ext := filepath.Ext(fh.Filename)
		if ext != ".pdf" && ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"err": "only pdf/png/jpg"})
			return
		}

		id := uuid.New()
		dst := filepath.Join(uploadDir, id.String()+ext)
		if err := c.SaveUploadedFile(fh, dst); err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		meta := models.File{
			ID:         id,
			UserID:     uid,
			OriginName: fh.Filename,
			Path:       dst,
			MimeType:   fh.Header.Get("Content-Type"),
			Size:       fh.Size,
		}
		db.Create(&meta)
		c.JSON(http.StatusOK, meta)
	}
}

func List(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetUint("uid")
		var files []models.File
		db.Where("user_id = ?", uid).Find(&files)
		c.JSON(http.StatusOK, files)
	}
}

func Delete(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetUint("uid")
		fileID := c.Param("id")
		if err := db.Where("id = ? AND user_id = ?", fileID, uid).Delete(&models.File{}).Error; err != nil {
			c.JSON(http.StatusNotFound, nil)
			return
		}
		c.Status(http.StatusNoContent)
	}
}
