package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Document struct {
	gorm.Model
	TickedID  string `gorm:"type:varchar(100);unique_index"`
	Content   string `gorm:"type:varchar(100)"`
	Title     string `gorm:"type:varchar(100)"`
	Author    string `gorm:"type:varcahr(100)"`
	Topic     string `gorm:"type:varcahr(100)"`
	Watermark string `gorm:"type:varcahr(100)"`
}

func Init(dialect, host, port, dbname, pass string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dialect, hsot, port, dbname, pass))

	return db, err
	defer db.close()
}
