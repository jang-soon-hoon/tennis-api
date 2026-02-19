package main

import (
	"tennis-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// --- [CORS 설정 시작] ---
	config := cors.DefaultConfig()
	// 리액트 배포 주소(Firebase)와 로컬 주소만 허용하여 보안 강화
	config.AllowOrigins = []string{
		"https://tennis-47e2b.web.app", // 실제 Firebase 호스팅 주소로 변경하세요!
		"http://localhost:5173",        // Vite 로컬 개발 주소
	}
	config.AllowMethods = []string{"GET", "POST", "PATCH", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(config))
	// --- [CORS 설정 끝] ---

	routes.SetupRoutes(r)

	// r.GET("/matches", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World",
	// 	})
	// })

	r.Run()
}
