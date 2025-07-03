package dto

import "github.com/AleksandrVishniakov/tgbots-tgapi/models"

const GetMyShortDescriptionMethod = "getMyShortDescription"

type GetMyShortDescriptionRequest struct {
	LanguageCode models.LanguageCode `json:"language_code,omitempty"`
}

type GetMyShortDescriptionResponse response[string]
