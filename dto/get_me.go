package dto

import "github.com/AleksandrVishniakov/tgbots-tgapi/models"

const GetMeMethod = "getMe"

type GetMeResponse response[*models.User]
