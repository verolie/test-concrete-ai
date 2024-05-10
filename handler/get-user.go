package handler

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Activities struct {
	Activity_id int `gorm:"primary_key;auto_increment;not_null"`
	Title       string
	Email       string
	CreateAt    time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func GetUsers(c *gin.Context) {

}
