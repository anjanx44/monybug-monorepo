package handlers

import (
	"monybug-backend/internal/models"
	"monybug-backend/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInsights(c *gin.Context) {
	insights, err := repositories.GetAllInsights()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, insights)
}

func CreateInsight(c *gin.Context) {
	var insight models.ExpenseInsight
	if err := c.ShouldBindJSON(&insight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if insight.Type == "" || insight.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insight_type and message are required"})
		return
	}

	if err := repositories.CreateInsight(&insight); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, insight)
}
