package dto

import "remainderTask/models"

type DefaultRespDto struct {
	Status           string `json:"status"`
	ErrorDescription string `json:"errorDescription"`
}

type UserDetailsRespDto struct {
	Users            []models.User `json:"users"`
	Status           string        `json:"status"`
	ErrorDescription string        `json:"errorDescription"`
}
