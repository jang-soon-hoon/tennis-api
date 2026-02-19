package database

import (
	"context"
	"log"
	"os"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var Client *firestore.Client

func init() {
	ctx := context.Background()
	var opt option.ClientOption
	// 1. Render 환경 변수에서 JSON 문자열 읽기
	configJSON := os.Getenv("FIREBASE_CONFIG_JSON")
	// if configJSON != "" {
	// 	log.Println("환경 변수 FIREBASE_CONFIG_JSON이 설정되지 않았습니다.")
	// 	// 로컬 환경: 파일이 있는지 확인 후 파일 경로로 사용
	// 	filePath := "./service-account-key.json"
	// 	if _, err := os.Stat(filePath); err == nil {
	// 		opt = option.WithCredentialsFile(filePath)
	// 		log.Println("로컬 JSON 파일에서 Firebase 설정 로드 완료")
	// 	} else {
	// 		log.Fatal("Firebase 설정(환경 변수 또는 파일)을 찾을 수 없습니다.")
	// 	}
	// } else {

	// 2. 파일 대신 JSON 바이트 배열로 Firebase 앱 초기화
	opt = option.WithCredentialsJSON([]byte(configJSON))
	//}

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Firebase 앱 초기화 실패: %v", err)
	}

	// 3. Firestore 클라이언트 생성
	var errFS error
	Client, errFS = app.Firestore(ctx)
	if errFS != nil {
		log.Fatalf("Firestore 클라이언트 생성 실패: %v", errFS)
	}
}
