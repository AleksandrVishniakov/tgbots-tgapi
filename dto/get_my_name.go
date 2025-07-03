package dto

import "github.com/AleksandrVishniakov/tgbots-tgapi/models"

const GetMyNameMethod = "getMyName"

type GetMyNameRequest struct {
	LanguageCode models.LanguageCode `json:"language_code,omitempty"`
}

type GetMyNameResponse response[string]
