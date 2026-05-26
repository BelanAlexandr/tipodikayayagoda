package models

import "time"

type Notification struct {
	ID         int       `json:"id"`
	User_ID    int       `json:"user_id"`
	Text       string    `json:"text"`
	Is_read    bool      `json:"is_read"`
	Created_at time.Time `json:"created_at"`
}
