package tables

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name string
	Url  string
}
