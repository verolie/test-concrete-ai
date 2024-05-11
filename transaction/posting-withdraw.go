package transaction

import (
	"context"
	"example/transaction/model"
	"example/transaction/prisma/db"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func WithdrawProcess(c *gin.Context) {
	var payment model.DetailTransaction
	var err error

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
    	_, err = client.TransactionDetail.CreateOne(
        	db.TransactionDetail.TrxID.Equals(payment.Trx_id),
			db.TransactionDetail.Timestamps.Equals(time.Now()),
			db.TransactionDetail.ReceiverPan.Equals(""),
			db.TransactionDetail.SenderPan.Equals(payment.Sender_pan),
			db.TransactionDetail.ApvCode.Equals(payment.Apv_code),
			db.TransactionDetail.TrxTyp.Equals(payment.Trx_typ),
			db.TransactionDetail.Amt.Decrement(float64(payment.Amt)),
			db.TransactionDetail.Status.Equals(payment.Status),
			db.TransactionDetail.Desc.Equals(payment.Desc),
			db.TransactionDetail.AcctDetail.Link(db.AccountDetail.LocAcct.Equals(payment.Loc_acct)),
    	).Exec(context.Background())
   		if err != nil {
        	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting payment data"})
        	return
    	}
	}else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error Cannot Find Data"})
        return
	}

	c.JSON(http.StatusOK, ResponseDataDetail("withdraw success"))
}
