package models

import "time"

type Event struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id"`
	EventName string    `json:"event_name"`
	EventDate time.Time `json:"event_date"`
}
