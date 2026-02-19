package main

import (
	"context"
	"log"
	"os"

	firestore "cloud.google.com/go/firestore/apiv1"
	firebase "firebase.google.com/go/v4"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

var db *firestore.Client

func init() {
	ctx := context.Background()

	// 1. Render 환경 변수에서 JSON 문자열 읽기
	configJSON := os.Getenv("FIREBASE_CONFIG_JSON")
	if configJSON == "" {
		log.Fatal("환경 변수 FIREBASE_CONFIG_JSON이 설정되지 않았습니다.")
	}

	// 2. 파일 대신 JSON 바이트 배열로 Firebase 앱 초기화
	opt := option.WithCredentialsJSON([]byte(configJSON))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Firebase 앱 초기화 실패: %v", err)
	}

	// 3. Firestore 클라이언트 생성
	var errFS error
	db, errFS = app.Firestore(ctx)
	if errFS != nil {
		log.Fatalf("Firestore 클라이언트 생성 실패: %v", errFS)
	}
}

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

	r.GET("/matches", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.Run()
}
