package handler

import (
	"context"
	"example/transaction/model"
	"example/transaction/prisma/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var loginRequest model.LoginRequest
	if err := c.BindJSON(&loginRequest); err != nil {
        c.JSON(http.StatusBadRequest, ResponseErrorDetail(CreateErrorResp("Invalid request body", err.Error())))
        return
    }

	client := db.NewClient()
    if err := client.Prisma.Connect(); err != nil {
        log.Fatalf("Error connecting to database: %v", err)
		return
    }

    defer client.Prisma.Disconnect()
	
	users, err := client.User.FindMany(
        db.User.Email.Equals(loginRequest.Email),
        db.User.Password.Equals(loginRequest.Password),
    ).Exec(context.Background())

    if err != nil {
        log.Fatalf("Error searching for users: %v", err)
        c.JSON(http.StatusInternalServerError, ResponseErrorDetail(CreateErrorResp("Internal Server Error", err.Error())))
        return
    }
  
    if (len(users) == 0) {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    } else if (len(users) > 1){
		c.JSON(http.StatusNotFound, gin.H{"error": "More than 1 account"})
        return
	}


	c.JSON(http.StatusOK, ResponseDataDetail("Success Login"))
}
