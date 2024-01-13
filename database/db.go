// database/db.go
package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	DB_PORT := "3306"
	DB_USER := "root"
	DB_PASSWORD := ""
	DB_HOST := "localhost"
	DB_NAME := "uas_pemrogramanplatform"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")

	// Auto Migrate
	DB.AutoMigrate(&User{}, &Product{})

}



// User model
type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
}

// Product model
type Product struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	NamaProduk string  `json:"nama_produk"`
	Harga      float64 `json:"harga"`
}
