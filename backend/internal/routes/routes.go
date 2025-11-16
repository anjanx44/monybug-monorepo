package routes

import (
	"monybug-backend/internal/database"
	"monybug-backend/internal/handlers"
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
		// Categories routes
		api.GET("/categories", handlers.GetCategories)
		api.POST("/categories", handlers.CreateCategory)

		// Transactions routes
		api.GET("/transactions", handlers.GetTransactions)
		api.POST("/transactions", handlers.CreateTransaction)
	}
}
