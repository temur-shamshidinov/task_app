package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/temur-shamshidinov/task_app/models"
	"github.com/temur-shamshidinov/task_app/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func (h *handler) Register(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "So'rov xatolik bilan yuborildi"})
		return
	}

	
	hashedPassword, err := utils.HashPassword(user.PasswordHash)
	if err != nil {
		c.JSON(500, gin.H{"error": "Parolni xashlashda xatolik"})
		return
	}

	user.PasswordHash = hashedPassword

	
	if err := h.storage.GetUserRepo().CreateUser(c, user); err != nil {
		c.JSON(500, gin.H{"error": "Foydalanuvchini yaratishda xatolik"})
		return
	}

	c.JSON(200, gin.H{"message": "Foydalanuvchi muvaffaqiyatli ro'yxatdan o'tdi"})
}

func (h *handler) Login(c *gin.Context) {

	var loginReq models.LoginReq

	
	if err := c.BindJSON(&loginReq); err != nil {
		c.JSON(400, gin.H{"error": "So'rov xatolik bilan yuborildi"})
		return
	}

	
	users, err := h.storage.GetUserRepo().GetUserByEmail(c, loginReq.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": "Foydalanuvchini olishda xatolik"})
		return
	}

	if len(users) == 0 {
		c.JSON(401, gin.H{"error": "Email yoki parol noto'g'ri"})
		return
	}

	user := users[0]

	
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginReq.PasswordHash)); err != nil {
		c.JSON(401, gin.H{"error": "Email yoki parol noto'g'ri"})
		return
	}

	
	token := "some-generated-jwt-token"

	c.JSON(200, gin.H{"token": token})
}
