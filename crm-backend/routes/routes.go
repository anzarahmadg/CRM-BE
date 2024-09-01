package routes

import (
    "github.com/gin-gonic/gin"
    "crm-backend/controllers"
)

func SetupRoutes(router *gin.Engine) {
    userGroup := router.Group("/api/users")
    {
        userGroup.POST("/", controllers.CreateUser)
        userGroup.GET("/", controllers.GetUsers)
        userGroup.PUT("/:id", controllers.UpdateUser)
        userGroup.DELETE("/:id", controllers.DeleteUser)
    }
    
    customerGroup := router.Group("/api/customers")
    {
        customerGroup.POST("/", controllers.CreateCustomer)
        customerGroup.GET("/", controllers.GetCustomers)
        customerGroup.PUT("/:id", controllers.UpdateCustomer)
        customerGroup.DELETE("/:id", controllers.DeleteCustomer)
    }

    interactionGroup := router.Group("/api/interactions")
    {
        interactionGroup.POST("/", controllers.CreateInteraction)
        interactionGroup.GET("/customer/:customerId", controllers.GetCustomerInteractions)
        interactionGroup.PUT("/:id", controllers.UpdateInteraction)
        interactionGroup.DELETE("/:id", controllers.DeleteInteraction)
    }
}
