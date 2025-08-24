package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CalculationResult struct {
	gorm.Model
	FloatValue float64
	IntValue   int
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
	calcInt()
	calcIntWithError()
	calcBigRat()
}

func calcFloat() {
	var f float64
	for i := 0; i < 10; i++ {
		f += 0.1
	}
	fmt.Println(f)
	DB.Create(&CalculationResult{FloatValue: f})
}

func calcInt() {
	var i int
	k := int(0.1 * 10000)
	for j := 0; j < 10; j++ {
		i += k
	}
	fmt.Println(i)
	calcResult := CalculationResult{IntValue: i}
	DB.Create(&calcResult)
	fmt.Println(float64(calcResult.IntValue) / 10000.0)
}

func calcIntWithError() {
	fmt.Println("--Error case with 1.0/49.0 --")
	var i int
	val := (1.0 / 49.0) * 100000000.0
	k := int(val)
	for j := 0; j < 49; j++ {
		i += k
	}
	calcResult := CalculationResult{IntValue: i}
	DB.Create(&calcResult)
	fmt.Printf("i = %d\n", i)
	result := float64(i) / 100000000.0
	fmt.Printf("Expected: 1.0, Actual: %.20f\n", result)
}

func calcBigRat() {
	fmt.Println("--- big.Rat case with 1/49 ---")
	// 1/49を表現するRatを作成
	r := big.NewRat(1, 49)

	// 合計を保持するRatを作成
	sum := big.NewRat(0, 1)

	// 49回足し合わせる
	for i := 0; i < 49; i++ {
		sum.Add(sum, r)
	}

	// 期待値である1 (1/1)
	one := big.NewRat(1, 1)

	// 結果を比較
	if sum.Cmp(one) == 0 {
		fmt.Println("Correct! The result is exactly 1.")
	} else {
		fmt.Printf("Error! The result is not 1. It is %s\n", sum.String())
	}

	// 結果を文字列や浮動小数点数で表示
	fmt.Printf("Result as a fraction: %s\n", sum.String())
	f64, _ := sum.Float64()
	fmt.Printf("Result as float64: %.20f\n", f64)
}
