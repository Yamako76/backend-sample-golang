package infra

import (
	"gorm.io/gorm"

	"github.com/Yamako76/backend-sample-golang/infra/mysql"
)

func NewDB() (*gorm.DB, error) {
	return mysql.NewDB()
}
