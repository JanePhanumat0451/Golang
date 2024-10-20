package models

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type JWTToken struct {
	ID        uint      `gorm:"primaryKey"`                // รหัสอัตโนมัติสำหรับแต่ละ token
	UserID    uint      `gorm:"not null"`                  // เชื่อมโยงกับผู้ใช้
	Token     string    `gorm:"not null"`                  // เก็บค่า JWT token
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"` // เวลาที่สร้าง token
	ExpiresAt time.Time `gorm:"not null"`                  // เวลาหมดอายุของ token
	IsRevoked bool      `gorm:"default:false"`             // แสดงว่า token นี้ถูกยกเลิกหรือไม่
}
