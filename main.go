package main

import (
	"video-collector/server/database"
	"video-collector/server/handlers"
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed all:web/dist
var embeddedFiles embed.FS

func main() {
	// 1. Initialize Database
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 2. Seed initial data
	database.SeedData(db)

	// 3. Initialize API handlers with the database connection
	h := &handlers.Handler{DB: db}

	// 4. Set up Gin for API routes
	apiRouter := gin.Default()
	apiRouter.RedirectTrailingSlash = false
	api := apiRouter.Group("/api")
	{
		api.GET("/products", h.GetProducts)
		api.POST("/products", h.CreateProduct)
		api.PUT("/products/:id", h.UpdateProduct)
		api.PUT("/products/:id/status", h.UpdateProductStatus)
		api.DELETE("/products/:id", h.DeleteProduct)

		api.GET("/videos", h.GetVideos)
		api.POST("/videos", h.CreateVideo)
		api.PUT("/videos/:id", h.UpdateVideo)
		api.DELETE("/videos/:id", h.DeleteVideo)
	}

	// 5. Set up frontend file server
	subFS, err := fs.Sub(embeddedFiles, "web/dist")
	if err != nil {
		log.Fatal(err)
	}

	// 6. Set up main router to delegate requests
	mux := http.NewServeMux()
	mux.Handle("/api/", apiRouter)
	mux.Handle("/", http.FileServer(http.FS(subFS)))

	// 7. Start the server
	port := "35781"
	log.Printf("Server starting on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
