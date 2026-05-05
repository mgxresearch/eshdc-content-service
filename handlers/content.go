package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/eshdc/content-service/config"
	"github.com/eshdc/content-service/models"
	"github.com/gin-gonic/gin"
)

// Hero Slides Handlers
func ListSlides(c *gin.Context) {
	var slides []models.HeroSlide
	config.DB.Where("is_active = ?", true).Order("\"order\" asc").Find(&slides)
	c.JSON(http.StatusOK, slides)
}

// News Handlers
func ListNews(c *gin.Context) {
	var news []models.News
	mockMode := getMockMode()
	
	query := config.DB.Where("is_visible = ?", true)
	if !mockMode {
		query = query.Where("is_mock = ?", false)
	}
	
	query.Order("created_at desc").Find(&news)
	c.JSON(http.StatusOK, news)
}

func GetNewsBySlug(c *gin.Context) {
	var news models.News
	slug := c.Param("slug")
	mockMode := getMockMode()

	query := config.DB.Where("slug = ?", slug)
	if !mockMode {
		query = query.Where("is_mock = ?", false)
	}

	if err := query.First(&news).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}
	c.JSON(http.StatusOK, news)
}

// Page Content Handlers
func GetPageContent(c *gin.Context) {
	pageName := c.Param("page_name")
	mockMode := getMockMode()
	
	var contents []models.PageContent
	query := config.DB.Where("page_name = ?", pageName)
	if !mockMode {
		query = query.Where("is_mock = ?", false)
	}
	query.Find(&contents)
	
	result := make(map[string]string)
	for _, item := range contents {
		result[item.Key] = item.Value
	}
	
	c.JSON(http.StatusOK, result)
}

// Settings Handlers
func GetSettings(c *gin.Context) {
	var settings []models.Setting
	config.DB.Find(&settings)
	c.JSON(http.StatusOK, settings)
}

func UpdateSetting(c *gin.Context) {
	var req struct {
		Key   string `json:"key" binding:"required"`
		Value string `json:"value" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var setting models.Setting
	config.DB.Where("key = ?", req.Key).FirstOrCreate(&setting, models.Setting{Key: req.Key})
	setting.Value = req.Value
	config.DB.Save(&setting)

	c.JSON(http.StatusOK, gin.H{"message": "Setting updated"})
}

// Helper
func GetSetting(c *gin.Context) {
	key := c.Param("key")
	var setting models.Setting
	if err := config.DB.Where("key = ?", key).First(&setting).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Setting not found"})
		return
	}
	c.JSON(http.StatusOK, setting)
}

func GetLegalContent(c *gin.Context) {
	slug := c.Param("slug")
	var content models.PageContent
	if err := config.DB.Where("page_name = ? AND key = ?", "legal", slug).First(&content).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Legal content not found"})
		return
	}
	c.JSON(http.StatusOK, content)
}

func ListJobs(c *gin.Context) {
	var jobs []models.Job
	config.DB.Where("is_active = ?", true).Find(&jobs)
	c.JSON(http.StatusOK, jobs)
}

func SubmitContactMessage(c *gin.Context) {
	var req struct {
		Name    string `json:"name" binding:"required"`
		Email   string `json:"email" binding:"required"`
		Subject string `json:"subject"`
		Message string `json:"message" binding:"required"`
		Captcha string `json:"captcha" binding:"required"`
		A       int    `json:"a"`
		B       int    `json:"b"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Dynamic captcha check: if A and B are provided, check their sum.
	// Otherwise fallback to the legacy "4" check.
	if req.A != 0 || req.B != 0 {
		expected := fmt.Sprintf("%d", req.A+req.B)
		if req.Captcha != expected {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect security answer. Please try again."})
			return
		}
	} else if req.Captcha != "4" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid captcha answer (2+2=?)"})
		return
	}

	msg := models.ContactMessage{
		Name:    req.Name,
		Email:   req.Email,
		Subject: req.Subject,
		Message: req.Message,
	}

	if err := config.DB.Create(&msg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

// Memo Handlers
func ListMemos(c *gin.Context) {
	var memos []models.Memo
	recipient := c.Query("recipient")
	
	query := config.DB
	if recipient != "" {
		// Basic "contains" check for comma-separated recipients
		query = query.Where("recipients LIKE ?", "%"+recipient+"%")
	}
	
	query.Order("created_at desc").Find(&memos)
	c.JSON(http.StatusOK, memos)
}

func CreateMemo(c *gin.Context) {
	var memo models.Memo
	if err := c.ShouldBindJSON(&memo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&memo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to dispatch memo"})
		return
	}

	// Trigger Email Notification via Notification Service
	notifURL := os.Getenv("NOTIFICATION_SERVICE_URL")
	if notifURL != "" {
		recipients := strings.Split(memo.Recipients, ",")
		for _, r := range recipients {
			email := strings.TrimSpace(r)
			if strings.Contains(email, "@") {
				templateName := "internal_memo"
				if memo.Type == "external" {
					templateName = "external_memo"
				}
				payload := map[string]interface{}{
					"template":  templateName,
					"recipient": email,
					"name":      email, // Could lookup real name if needed
					"data": map[string]interface{}{
						"subject":     memo.Subject,
						"sender_name": memo.SenderName,
						"memo_serial": memo.Serial,
						"category":    memo.Category,
						"content":     memo.Content,
						"date":        time.Now().Format("02 Jan 2006"),
					},
				}
				jsonPayload, _ := json.Marshal(payload)
				http.Post(notifURL+"/api/v1/notifications/send", "application/json", bytes.NewBuffer(jsonPayload))
			}
		}
	}

	c.JSON(http.StatusOK, memo)
}

func CreateNews(c *gin.Context) {
	var news models.News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create news"})
		return
	}
	c.JSON(http.StatusCreated, news)
}

func UpdateNews(c *gin.Context) {
	id := c.Param("id")
	var news models.News
	if err := config.DB.First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update news"})
		return
	}
	c.JSON(http.StatusOK, news)
}

func CreateSlide(c *gin.Context) {
	var slide models.HeroSlide
	if err := c.ShouldBindJSON(&slide); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&slide).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create slide"})
		return
	}
	c.JSON(http.StatusCreated, slide)
}

func UpdateSlide(c *gin.Context) {
	id := c.Param("id")
	var slide models.HeroSlide
	if err := config.DB.First(&slide, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Slide not found"})
		return
	}
	if err := c.ShouldBindJSON(&slide); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&slide).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update slide"})
		return
	}
	c.JSON(http.StatusOK, slide)
}

func getMockMode() bool {
	var setting models.Setting
	if err := config.DB.Where("key = ?", "mock_mode").First(&setting).Error; err == nil {
		return setting.Value == "\"true\"" || setting.Value == "true"
	}
	return false
}



func NuclearReset(c *gin.Context) {
	// User wants to keep News and Hero Slides
	config.DB.Exec("TRUNCATE TABLE memos, contact_messages RESTART IDENTITY CASCADE")
	config.DB.Exec("DELETE FROM page_contents WHERE is_mock = true")

	// Set Production Live Flag locally
	var setting models.SystemSetting
	config.DB.Where("key = ?", "production_live").FirstOrCreate(&setting, models.SystemSetting{Key: "production_live"})
	setting.Value = "true"
	setting.UpdatedAt = time.Now()
	config.DB.Save(&setting)

	c.JSON(http.StatusOK, gin.H{"message": "Memos and Messages purged. Slides and News preserved. Production mode locked."})
}
