package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "crm-backend/models"
    "crm-backend/services"
)

func CreateCustomer(c *gin.Context) {
    var customer models.Customer
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := services.CreateCustomer(&customer); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, customer)
}

func GetCustomers(c *gin.Context) {
    customers, err := services.GetCustomers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, customers)
}

func UpdateCustomer(c *gin.Context) {
    customerId := c.Param("id")
    var customer models.Customer
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := services.UpdateCustomer(customerId, &customer); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, customer)
}

func DeleteCustomer(c *gin.Context) {
    customerId := c.Param("id")
    if err := services.DeleteCustomer(customerId); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, nil)
}
