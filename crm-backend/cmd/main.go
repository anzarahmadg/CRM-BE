package main

import (
    "github.com/gin-gonic/gin"
    "crm-backend/config"
    "crm-backend/routes"
)

func main() {
    // Load configuration
    cfg := config.LoadConfig()

    // Initialize Gin
    router := gin.Default()

    // Setup routes
    routes.SetupRoutes(router)

    // Start server
    router.Run(":8080")
}
