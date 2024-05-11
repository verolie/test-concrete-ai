package handler

import (
	"context"
	"log"
	"net/http"

	"example/transaction/model"
	"example/transaction/prisma/db"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var registerRequest model.RegisUser
	var err error

	if err := c.BindJSON(&registerRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

	client := db.NewClient()
    if err := client.Prisma.Connect(); err != nil {
        log.Fatalf("Error connecting to database: %v", err)
		return
    }
	
    defer client.Prisma.Disconnect()

	status, err := CheckAccount(client , registerRequest.Acct_num)
	if err != nil {
       	c.JSON(http.StatusInternalServerError, gin.H{"error": "When check account" })
       	return
	}

	if (!status) {
		_ , err = client.User.CreateOne(
			db.User.AcctNum.Set(registerRequest.Acct_num),
			db.User.Name.Set(registerRequest.Name),
			db.User.Email.Set(registerRequest.Email),
			db.User.Password.Set(registerRequest.Password),
			db.User.Address.Set(registerRequest.Address),
		).Exec(context.Background())
   		if err != nil {
        	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting payment data" })
        	return
    	}
	}

	_, err = client.AccountDetail.CreateOne(
		db.AccountDetail.LocAcct.Set(registerRequest.DetailAccount.Loc_acct),
		db.AccountDetail.PrinPan.Set(registerRequest.DetailAccount.Prin_pan),
		db.AccountDetail.AcctTyp.Set(registerRequest.DetailAccount.Acct_typ),
		db.AccountDetail.ActvTyp.Set(registerRequest.DetailAccount.Actv_typ),
		db.AccountDetail.BlncAmt.Divide(registerRequest.DetailAccount.Blnc_amt),
		db.AccountDetail.LoanAmt.Divide(registerRequest.DetailAccount.Loan_amt),
		db.AccountDetail.CyccDay.Divide(registerRequest.DetailAccount.Cycc_day),
		db.AccountDetail.MinLoanPymnt.Divide(registerRequest.DetailAccount.Min_loan_pymnt),
		db.AccountDetail.Acct.Link(
        	db.User.AcctNum.Equals(registerRequest.Acct_num),
    	),
	).Exec(context.Background())
   		if err != nil {
        	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting payment data" })
        	return
    	}



	c.JSON(http.StatusOK, ResponseDataDetail("user succesfully created"))
}

func CheckAccount(client *db.PrismaClient, Acctnum string) (bool, error) {
	userAcct, err := client.User.FindUnique(
        db.User.AcctNum.Equals(Acctnum),
    ).Exec(context.Background())
	if  userAcct == nil {
        return false, err
    }
	return true, err
}
