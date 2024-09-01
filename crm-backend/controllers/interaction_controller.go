package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "crm-backend/models"
    "crm-backend/services"
)

func CreateInteraction(c *gin.Context) {
    var interaction models.Interaction
    if err := c.ShouldBindJSON(&interaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := services.CreateInteraction(&interaction); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, interaction)
}

func GetCustomerInteractions(c *gin.Context) {
    customerId := c.Param("customerId")
    interactions, err := services.GetInteractionsByCustomerId(customerId)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, interactions)
}

func UpdateInteraction(c *gin.Context) {
    interactionId := c.Param("id")
    var interaction models.Interaction
    if err := c.ShouldBindJSON(&interaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := services.UpdateInteraction(interactionId, &interaction); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, interaction)
}

func DeleteInteraction(c *gin.Context) {
    interactionId := c.Param("id")
    if err := services.DeleteInteraction(interactionId); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}
