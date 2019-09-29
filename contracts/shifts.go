package contracts

import "time"

type ShiftBO struct{
	Id uint64 `json:"id"`
	Worker string `json:"worker"`
	StartedAt time.Time `json:"startedat"`
	EndAt time.Time `json:"endat"`
}
