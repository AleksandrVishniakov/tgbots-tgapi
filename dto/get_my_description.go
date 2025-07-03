package dto

import "github.com/AleksandrVishniakov/tgbots-tgapi/models"

const GetMyDescriptionMethod = "getMyDescription"

type GetMyDescriptionRequest struct {
	LanguageCode models.LanguageCode `json:"language_code,omitempty"`
}

type GetMyDescriptionResponse response[string]
