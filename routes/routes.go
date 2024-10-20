package routes

import (
	"Gofinal1/controllers" // นำเข้า controllers ที่เก็บฟังก์ชันต่าง ๆ เช่น Login, Register
	"Gofinal1/middlewares" // นำเข้า middlewares ที่เก็บ AuthMiddleware

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// เส้นทางสำหรับการลงทะเบียนและเข้าสู่ระบบ
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// กลุ่มเส้นทางที่ใช้ Middleware สำหรับตรวจสอบ JWT Token
	protected := r.Group("/", middlewares.AuthMiddleware())
	{
		protected.GET("/expenses", controllers.GetExpenses)
		protected.POST("/expenses", controllers.CreateExpense)
		protected.PUT("/expenses/:id", controllers.UpdateExpense)
		protected.DELETE("/expenses/:id", controllers.DeleteExpense)
		protected.GET("/expenses/:id", controllers.GetExpenseByID)

	}
}
