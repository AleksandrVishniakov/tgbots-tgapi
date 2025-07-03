package dto

import "github.com/AleksandrVishniakov/tgbots-tgapi/models"

const SetMyShortDescriptionMethod = "setMyShortDescription"

type SetMyShortDescriptionRequest struct {
	ShortDescription string              `json:"short_description,omitempty"`
	LanguageCode     models.LanguageCode `json:"language_code,omitempty"`
}

type SetMyShortDescriptionResponse response[bool]
