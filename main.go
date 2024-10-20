package main

import (
	"Gofinal1/config"
	"Gofinal1/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// เสิร์ฟไฟล์สาธารณะจากโฟลเดอร์ public
	r.Static("/public", "./public")

	// โหลดเทมเพลต HTML หลายไฟล์
	r.LoadHTMLGlob("./public/*.html")

	// เส้นทางนี้จะส่งหน้า index.html เป็นหน้าแรก
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// เส้นทางสำหรับ register.html
	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register.html", nil)
	})

	// เชื่อมต่อกับฐานข้อมูล
	config.ConnectDatabase()

	// ตั้งค่าเส้นทาง API
	routes.SetupRoutes(r)

	// รันเซิร์ฟเวอร์ที่ port 8080
	r.Run(":8080")
}
