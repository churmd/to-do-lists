package models

import "time"

type UserID string

type Todo struct {
	ID          string
	UserID      UserID
	Description string
	Created     time.Time
	Complete    bool
}
