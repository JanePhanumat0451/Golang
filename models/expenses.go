package models

type Expense struct {
	ID       uint    `json:"id" gorm:"primary_key"`
	Amount   float64 `json:"amount"`   // จำนวนเงินที่ใช้จ่าย
	Category string  `json:"category"` // ประเภทการใช้จ่าย เช่น Food, Travel
	Date     string  `json:"date"`     // วันที่ใช้จ่าย
	Notes    string  `json:"notes"`    // หมายเหตุเกี่ยวกับการใช้จ่าย
	UserID   uint    `json:"user_id"`  // ผู้ใช้ที่บันทึกการใช้จ่าย (เชื่อมโยงกับตาราง users)
}

type ExpenseInput struct {
	Amount   float64 `json:"amount"`   // จำนวนเงินที่ใช้จ่าย
	Category string  `json:"category"` // ประเภทการใช้จ่าย เช่น Food, Travel
	Date     string  `json:"date"`     // วันที่ใช้จ่าย
	Notes    string  `json:"notes"`    // หมายเหตุเกี่ยวกับการใช้จ่าย
}
