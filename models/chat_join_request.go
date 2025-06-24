package models

type ChatJoinRequest struct {
	Chat Chat `json:"chat"`
	From User `json:"from"`
	Date int    `json:"date"`
	Bio  string `json:"bio,omitempty"`
}
