package controllers

import (
	"context"
	"net/http"
	"tennis-api/database"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func GetMatches(c *gin.Context) {
	ctx := context.Background()
	var matches []map[string]interface{}

	// Firestore "matches" 컬렉션 데이터 조회
	iter := database.Client.Collection("matches").Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "데이터 조회 중 오류 발생"})
			return
		}

		data := doc.Data()
		data["id"] = doc.Ref.ID // 문서 고유 ID 포함
		matches = append(matches, data)
	}

	c.JSON(http.StatusOK, matches)
}
