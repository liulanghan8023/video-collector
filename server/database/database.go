package database

import (
	"video-collector/server/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

func Connect() (*gorm.DB, error) {
	exePath, err := os.Executable()
	if err != nil {
		return nil, err
	}
	exeDir := filepath.Dir(exePath)
	dbPath := filepath.Join(exeDir, "douyin.db")
	log.Printf("Database path: %s", dbPath)

	log.Println("Opening database...")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to open database: %v", err)
		return nil, err
	}
	log.Println("Database opened successfully.")

	log.Println("Auto-migrating database...")
	if err := db.AutoMigrate(&models.Product{}, &models.Video{}); err != nil {
		log.Printf("Failed to auto-migrate database: %v", err)
		return nil, err
	}
	log.Println("Database auto-migrated successfully.")

	return db, nil
}

func SeedData(db *gorm.DB) {
	var productCount int64
	db.Model(&models.Product{}).Count(&productCount)
	if productCount == 0 {
		log.Println("Seeding products...")
		products := []models.Product{
			{Name: "Product A", URL: "http://example.com/product_a"},
			{Name: "Product B", URL: "http://example.com/product_b"},
		}
		for _, p := range products {
			db.Create(&p)
		}
	}

	var videoCount int64
	db.Model(&models.Video{}).Count(&videoCount)
	if videoCount == 0 {
		log.Println("Seeding videos...")
		videos := []models.Video{
			{URL: "http://example.com/video1.mp4", ProductID: 1},
			{URL: "http://example.com/video2.mp4", ProductID: 1},
			{URL: "http://example.com/video3.mp4", ProductID: 2},
		}
		for _, v := range videos {
			db.Create(&v)
		}
	}
}
