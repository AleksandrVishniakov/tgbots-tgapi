package dto

import "github.com/AleksandrVishniakov/tgbots-tgapi/models"

const GetUpdatesMethod = "getUpdates"

type GetUpdatesRequest struct {
	Offset         int      `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

type GetUpdatesResponse response[[]models.Update]
