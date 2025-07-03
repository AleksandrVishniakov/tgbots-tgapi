package dto

import "github.com/AleksandrVishniakov/tgbots-tgapi/models"

const SetMyDescriptionMethod = "setMyDescription"

type SetMyDescriptionRequest struct {
	Description  string              `json:"description,omitempty"`
	LanguageCode models.LanguageCode `json:"language_code,omitempty"`
}

type SetMyDescriptionResponse response[bool]
