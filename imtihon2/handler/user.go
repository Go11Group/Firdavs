package handler

import (
	"fmt"
	"n11/Firdavs/imtihon2/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser - yangi foydalanuvchi yaratish uchun handler funksiyasi
func (hand *Handler) CreateUser(g *gin.Context) {
	var user model.User
	// JSON requestini parslash
	err := g.BindJSON(&user)
	if err != nil {
		// Agar xato bo'lsa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.String(http.StatusBadRequest, err.Error())
		return
	}
	// Foydalanuvchini yaratish
	err = hand.User.CreateUser(&user)
	if err != nil {
		// Agar xato bo'lsa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Yaratilgan foydalanuvchini 201 Created status kodi bilan qaytaradi
	g.JSON(http.StatusCreated, user)
}

// UpdateUser - mavjud foydalanuvchini yangilash uchun handler funksiyasi
func (hand *Handler) UpdateUser(g *gin.Context) {
	var user model.User
	// Requestni parslash
	err := g.Bind(&user)
	if err != nil {
		// Agar xato bo'lsa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.String(http.StatusBadRequest, err.Error())
		return
	}
	// Foydalanuvchini yangilash
	err = hand.User.UpdateUser(&user)
	if err != nil {
		// Agar xato bo'lsa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Muvaffaqiyatli yangilanganini bildiradi
	g.String(http.StatusOK, "SUCCESS")
}

// GetAllUsers - barcha foydalanuvchilarni olish uchun handler funksiyasi
func (hand *Handler) GetAllUsers(g *gin.Context) {
	// Barcha foydalanuvchilarni olish
	users, err := hand.User.ReadAllUsers()
	if err != nil {
		// Agar xato bo'lsa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Barcha foydalanuvchilarni 200 OK status kodi bilan qaytaradi
	g.JSON(http.StatusOK, users)
}

// GetUserById - foydalanuvchini ID orqali olish uchun handler funksiyasi
func (hand *Handler) GetUserById(g *gin.Context) {
	id := g.Param("user_id")
	if id == "" {
		// Agar ID kiritilmagan bo'lsa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.String(http.StatusBadRequest, "There is an error entering ID")
		return
	}

	// Foydalanuvchini ID orqali olish
	user, err := hand.User.ReadUser(id)
	if err != nil {
		// Agar xato bo'lsa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		fmt.Println("he;;d", err)
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Foydalanuvchini 200 OK status kodi bilan qaytaradi
	g.JSON(http.StatusOK, user)
}

// DeleteUsers - foydalanuvchini ID orqali o'chirish uchun handler funksiyasi
func (hand *Handler) DeleteUsers(g *gin.Context) {
	id := g.Param("user_id")
	if id == "" {
		// Agar ID kiritilmagan bo'lsa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.String(http.StatusBadRequest, "There is an error entering ID")
		return
	}

	// Foydalanuvchini o'chirish
	err := hand.User.DeleteUser(id)
	if err != nil {
		// Agar xato bo'lsa, 500 Internal Server Error status kodi bilan xatolikni qaytaradi
		g.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Muvaffaqiyatli o'chirilganini bildiradi
	g.String(http.StatusOK, "SUCCESS")
}

// GetFilterUsers - foydalanuvchilarni filtrlar orqali olish uchun handler funksiyasi
func (hand *Handler) GetFilterUsers(g *gin.Context) {
	// Filtr parametrlari
	userFilter := model.FilterUsers{
		Id:       g.Param("user_id"),
		Name:     g.Param("name"),
		Email:    g.Param("email"),
		Birthday: g.Param("birthday"),
		Password: g.Param("password"),
	}

	// Filtrlangan foydalanuvchilarni olish
	user, err := hand.User.FilterUsers(userFilter)
	if err != nil {
		// Agar xato bo'lsa, 400 Bad Request status kodi bilan xatolikni qaytaradi
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Filtrlangan foydalanuvchilarni 200 OK status kodi bilan qaytaradi
	g.JSON(http.StatusOK, user)
}
