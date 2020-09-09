package model

import "time"

// MessageTrn has message data
type MessageTrn struct {
	ID       int       `json:"id"`
	Text     string    `json:"text"`
	Datetime time.Time `json:"datetime"`
}
