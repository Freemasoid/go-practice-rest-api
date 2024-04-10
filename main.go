package main

import (
	"github.com/Freemasoid/go-practice-rest-api/db"
	"github.com/Freemasoid/go-practice-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
