package handlers

import (
	"fmt"
	"monybug-backend/internal/models"
	"monybug-backend/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTransactions(c *gin.Context) {
	transactions, err := repositories.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if transaction.CategoryID == 0 || transaction.Amount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category_id and amount are required"})
		return
	}

	if err := repositories.CreateTransaction(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var transactionID uint
	if _, err := fmt.Sscanf(id, "%d", &transactionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := repositories.UpdateTransaction(transactionID, &transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "transaction updated"})
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	var transactionID uint
	if _, err := fmt.Sscanf(id, "%d", &transactionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := repositories.DeleteTransaction(transactionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "transaction deleted"})
}
