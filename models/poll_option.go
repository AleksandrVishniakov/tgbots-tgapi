package models

type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}
