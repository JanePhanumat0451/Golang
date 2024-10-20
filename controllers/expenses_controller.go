package controllers

import (
	"Gofinal1/config"
	"Gofinal1/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateExpense(c *gin.Context) {
	amountStr := c.PostForm("amount")
	category := c.PostForm("category")
	notes := c.PostForm("notes")
	date := c.PostForm("date")

	// แปลง amount จาก string เป็น float64
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Amount must be a valid number"})
		return
	}

	// ตรวจสอบว่าผู้ใช้ได้ส่ง token และยืนยันว่าเป็นผู้ใช้ที่ถูกต้องหรือไม่
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// บันทึกข้อมูลรายจ่าย โดยเชื่อมโยงกับ UserID ของผู้ใช้ที่เข้าสู่ระบบ
	expense := models.Expense{
		Amount:   amount,
		Category: category,
		Notes:    notes,
		Date:     date,
		UserID:   userID.(uint), // เพิ่ม userID ที่ได้จาก token
	}
	result := config.DB.Create(&expense)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create expense"})
		return
	}

	// ส่งข้อมูลกลับไป
	c.JSON(http.StatusOK, gin.H{"data": expense})
}

// GetExpenses ฟังก์ชันสำหรับดึงข้อมูลรายจ่ายเฉพาะของผู้ใช้ที่เข้าสู่ระบบ
func GetExpenses(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var expenses []models.Expense
	config.DB.Where("user_id = ?", userID).Find(&expenses) // ดึงเฉพาะข้อมูลของผู้ใช้ที่เข้าสู่ระบบ
	c.JSON(http.StatusOK, gin.H{"data": expenses})
}

func GetExpenseByID(c *gin.Context) {
	var expense models.Expense
	// ดึงข้อมูลรายการรายจ่ายตาม ID
	if err := config.DB.Where("id = ?", c.Param("id")).First(&expense).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"}) // ส่งกลับ 404 ถ้าไม่พบข้อมูล
		return
	}

	// ส่งข้อมูลรายการรายจ่ายกลับไป
	c.JSON(http.StatusOK, gin.H{"data": expense})
}

// UpdateExpense ฟังก์ชันสำหรับอัปเดตรายจ่ายตาม ID
func UpdateExpense(c *gin.Context) {
	// ดึงข้อมูลรายการรายจ่ายตาม ID
	var expense models.Expense
	if err := config.DB.Where("id = ?", c.Param("id")).First(&expense).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Expense not found"})
		return
	}

	// เก็บค่า UserID เดิมเพื่อป้องกันการเขียนทับ
	originalUserID := expense.UserID

	// อัปเดตฟิลด์ที่ส่งเข้ามาใหม่
	var input models.ExpenseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// อัปเดตฟิลด์ที่อนุญาตให้อัปเดตเท่านั้น
	expense.Amount = input.Amount
	expense.Category = input.Category
	expense.Notes = input.Notes
	expense.Date = input.Date

	// ยืนยันว่าค่า UserID ยังคงเหมือนเดิม
	expense.UserID = originalUserID

	// บันทึกการเปลี่ยนแปลง
	if err := config.DB.Save(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update expense"})
		return
	}

	// ส่งข้อมูลที่อัปเดตกลับไป
	c.JSON(http.StatusOK, gin.H{"data": expense})
}

// DeleteExpense ฟังก์ชันสำหรับลบรายจ่ายตาม ID
func DeleteExpense(c *gin.Context) {
	var expense models.Expense
	if err := config.DB.Where("id = ?", c.Param("id")).First(&expense).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Expense not found"})
		return
	}

	config.DB.Delete(&expense)
	c.JSON(http.StatusOK, gin.H{"message": "Expense deleted"})
}
