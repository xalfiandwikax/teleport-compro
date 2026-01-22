package main

import (
    "log"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    
    router.GET("/api/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })
    
    router.POST("/api/auth/login", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Login endpoint - under development",
            "token": "dummy-token",
            "user": map[string]interface{}{
                "id": 1,
                "name": "Test User",
                "email": "test@mail.com",
                "role": "admin",
            },
        })
    })
    
    log.Println("🚀 Server starting on :8080")
    log.Println("📝 GET  http://localhost:8080/api/health")
    log.Println("📝 POST http://localhost:8080/api/auth/login")
    
    router.Run(":8080")
}
