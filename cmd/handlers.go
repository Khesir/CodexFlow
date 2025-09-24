package main

import (
	"app/internal/core/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
}

func initHTTPHandlers(g *gin.Engine, a *App) {
	// Default Error Handling

	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Public
	// g.POST("/login", func(c *gin.Context) { auth.Login(c, a.DB) })
	g.POST("/refresh", func(c *gin.Context) { auth.RefreshToken(c, a.DB) })
	g.POST("/signin", func(c *gin.Context) { auth.SignIn(c, a.DB) })

	protected := g.Group("/api", auth.Middleware())
	{
		protected.POST("/signout", func(c *gin.Context) { auth.Signout(c, a.DB) })
		protected.GET("/profile", func(c *gin.Context) {
			userID := c.GetUint("user_id") // extracted from middleware
			c.JSON(http.StatusOK, gin.H{"message": "Hello user", "user_id": userID})
		})
		protected.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "protected pong",
			})
		})
	}
}
