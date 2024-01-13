// models/models.go
package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
}

type Product struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	NamaProduk string  `json:"nama_produk"`
	Harga      float64 `json:"harga"`
}
