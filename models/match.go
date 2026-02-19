package models

import "time"

type Match struct {
	ID        string    `json:"id" firestore:"-"` // ID는 문서 이름에서 가져오므로 생략
	PlayerA   string    `json:"playerA" firestore:"PlayerA"`
	PlayerB   string    `json:"playerB" firestore:"PlayerB"`
	ScoreA    int       `json:"scoreA" firestore:"scoreA"`
	ScoreB    int       `json:"scoreB" firestore:"scoreB"`
	Status    string    `json:"status" firestore:"status"` // "scheduled", "live", "finished"
	CreatedAt time.Time `json:"createdAt" firestore:"createdAt"`
}
