package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
}

func initHTTPHandlers(g *gin.Engine, a *App) {
	// Default Error Handling

	g.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

}
