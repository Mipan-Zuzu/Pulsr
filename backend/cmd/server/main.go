package main

import (
	"fmt"
	"pulsr/internal/database"
	"pulsr/internal/routing"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"pulsr/internal/model"
)

func main () {
	godotenv.Load()
	route := gin.Default()
	port := "3031"
	db, err := database.PostgresConnection()
	ErrModelAccsesToken := db.AutoMigrate(&model.AccsesTokenSupabase{})
	
	if ErrModelAccsesToken != nil {
		fmt.Println(ErrModelAccsesToken)
	}
	if err != nil {
		fmt.Println(err)
	}
	routing.SetupRouting(db, route)
	route.Run(fmt.Sprintf(":%s", port))
	fmt.Println(fmt.Printf("Server Listen in Port %s", port))
}
