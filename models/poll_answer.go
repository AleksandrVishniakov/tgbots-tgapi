package models

type PollAnswer struct {
	PollID    string `json:"poll_id"`
	User      User   `json:"user"`
	OptionIDs []int  `json:"option_ids"`
}
