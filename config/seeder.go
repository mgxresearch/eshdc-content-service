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
			ImageURL: "/uploads/hero1.png",
			Tag:      "Digital Land Allocation",
			Title:    "68+ Massive Estates: Digitizing Enugu",
			Description: "The ESHDC digital transformation is now live with 68+ premium estates and 34,000+ digitized plots across the state.",
			Order:    1,
		},
		{
			ImageURL: "/uploads/hero2.png",
			Tag:      "Premium Living",
			Title:    "Centenary City: The Future of Enugu",
			Description: "Experience smart city living in Centenary City. Secure your plot through our 100% verified digital portal.",
			Order:    2,
		},
		{
			ImageURL: "/uploads/hero3.png",
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
			Title:     "ESHDC Digitization Process: Step-by-Step",
			Slug:      "eshdc-digitization-process",
			Summary:   "Learn how the Enugu State Housing Development Corporation is digitizing every property record for maximum security.",
			Content:   "<h1>Digitizing Enugu's Future</h1><p>The ESHDC digitization process involves scanning historical files, indexing financial metadata (receipts and revenue codes), and securing them in a blockchain-ready digital vault. This ensures that no property record can be tampered with, and every transaction is fully traceable.</p><p>Staff members use high-precision scanners to capture every page of the property file, which is then verified by the registry department before being committed to the ledger.</p>",
			ImageURL:  "/uploads/news/digitization.png",
			Category:  "Digital Land",
			IsVisible: true,
			IsFeatured: true,
			IsMock:    true,
		},
		{
			Title:     "How to Use the ESHDC Portal: A Guide for Citizens",
			Slug:      "citizen-portal-guide",
			Summary:   "A comprehensive guide on how citizens can register, search for estates, and apply for land allocation online.",
			Content:   "<h1>Empowering Citizens Digitally</h1><p>The new ESHDC Public Portal allows citizens to create a verified profile using their NIN. Once logged in, you can browse over 68 estates, view available plots on our interactive GIS map, and initiate applications for allocation without visiting our physical office.</p><p>Payments are integrated via Remita and secure bank transfers, with digital receipts issued immediately upon verification.</p>",
			ImageURL:  "/uploads/news/portal_guide.png",
			Category:  "Housing",
			IsVisible: true,
			IsMock:    true,
		},
		{
			Title:     "Land Transfer: What You Should Know",
			Slug:      "land-transfer-guide",
			Summary:   "Understanding the legal and administrative steps required for a successful land transfer in Enugu State.",
			Content:   "<h1>Seamless Property Transfers</h1><p>Transferring property in Enugu has been simplified. The new process requires both the assignor and assignee to have verified digital accounts. The assignor initiates the transfer from their property dashboard, uploads the deed of assignment, and the system automatically routes it for administrative review.</p><p>Once approved, the property ledger is updated in real-time, reflecting the new owner and maintaining a complete history of the asset.</p>",
			ImageURL:  "/uploads/news/land_transfer.png",
			Category:  "Legal",
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

	// Seed Memos
	memos := []models.Memo{
		{
			Serial: "ESHDC/MEMO/2026/001",
			Subject: "Digitization Phase 1 Completion",
			Content: "<h1>Phase 1 Complete</h1><p>We have successfully digitized all files for Independence Layout. Registry staff should now begin verification of financial metadata.</p>",
			SenderName: "Admin Office",
			SenderEmail: "admin@eshdc.gov.ng",
			Recipients: "staff@eshdc.gov.ng",
			RecipientNames: "All Staff",
			Category: "General Administration",
			Priority: "High",
			Status: "Official",
		},
		{
			Serial: "ESHDC/MEMO/2026/002",
			Subject: "New Security Protocol for Land Transfers",
			Content: "<h1>Security Update</h1><p>From Monday, all land transfers must be initiated digitally. Physical submissions will no longer be accepted without a portal tracking ID.</p>",
			SenderName: "Legal Department",
			SenderEmail: "legal@eshdc.gov.ng",
			Recipients: "registry@eshdc.gov.ng",
			RecipientNames: "Registry Officers",
			Category: "Legal & Secretarial",
			Priority: "Urgent",
			Status: "Official",
		},
	}

	for _, m := range memos {
		var existing models.Memo
		if err := DB.Where("serial = ?", m.Serial).First(&existing).Error; err != nil {
			DB.Create(&m)
		}
	}

	fmt.Println("✅ Settings, Metadata, Jobs & Memos seeded")
}
