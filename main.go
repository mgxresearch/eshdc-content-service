package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eshdc/content-service/config"
	"github.com/eshdc/content-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	if os.Getenv("SEED_DB") == "true" {
		config.SeedDatabase()
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8087"
	}

	r := gin.Default()

	// Static folder for uploads
	r.Static("/uploads", "./uploads")

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "up", "service": "content-service"})
	})

	api := r.Group("/api/v1/content")
	{
		// Hero Slides
		api.GET("/slides", handlers.ListSlides)
		api.POST("/slides", handlers.CreateSlide)
		api.PUT("/slides/:id", handlers.UpdateSlide)

		// News
		api.GET("/news", handlers.ListNews)
		api.POST("/news", handlers.CreateNews)
		api.PUT("/news/:id", handlers.UpdateNews)
		api.GET("/news/:slug", handlers.GetNewsBySlug)
		
		// Page Content
		api.GET("/pages/:page_name", handlers.GetPageContent)
		api.GET("/legal/:slug", handlers.GetLegalContent)
		
		// Careers & Contact
		api.GET("/jobs", handlers.ListJobs)
		api.POST("/contact", handlers.SubmitContactMessage)

		// Settings & Metadata
		api.GET("/settings", handlers.GetSettings)
		api.GET("/settings/:key", handlers.GetSetting)
		api.POST("/settings", handlers.UpdateSetting)

		// Media
		api.POST("/media/upload", handlers.UploadMedia)
		api.GET("/media", handlers.GetMedia)

		// Memos
		api.GET("/memos", handlers.ListMemos)
		api.POST("/memos", handlers.CreateMemo)

		api.POST("/nuclear-reset", handlers.NuclearReset)
		api.POST("/seed", func(c *gin.Context) {
			config.SeedDatabase()
			c.JSON(http.StatusOK, gin.H{"message": "Content Seeded Successfully"})
		})
	}

	log.Printf("Content Service starting on port %s", port)
	r.Run(":" + port)
}
