package transaction

import (
	"context"
	"example/transaction/model"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"example/transaction/prisma/db"
)

var acctTyp string
var actvTyp string
var blncAmt, loanAmt, minLoanPymnt float64

func PostingPayment(c *gin.Context) {
	var payment model.DetailTransaction
	var err error
	var resp *db.TransactionDetailModel

	if err = c.BindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

	client := db.NewClient()
    if err = client.Prisma.Connect(); err != nil {
        log.Fatalf("Error connecting to database: %v", err)
		return
    }
    defer client.Prisma.Disconnect()

	if(!CheckAccount(client, payment.Loc_acct)){
		//try to update account
		if (actvTyp != "W") {
			UpdateAccount(c, client, payment);
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Account Already Write Off"})
        	return
		}

		//insert when payment success
    	resp, err = client.TransactionDetail.CreateOne(
        	db.TransactionDetail.TrxID.Set(payment.Trx_id),
        	db.TransactionDetail.Timestamps.Set(time.Now()),
        	db.TransactionDetail.ApvCode.Set(payment.Apv_code),
        	db.TransactionDetail.TrxTyp.Set(payment.Trx_typ),
        	db.TransactionDetail.Amt.Set(float64(payment.Amt)),
        	db.TransactionDetail.Status.Set(payment.Status),
        	db.TransactionDetail.Desc.Set(payment.Desc),
        	db.TransactionDetail.LocAcct.Set(payment.Loc_acct),
			// db.TransactionDetail.ReceiverPan.Set(payment.Receiver_account),
			db.TransactionDetail.SenderPan.Set(payment.Sender_account),
    	).Exec(context.Background())
   		if err != nil {
        	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting payment data"})
        	return
    	}
	}else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Cannot Find Data"})
        return
	}
	
	c.JSON(http.StatusOK, ResponseDataDetail(resp))
}

func CheckAccount(client *db.PrismaClient, Loc_acct string) bool {
	accountDetail, err := client.AccountDetail.FindUnique(
        db.AccountDetail.LocAcct.Equals(Loc_acct),
    ).Exec(context.Background())
	if err != nil || accountDetail == nil {
        return false
    }
		acctTyp = accountDetail.AcctTyp
		actvTyp = accountDetail.ActvTyp
		blncAmt = accountDetail.BlncAmt
		loanAmt = accountDetail.LoanAmt
		minLoanPymnt = accountDetail.MinLoanPymnt
	
	return true
}

func UpdateAccount(c *gin.Context, client *db.PrismaClient, payment model.DetailTransaction) {
	if (payment.Trx_typ == "C"){
		blncAmt -= float64(payment.Amt)

		if (blncAmt < 0) {
			if (acctTyp == "C" || acctTyp == "PL") {
				loanAmt += float64(payment.Amt)
				minLoanPymnt = (loanAmt + float64(payment.Amt)) * 0.1
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Account Balance Not enough"})
            	return
			}
		}
		
	} else if (payment.Trx_typ == "D"){
		blncAmt += float64(payment.Amt)
		if (acctTyp == "C" || acctTyp == "PL") {
			loanAmt -= float64(payment.Amt)
			minLoanPymnt = (loanAmt + float64(payment.Amt)) * 0.1
			if loanAmt < 0 {
				blncAmt += math.Abs(loanAmt)
			}
		}
	}
	
	_ , err :=   client.AccountDetail.FindMany(db.AccountDetail.LocAcct.Equals(payment.Loc_acct)).Update(
		db.AccountDetail.BlncAmt.Set(blncAmt),
		db.AccountDetail.LoanAmt.Set(loanAmt),
		db.AccountDetail.MinLoanPymnt.Set(minLoanPymnt),
	).Exec(context.Background())

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Update "})
        return
    }
}
