package start

import (
	"fmt"
	"e1/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDBEngine() (*gorm.DB, error) {

	dsn := "giant:1234@tcp(127.0.0.1:3306)/blog_service?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	global.DBEngine = db
	return db, nil
}
func DoTest() {

	fmt.Println("ddd")
}
