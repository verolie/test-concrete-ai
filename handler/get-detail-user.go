package handler

import (
	"context"
	"log"
	"net/http"

	"example/transaction/prisma/db"

	"github.com/gin-gonic/gin"
)

func DetaiUserAccount(c *gin.Context) {
	client := db.NewClient()
    if err := client.Prisma.Connect(); err != nil {
        log.Fatalf("Error connecting to database: %v", err)
		return
    }
    defer client.Prisma.Disconnect()

	acctNumber := c.Param("acct_num")
    if acctNumber == "" {
        c.JSON(http.StatusBadRequest, ResponseErrorDetail(CreateErrorResp("Transaction ID is required", "")))
        return
    }

	resp, err := client.AccountDetail.FindMany(db.AccountDetail.AcctNum.Equals(acctNumber)).Exec(context.Background())
    if err != nil {
        c.JSON(http.StatusInternalServerError, ResponseErrorDetail(CreateErrorResp("Error retrieving transaction details", "")))
        return
    }

	c.JSON(http.StatusOK, ResponseDataDetail(resp))
}
