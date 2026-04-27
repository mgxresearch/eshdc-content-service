package handlers

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/eshdc/content-service/config"
	"github.com/eshdc/content-service/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadMedia(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Generate unique filename
	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension
	filePath := filepath.Join("uploads", newFileName)

	// Save to disk
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Save metadata to DB
	media := models.Media{
		FileName: file.Filename,
		FilePath: "/uploads/" + newFileName,
		FileType: file.Header.Get("Content-Type"),
		Size:     file.Size,
		CreatedAt: time.Now(),
	}
	config.DB.Create(&media)

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"url":     media.FilePath,
	})
}

func GetMedia(c *gin.Context) {
	var media []models.Media
	config.DB.Find(&media)
	c.JSON(http.StatusOK, media)
}
