package routes

import (
	"net/http"

	"github.com/Freemasoid/go-practice-rest-api/models"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	var user models.Event

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not save user"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}
