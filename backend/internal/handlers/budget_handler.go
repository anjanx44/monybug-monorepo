package handlers

import (
	"fmt"
	"monybug-backend/internal/models"
	"monybug-backend/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBudgets(c *gin.Context) {
	budgets, err := repositories.GetAllBudgets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, budgets)
}

func CreateBudget(c *gin.Context) {
	var budget models.Budget
	if err := c.ShouldBindJSON(&budget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if budget.CategoryID == 0 || budget.MonthlyLimit == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category_id and monthly_limit are required"})
		return
	}

	if err := repositories.CreateBudget(&budget); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, budget)
}

func UpdateBudget(c *gin.Context) {
	id := c.Param("id")
	var budget models.Budget
	if err := c.ShouldBindJSON(&budget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var budgetID uint
	if _, err := fmt.Sscanf(id, "%d", &budgetID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := repositories.UpdateBudget(budgetID, &budget); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "budget updated"})
}

func DeleteBudget(c *gin.Context) {
	id := c.Param("id")
	var budgetID uint
	if _, err := fmt.Sscanf(id, "%d", &budgetID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := repositories.DeleteBudget(budgetID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "budget deleted"})
}
