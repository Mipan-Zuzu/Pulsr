package main

import (
	"fmt"
	"net/http"
	"pulsr/internal/cache"
	"pulsr/internal/database"
	"pulsr/internal/model"
	"pulsr/internal/routing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupCORS(route *gin.Engine) {
	route.Use(func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Vary", "Origin")
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})
}

func main() {
	godotenv.Load()
	route := gin.Default()
	setupCORS(route)
	port := "3031"
	db, err := database.PostgresConnection()
	ErrModelAccsesToken := db.AutoMigrate(&model.AccsesTokenSupabase{})

	rdb := cache.GetClient()

	if ErrModelAccsesToken != nil {
		fmt.Println(ErrModelAccsesToken)
	}
	if err != nil {
		fmt.Println(err)
	}
	routing.SetupRouting(db, route)
	routing.SetupRedisUpstash(route, rdb)
	route.Run(fmt.Sprintf(":%s", port))
	fmt.Println(fmt.Printf("Server Listen in Port %s", port))
}
