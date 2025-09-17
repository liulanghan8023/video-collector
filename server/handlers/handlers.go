package handlers

import "gorm.io/gorm"

// Handler is a struct that holds dependencies for API handlers, such as a database connection.
type Handler struct {
	DB *gorm.DB
}