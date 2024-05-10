package server

import (
	"example/transaction/handler"
	"example/transaction/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunServer() {

	e := gin.Default()

	registerServer(e)

	e.Run()
}

func registerServer(e *gin.Engine) {
	//Acount Manager Service
	e.GET("/user", usersHandler)
	e.POST("/user", usersHandler)

	//Transaction
	e.POST("/transaction/send", paymentProcess)
	e.POST("/transaction/withdraw", withdrawProcess)
	e.GET("/transaction/detail", detailTransaction)
	e.GET("/transaction/detail/:id", detailTransactionParam)
}

func usersHandler(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		handler.GetUsers(c)
	case http.MethodPost:
		handler.CreateUser(c)
	default:
		 c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func paymentProcess(c *gin.Context) {
	transaction.PostingPayment(c)
}

func withdrawProcess(c *gin.Context) {
	transaction.Withdraw(c)
}

func detailTransaction(c *gin.Context) {
	transaction.GetDetailTransaction(c)
}

func detailTransactionParam(c *gin.Context) {
	transaction.GetDetailTransactionParam(c)
}
