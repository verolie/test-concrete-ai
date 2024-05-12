package handler

import (
	"context"
	"log"
	"net/http"

	"example/transaction/prisma/db"

	"github.com/gin-gonic/gin"
)

func UserTrnxHist(c *gin.Context) {
	client := db.NewClient()
    if err := client.Prisma.Connect(); err != nil {
        log.Fatalf("Error connecting to database: %v", err)
		return
    }
    defer client.Prisma.Disconnect()

	locAcct := c.Param("loc_acct")
    if locAcct == "" {
        c.JSON(http.StatusBadRequest, ResponseErrorDetail(CreateErrorResp("Transaction ID is required", "")))
        return
    }

	resp, err := client.TransactionDetail.FindMany(db.TransactionDetail.LocAcct.Equals(locAcct)).Exec(context.Background())
    if err != nil {
        c.JSON(http.StatusInternalServerError,  ResponseErrorDetail(CreateErrorResp("Error retrieving transaction details", err.Error())))
        return
    }

	c.JSON(http.StatusOK, ResponseDataDetail(resp))
}
