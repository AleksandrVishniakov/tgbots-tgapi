package models

type ChosenInlineResult struct {
	ResultID        string `json:"result_id"`
	From            User   `json:"from"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
	Query           string `json:"query"`
}
