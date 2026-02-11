package model

import "time"

type Event struct {
	Id int64
	Payload string
	CreatedAt time.Time
}