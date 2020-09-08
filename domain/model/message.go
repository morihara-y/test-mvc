package model

import "time"

// Message has message data
type Message struct {
	ID       int
	Text     string
	Datetime time.Time
}
