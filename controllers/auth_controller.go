package controllers

import (
	"Gofinal1/config"
	"Gofinal1/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key")

// Claims โครงสร้างข้อมูลที่บรรจุข้อมูลโทเค็น
type Claims struct {
	UserID uint   `json:"user_id"` // เพิ่มฟิลด์สำหรับ UserID
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.StandardClaims
}

// Login ฟังก์ชันสำหรับการเข้าสู่ระบบ
func Login(c *gin.Context) {
	var input models.LoginInput
	var user models.User

	// ตรวจสอบข้อมูลที่รับเข้ามาว่าถูกต้องหรือไม่
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบอีเมลและรหัสผ่านในฐานข้อมูล
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// เปรียบเทียบรหัสผ่านที่เข้ารหัส
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// กำหนดเวลาหมดอายุของโทเค็น
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		UserID: user.ID, // บันทึก UserID ของผู้ใช้ที่เข้าสู่ระบบ
		Email:  user.Email,
		Name:   user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// สร้าง JWT โทเค็น
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	// บันทึก token ลงในฐานข้อมูล jwt_tokens
	jwtToken := models.JWTToken{
		UserID:    user.ID,
		Token:     tokenString,
		ExpiresAt: expirationTime,
	}

	if err := config.DB.Create(&jwtToken).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save token"})
		return
	}

	// ส่งโทเค็นกลับไปยังผู้ใช้
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Register(c *gin.Context) {
	var input models.RegisterInput
	var existingUser models.User

	// ตรวจสอบว่ามีข้อมูลการสมัครเข้ามาครบถ้วน
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// ตรวจสอบว่า email ถูกใช้ไปแล้วหรือไม่
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
		return
	}

	// ตรวจสอบว่าชื่อและรหัสผ่านไม่ว่างเปล่า
	if input.Name == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name or Password cannot be empty"})
		return
	}

	// เข้ารหัสรหัสผ่านก่อนบันทึก
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encrypting password"})
		return
	}

	// สร้าง user ใหม่
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	// บันทึกข้อมูลผู้ใช้ในฐานข้อมูล
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งข้อความสำเร็จ
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
