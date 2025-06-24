package models

type Poll struct {
	ID              string       `json:"id"`
	Question        string       `json:"question"`
	Options         []PollOption `json:"options"`
	IsClosed        bool         `json:"is_closed"`
	TotalVoterCount int          `json:"total_voter_count"`
}
