package handler

import (
	"context"
	"example/transaction/model"
	"example/transaction/prisma/db"
	"example/transaction/utils/token"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	
	users, err := client.User.FindFirst(
        db.User.Email.Equals(loginRequest.Email),
    ).Exec(context.Background())

    if err != nil &&  err.Error() != "ErrNotFound" {
        log.Fatalf("Error searching for users: %v", err)
        c.JSON(http.StatusInternalServerError, ResponseErrorDetail(CreateErrorResp("Internal Server Error", err.Error())))
        return
    }

    if (users == nil) {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    err = VerifyPassword(loginRequest.Password, users.Password)

    if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusInternalServerError, ResponseErrorDetail(CreateErrorResp("Mismatch Hash Password", err.Error())))
        return
	}
  
    token,err := token.GenerateToken(users.Email)

    if err != nil {
        c.JSON(http.StatusInternalServerError, ResponseErrorDetail(CreateErrorResp("Failed Generate Token", err.Error())))
		return 
	}

	c.JSON(http.StatusOK, ResponseDataDetail(token))
}


func VerifyPassword(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}