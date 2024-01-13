// uas_pemrogramanplatform/handlers/user_handler.go
package handlers

import (
	"net/http"
	"uas_pemrogramanplatform/database"
	"uas_pemrogramanplatform/models"

	"github.com/gin-gonic/gin"
)

// CreateUserHandler handles the creation of a new user
// @Summary Membuat user baru
// @Description Membuat user baru dengan detail yang diberikan
// @Accept json
// @Produce json
// @Param user body models.User true "Objek user yang akan dibuat"
// @Success 201 {object} models.User
// @Router /api/users/ [post]
func CreateUserHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}


// UpdateUserHandler handles the update of a user by ID
// @Summary Memperbarui user berdasarkan ID
// @Description Memperbarui user dengan detail yang diberikan
// @Accept json
// @Produce json
// @Param id path int true "ID user"
// @Param user body models.User true "Objek user yang akan diperbarui"
// @Success 200 {object} models.User
// @Router /api/users/{id} [put]
func UpdateUserHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DeleteUserHandler handles the deletion of a user by ID
// @Summary Menghapus user berdasarkan ID
// @Description Menghapus user berdasarkan ID mereka
// @Produce json
// @Param id path int true "ID user"
// @Success 200 {string} string
// @Router /api/users/{id} [delete]
func DeleteUserHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user tidak ditemukan"})
		return
	}

	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "user berhasil dihapus"})
}

// LoginHandler handles the user login operation
// @Summary Login user
// @Description Autentikasi user
// @Accept json
// @Produce json
// @Param loginUser body models.User true "Kredensial login user"
// @Success 200 {object} models.User
// @Router /api/login [post]
func LoginHandler(c *gin.Context) {
	var loginUser models.User
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFromDB models.User
	if err := database.DB.Where("email = ?", loginUser.Email).First(&userFromDB).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Kredensial tidak valid"})
		return
	}

	// TODO: Implementasi autentikasi login (misalnya, menggunakan bcrypt)

	c.JSON(http.StatusOK, gin.H{"message": "Login berhasil"})
}

// GetAllUsersHandler handles the retrieval of all users
// @Summary Mengambil semua user
// @Description Mengambil daftar semua user
// @Produce json
// @Success 200 {array} models.User
// @Router /api/users/ [get]
func GetAllUsersHandler(c *gin.Context) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kesalahan Server Internal"})
		return
	}

	c.JSON(http.StatusOK, users)
}
