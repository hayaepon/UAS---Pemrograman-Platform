package main

import (
	"uas_pemrogramanplatform/database"
	"uas_pemrogramanplatform/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "uas_pemrogramanplatform/docs" // Import generated docs
)

// @title UAS - Pemrograman Berbasis Platform
// @description Membuat API CRUD dengan menggunakan bahasa GO
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http https

func main() {
	// Inisialisasi koneksi database
	database.InitDB()

	// Membuat router baru
	r := gin.Default()

	// Middleware untuk mendukung CORS
	r.Use(cors.Default())

	// Grouping untuk API Pengguna (Users)
	usersGroup := r.Group("/api/users")
	{
		// @Summary Memperbarui user berdasarkan ID
		// @Description Memperbarui user dengan payload JSON
		// @Accept json
		// @Produce json
		// @Param id path int true "ID user"
		// @Param user body User true "Objek user yang akan diperbarui"
		// @Success 200 {object} User
		// @Router /api/users/{id} [put]
		usersGroup.PUT("/{id}", handlers.UpdateUserHandler)

		// @Summary Hapus user berdasarkan ID
		// @Description Hapus user berdasarkan ID mereka
		// @Produce json
		// @Param id path int true "ID user"
		// @Success 200 {string} string
		// @Router /api/users/{id} [delete]
		usersGroup.DELETE("/{id}", handlers.DeleteUserHandler)

		// @Summary Buat user baru
		// @Description Buat user baru dengan payload JSON
		// @Accept json
		// @Produce json
		// @Param user body User true "Objek user yang akan dibuat"
		// @Success 201 {object} User
		// @Router /api/users/ [post]
		usersGroup.POST("/", handlers.CreateUserHandler)

		// @Summary Ambil semua user
		// @Description Ambil daftar semua user
		// @Produce json
		// @Success 200 {array} User
		// @Router /api/users/ [get]
		usersGroup.GET("/", handlers.GetAllUsersHandler)
	}

	// Grouping untuk API Produk (Products)
	productsGroup := r.Group("/api/products")
	{
		// @Summary Buat produk baru
		// @Description Buat produk baru dengan payload JSON
		// @Accept json
		// @Produce json
		// @Param product body Product true "Objek produk yang akan dibuat"
		// @Success 201 {object} Product
		// @Router /api/products/ [post]
		productsGroup.POST("/", handlers.CreateProductHandler)

		// @Summary Memperbarui produk berdasarkan ID
		// @Description Memperbarui produk dengan payload JSON
		// @Accept json
		// @Produce json
		// @Param id path int true "ID Produk"
		// @Param product body Product true "Objek produk yang akan diperbarui"
		// @Success 200 {object} Product
		// @Router /api/products/{id} [put]
		productsGroup.PUT("/{id}", handlers.UpdateProductHandler)

		// @Summary Hapus produk berdasarkan ID
		// @Description Hapus produk berdasarkan ID mereka
		// @Produce json
		// @Param id path int true "ID Produk"
		// @Success 200 {string} string
		// @Router /api/products/{id} [delete]
		productsGroup.DELETE("/{id}", handlers.DeleteProductHandler)

		// @Summary Ambil semua produk
		// @Description Ambil daftar semua produk
		// @Produce json
		// @Success 200 {array} Product
		// @Router /api/products/ [get]
		productsGroup.GET("/", handlers.GetAllProductsHandler)
	}

	// @Summary Login pengguna
	// @Description Melakukan login pengguna dengan payload JSON
	// @Accept json
	// @Produce json
	// @Param user body User true "Kredensial login pengguna"
	// @Success 200 {object} LoginResponse
	// @Router /api/login [post]
	r.POST("/api/login", handlers.LoginHandler)

	// @Summary Endpoint Swagger
	// @Description Endpoint Swagger UI
	// @Produce html
	// @Router /swagger/*any [get]
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

	// Menjalankan server
	r.Run(":8080")
}
