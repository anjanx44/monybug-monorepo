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
		api.PUT("/categories/:id", handlers.UpdateCategory)
		api.DELETE("/categories/:id", handlers.DeleteCategory)

		// Transactions routes
		api.GET("/transactions", handlers.GetTransactions)
		api.POST("/transactions", handlers.CreateTransaction)
		api.PUT("/transactions/:id", handlers.UpdateTransaction)
		api.DELETE("/transactions/:id", handlers.DeleteTransaction)

		// Budgets routes
		api.GET("/budgets", handlers.GetBudgets)
		api.POST("/budgets", handlers.CreateBudget)
		api.PUT("/budgets/:id", handlers.UpdateBudget)
		api.DELETE("/budgets/:id", handlers.DeleteBudget)

		// Insights routes
		api.GET("/insights", handlers.GetInsights)
		api.POST("/insights", handlers.CreateInsight)
	}
}
