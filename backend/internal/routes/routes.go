package routes

import (
	"monybug-backend/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		dbStatus := "ok"
		if err := database.CheckHealth(); err != nil {
			dbStatus = "error: " + err.Error()
		}
		c.JSON(http.StatusOK, gin.H{
			"status":   "ok",
			"message":  "MonyBug API is running",
			"database": dbStatus,
		})
	})

	// API v1 routes
	api := r.Group("/api/v1")
	{
		// Categories routes (placeholder)
		api.GET("/categories", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Categories endpoint"})
		})

		// Transactions routes (placeholder)
		api.GET("/transactions", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Transactions endpoint"})
		})

		api.POST("/transactions", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Add transaction endpoint"})
		})
	}
}
