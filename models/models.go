package models

import (
	"time"
	"gorm.io/gorm"
)

type News struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"not null" json:"title"`
	Slug      string         `gorm:"uniqueIndex;not null" json:"slug"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	Summary   string         `json:"summary"`
	ImageURL   string         `json:"image_url"`
	Category   string         `json:"category"`
	IsFeatured bool           `gorm:"default:false" json:"is_featured"`
	IsVisible  bool           `gorm:"default:true" json:"is_visible"`
	IsMock    bool           `gorm:"default:false" json:"is_mock"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type HeroSlide struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ImageURL  string    `json:"image_url"`
	Tag       string    `json:"tag"`
	Title     string    `json:"title"`
	Description string  `json:"description"`
	Order     int       `json:"order"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PageContent struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	PageName  string         `gorm:"uniqueIndex;not null" json:"page_name"` // e.g. "home_hero", "about_us"
	Key       string         `gorm:"not null" json:"key"`       // e.g. "title", "subtitle", "body"
	Value     string         `gorm:"type:text" json:"value"`
	IsMock    bool           `gorm:"default:false" json:"is_mock"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type Setting struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Key   string `gorm:"uniqueIndex;not null" json:"key"`
	Value string `gorm:"type:jsonb" json:"value"` // JSON array or object
}

type Job struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Department  string         `json:"department"`
	Location    string         `json:"location"`
	Type        string         `json:"type"` // Full-time, Contract, etc.
	Description string         `gorm:"type:text" json:"description"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
}

type ContactMessage struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"not null" json:"email"`
	Subject   string    `json:"subject"`
	Message   string    `gorm:"type:text;not null" json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type Media struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FileName  string    `json:"file_name"`
	FilePath  string    `json:"file_path"`
	FileType  string    `json:"file_type"`
	Size      int64     `json:"size"`
	CreatedAt time.Time `json:"created_at"`
}

type Memo struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Serial      string         `gorm:"uniqueIndex;not null" json:"serial"`
	Subject     string         `gorm:"not null" json:"subject"`
	Content     string         `gorm:"type:text;not null" json:"content"`
	SenderID    uint           `json:"sender_id"`
	SenderName  string         `json:"sender_name"`
	SenderEmail string         `json:"sender_email"`
	Recipients      string         `json:"recipients"` // Comma separated emails
	RecipientNames  string         `json:"recipient_names"` // Comma separated names
	Category        string         `json:"category"`
	Type            string         `gorm:"default:'internal'" json:"type"` // internal, external
	Status          string         `gorm:"default:'Official'" json:"status"`
	RefFile     string         `json:"ref_file"`
	Signature   string         `json:"signature"`
	Attachments string         `json:"attachments"` // JSON array
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
