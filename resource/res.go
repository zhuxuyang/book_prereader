package resource

import (
	"github.com/jinzhu/gorm"
	"github.com/book_prereader/tables"
)

var (
	DB        *gorm.DB
)

func InitResource() {
	DB = tables.GetDBFromVipper()
}
