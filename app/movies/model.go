package movies

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Movie struct {
	gorm.Model
	Title string `gorm:"unique" json:"title"`
}


func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Movie{})
	return db
}
