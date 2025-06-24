package models

type Message struct {
	MessageID int    `json:"message_id"`
	From      *User  `json:"from,omitempty"`
	Date      int    `json:"date"`
	Chat      Chat   `json:"chat"`
	Text      string `json:"text,omitempty"`
}
