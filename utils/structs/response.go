package structs

import "apicars/models"

type ResponseUser struct {
	Message string      `json:"message"`
	User    models.User `json:"user"`
	Token   string      `json:"token"`
}

type RefreshTokenResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
