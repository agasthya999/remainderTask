package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBRead  *gorm.DB
	DBWrite *gorm.DB
	DB      *gorm.DB
	err     error
)

func GormConnect() {
	dsn_rd := fmt.Sprintf("host=localhost user=postgres password=******* dbname=remainderTask port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	DBRead, err = gorm.Open(postgres.Open(dsn_rd), &gorm.Config{})
	if err != nil {
		panic("failed to connect database reader")
	}
	fmt.Println("Success DB")
}
