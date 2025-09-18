package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Start starts the Gin server on addr. If distEmbed is non-nil it will try
// to serve the embedded frontend from "frontend/dist" (use nil in pure-server mode).
func Start(addr string, distEmbed embed.FS) error {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
	}
	// If an embedded FS is provided, try to serve the frontend
	if distEmbed != (embed.FS{}) {
		// Try to get the dist subdir; if that fails we just log and continue
		if distFS, err := fs.Sub(distEmbed, "frontend/dist"); err == nil {
			r.StaticFS("/assets", http.FS(distFS))
			r.NoRoute(func(c *gin.Context) {
				c.FileFromFS("index.html", http.FS(distFS))
			})
		} else {
			log.Println("embedded frontend not found (ok for dev):", err)
		}
	}
	return r.Run(addr)
}
