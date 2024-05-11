package server

import (
	"example/transaction/handler"
	"example/transaction/transaction"

	"github.com/gin-gonic/gin"
)

func RunServer() {

	e := gin.Default()

	registerServer(e)

	e.Run()
}

func registerServer(e *gin.Engine) {
	//Acount Manager Service
	e.POST("/user/login", getUser)
	e.POST("/user/register", createUser)
	e.GET("/user/account/detail/:acct_num", getDetaiUserAccount)
	e.GET("/user/payment/history/:loc_acct", getUserTrnxHist) //blom test
	
	//Transaction
	e.POST("/transaction/send", paymentProcess)
	e.POST("/transaction/withdraw", withdrawProcess)
	e.GET("/transaction/detail", detailTransaction)
	e.GET("/transaction/detail/:trx_id", detailTransactionParam)
}

func getUser(c *gin.Context) {
	handler.GetUsers(c)
}
func createUser(c *gin.Context) {
	handler.CreateUser(c)
}
func getDetaiUserAccount(c *gin.Context) {
	handler.DetaiUserAccount(c)
}
func getUserTrnxHist(c *gin.Context) {
	handler.UserTrnxHist(c)
}
func paymentProcess(c *gin.Context) {
	transaction.PostingPayment(c)
}

func withdrawProcess(c *gin.Context) {
	transaction.WithdrawProcess(c)
}

func detailTransaction(c *gin.Context) {
	transaction.GetDetailTransaction(c)
}

func detailTransactionParam(c *gin.Context) {
	transaction.GetDetailTransactionParam(c)
}
