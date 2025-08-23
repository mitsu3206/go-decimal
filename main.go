package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CalculationResult struct {
	gorm.Model
	FloatValue float64
}

var DB *gorm.DB

func main() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Println(err.Error())
	}
	conf := &mysql.Config{
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Net:       "tcp",
		Addr:      fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName:    os.Getenv("DB_NAME"),
		Loc:       jst,
		ParseTime: true,
	}
	db, err := gorm.Open(gmysql.Open(conf.FormatDSN()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db

	if err := db.AutoMigrate(&CalculationResult{}); err != nil {
		panic("failed to migrate database")
	}
	calcFloat()
}

func calcFloat() {
	var f float64
	for i := 0; i < 10; i++ {
		f += 0.1
	}
	fmt.Println(f)
	DB.Create(&CalculationResult{FloatValue: f})
}
