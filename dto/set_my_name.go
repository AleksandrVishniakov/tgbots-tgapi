package dto

import "github.com/AleksandrVishniakov/tgbots-tgapi/models"

const SetMyNameMethod = "setMyName"

type SetMyNameRequest struct {
	Name         string              `json:"name,omitempty"`
	LanguageCode models.LanguageCode `json:"languageCode,omitempty"`
}

type SetMyNameResponse response[bool]
