package handler

import (
	"context"
	"log"
	"net/http"

	"example/transaction/model"
	"example/transaction/prisma/db"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var registerRequest model.RegisUser
	var err error

	if err := c.BindJSON(&registerRequest); err != nil {
        c.JSON(http.StatusBadRequest,  ResponseErrorDetail(CreateErrorResp("Invalid request body", err.Error())))
        return
    }

	client := db.NewClient()
    if err := client.Prisma.Connect(); err != nil {
        log.Fatalf("Error connecting to database: %v", err)
		return
    }
	
    defer client.Prisma.Disconnect()

	status, err := CheckAccount(client , registerRequest.Acct_num)
	if  err != nil && err.Error() != "ErrNotFound"  {
		println(err)
       	c.JSON(http.StatusInternalServerError, ResponseErrorDetail(CreateErrorResp("Error when check account", err.Error())))
       	return
	}

	if (!status) {
		hashPass, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password),bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ResponseErrorDetail(CreateErrorResp("error when hash password", err.Error())))
			return
		}

		registerRequest.Password = string(hashPass)

		print(registerRequest.Password)

		_, err = client.User.CreateOne(
			db.User.AcctNum.Set(registerRequest.Acct_num),
			db.User.Name.Set(registerRequest.Name),
			db.User.Email.Set(registerRequest.Email),
			db.User.Password.Set(registerRequest.Password),
			db.User.Address.Set(registerRequest.Address),
		).Exec(context.Background())
   		if err != nil && err.Error() != "ErrNotFound"  {
        	c.JSON(http.StatusInternalServerError,  ResponseErrorDetail(CreateErrorResp("Error inserting User data", err.Error())))
        	return
    	}
	}

	_, err = client.AccountDetail.CreateOne(
		db.AccountDetail.LocAcct.Set(registerRequest.DetailAccount.Loc_acct),
		db.AccountDetail.PrinPan.Set(registerRequest.DetailAccount.Prin_pan),
		db.AccountDetail.AcctTyp.Set(registerRequest.DetailAccount.Acct_typ),
		db.AccountDetail.ActvTyp.Set(registerRequest.DetailAccount.Actv_typ),
		db.AccountDetail.BlncAmt.Set(registerRequest.DetailAccount.Blnc_amt),
		db.AccountDetail.LoanAmt.Set(registerRequest.DetailAccount.Loan_amt),
		db.AccountDetail.CyccDay.Set(registerRequest.DetailAccount.Cycc_day),
		db.AccountDetail.MinLoanPymnt.Set(registerRequest.DetailAccount.Min_loan_pymnt),
		db.AccountDetail.Acct.Link(
        	db.User.AcctNum.Equals(registerRequest.Acct_num),
    	),
	).Exec(context.Background())

   		if err != nil && err.Error() != "ErrNotFound"  {
        	c.JSON(http.StatusInternalServerError, ResponseErrorDetail(CreateErrorResp("Error inserting payment data", err.Error())))
        	return
    	}



	c.JSON(http.StatusOK, ResponseDataDetail("user succesfully created"))
}

func CheckAccount(client *db.PrismaClient, Acctnum string) (bool, error) {
	userAcct, err := client.User.FindUnique(
        db.User.AcctNum.Equals(Acctnum),
    ).Exec(context.Background())
	if err != nil && err.Error() != "ErrNotFound" {
		println("masuk sini")
        return false, err
    }
    if userAcct == nil {
        return false, nil // User account does not exist
    }
    return true, nil // User account exists
}
