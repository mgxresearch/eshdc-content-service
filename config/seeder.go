package config

import (
	"fmt"
	"github.com/eshdc/content-service/models"
)

func SeedDatabase() {
	fmt.Println("🌱 Seeding Content Database...")

	// 0. Hero Slides
	slides := []models.HeroSlide{
		{
			ImageURL: "/images/hero1.png",
			Tag:      "Digital Land Allocation",
			Title:    "68+ Massive Estates: Digitizing Enugu",
			Description: "The ESHDC digital transformation is now live with 68+ premium estates and 34,000+ digitized plots across the state.",
			Order:    1,
		},
		{
			ImageURL: "/images/hero2.png",
			Tag:      "Premium Living",
			Title:    "Centenary City: The Future of Enugu",
			Description: "Experience smart city living in Centenary City. Secure your plot through our 100% verified digital portal.",
			Order:    2,
		},
		{
			ImageURL: "/images/hero3.png",
			Tag:      "Estate Vault",
			Title:    "Registry-Safe: Digital Property Folders",
			Description: "Every asset is protected in a centralized digital vault. Transparent, secure, and fully traceable.",
			Order:    3,
		},
	}

	for _, s := range slides {
		var existing models.HeroSlide
		if err := DB.Where("title = ?", s.Title).First(&existing).Error; err != nil {
			DB.Create(&s)
		}
	}

	// 1. Initial Page Content
	pages := []models.PageContent{
		{PageName: "home_hero", Key: "title", Value: "Digitizing Housing in Enugu State", IsMock: true},
		{PageName: "home_hero", Key: "subtitle", Value: "Experience a seamless, transparent, and secure estate management system powered by ESHDC.", IsMock: true},
		{PageName: "home_hero", Key: "cta_text", Value: "Explore Estates", IsMock: true},
		
		{PageName: "about_us", Key: "mission", Value: "To provide affordable and quality housing solutions for all residents of Enugu State.", IsMock: true},
		{PageName: "about_us", Key: "vision", Value: "To be the leading housing development corporation in Nigeria through innovation and digitization.", IsMock: true},
		{PageName: "legal", Key: "terms_of_service", Value: "<h1>Terms of Service</h1><p>By using our digital portal, you agree to comply with all ESHDC regulations regarding land allocation and property management...</p>", IsMock: true},
		{PageName: "legal", Key: "privacy_policy", Value: "<h1>Privacy Policy</h1><p>Your data is secured using banking-grade encryption. We do not share your KYC information with third parties...</p>", IsMock: true},
		{PageName: "legal", Key: "compliance", Value: "<h1>Compliance & Regulatory</h1><p>All property allocations are subject to the Enugu State Land Use Act and ESHDC digitization guidelines 2026...</p>", IsMock: true},
		{PageName: "contact_info", Key: "phone", Value: "+234 (0) 42 255 123", IsMock: true},
		{PageName: "contact_info", Key: "email", Value: "support@eshdc.en.gov.ng", IsMock: true},
		{PageName: "contact_info", Key: "address", Value: "No. 1 ESHDC Way, Independence Layout, Enugu, Nigeria", IsMock: true},
	}

	for _, p := range pages {
		var existing models.PageContent
		if err := DB.Where("page_name = ? AND key = ?", p.PageName, p.Key).First(&existing).Error; err != nil {
			DB.Create(&p)
		}
	}

	// 2. Initial News
	news := []models.News{
		{
			Title:     "ESHDC Launches New Digitization Portal",
			Slug:      "eshdc-launches-digitization-portal",
			Summary:   "The Enugu State Housing Development Corporation has officially launched its new portal for estate management.",
			Content:   "<p>Today marks a new era in housing...</p>",
			ImageURL:  "/uploads/news/hero1.png",
			Category:  "Digital Land",
			IsVisible: true,
			IsFeatured: true,
			IsMock:    true,
		},
		{
			Title:     "New Housing Project in Independence Layout",
			Slug:      "new-housing-project-independence-layout",
			Summary:   "Phase 3 of the Independence Layout project is now open for allocation to verified citizens.",
			Content:   "<p>Explore the premium plots available...</p>",
			ImageURL:  "/uploads/news/hero2.png",
			Category:  "Housing",
			IsVisible: true,
			IsMock:    true,
		},
		{
			Title:     "Sustainable Urban Development Workshop",
			Slug:      "sustainable-urban-development-workshop",
			Summary:   "ESHDC hosts a workshop on sustainable urban planning and the future of Enugu's landscape.",
			Content:   "<p>Experts gather to discuss...</p>",
			ImageURL:  "/uploads/news/hero3.png",
			Category:  "ESHDC Updates",
			IsVisible: true,
			IsMock:    true,
		},
		{
			Title:     "Digital Certificates of Occupancy Now Live",
			Slug:      "digital-c-of-o-live",
			Summary:   "Citizens can now apply for and receive their C of O digitally through the ESHDC platform.",
			Content:   "<p>Security and transparency at its peak...</p>",
			ImageURL:  "/uploads/news/enugu_premium.png",
			Category:  "Digital Land",
			IsVisible: true,
			IsMock:    true,
		},
	}

	for _, n := range news {
		var existing models.News
		if err := DB.Where("slug = ?", n.Slug).First(&existing).Error; err != nil {
			DB.Create(&n)
		} else {
			// Update existing record with new data
			n.ID = existing.ID
			DB.Save(&n)
		}
	}

	// 3. Metadata Settings
	settings := []models.Setting{
		{Key: "titles", Value: `["Mr.", "Mrs.", "Ms.", "Dr.", "Prof.", "Chief", "Engr.", "Rev.", "Hon."]`},
		{Key: "genders", Value: `["Male", "Female"]`},
		{Key: "occupations", Value: `["Civil Servant", "Private Sector", "Self Employed", "Business Owner", "Student", "Retired", "Other"]`},
		{Key: "marital_statuses", Value: `["Single", "Married", "Divorced", "Widowed"]`},
		{Key: "locations", Value: `["Enugu North", "Enugu South", "Enugu East", "Nsukka", "Awgu", "Udi", "Ezeagu", "Nkanu West", "Nkanu East"]`},
		{Key: "memo_categories", Value: `["General Administration", "Legal & Secretarial", "Finance & Accounts", "Registry & Verification", "Engineering & GIS", "Security & Compliance"]`},
		{Key: "mock_mode", Value: `"true"`},
	}

	for _, s := range settings {
		var existing models.Setting
		if err := DB.Where("key = ?", s.Key).First(&existing).Error; err != nil {
			DB.Create(&s)
		}
	}

	// Seed Jobs
	jobs := []models.Job{
		{Title: "Senior Estate Surveyor", Department: "Valuation", Location: "Enugu", Type: "Full-time", Description: "Responsible for property valuation and estate management."},
		{Title: "Front-end Developer (Next.js)", Department: "IT", Location: "Remote / Hybrid", Type: "Full-time", Description: "Building the future of digital land management."},
		{Title: "Legal Officer", Department: "Legal", Location: "Enugu", Type: "Full-time", Description: "Managing land titles and legal documentation."},
	}

	for _, j := range jobs {
		var existing models.Job
		if err := DB.Where("title = ?", j.Title).First(&existing).Error; err != nil {
			DB.Create(&j)
		}
	}

	fmt.Println("✅ Settings, Metadata & Jobs seeded")
}
