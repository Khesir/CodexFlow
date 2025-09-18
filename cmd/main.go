package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var embeddedFrontend embed.FS

func main() {
	r := gin.Default()

	// API routes
	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
	}

	// Serve frontend static files
	distFS, _ := fs.Sub(embeddedFrontend, "frontend/dist")
	r.StaticFS("/assets", http.FS(distFS)) // serve static assets

	// SPA fallback for React Router
	r.NoRoute(func(c *gin.Context) {
		c.FileFromFS("index.html", http.FS(distFS))
	})

	log.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}
