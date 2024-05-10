package transaction

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type DetailTransaction struct {
	Trx_id 		string
	TimeStamps  time.Time
	Apv_code    string
	Trx_typ		string
	Amt			float32
	Status		string
	Desc		string
	Loc_acct    string
	CreateAt    time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func GetDetailTransaction(c *gin.Context) {
	// client := db.NewClient
	// result := client.DetailTransaction

	client := db.NewClient()
    if err := client.Prisma.Connect(); err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    defer client.Prisma.Disconnect()
	


	// if result.Error != nil {
	// 	c.JSON(http.StatusBadRequest, "bad request")
	// }

	// resp := lastUpdateActivity(db, 0)

	// c.JSON(http.StatusOK, ResponseDataDetail(resp))
	c.JSON(http.StatusOK, "sa")
}

func GetDetailTransactionParam(c *gin.Context) {

}
