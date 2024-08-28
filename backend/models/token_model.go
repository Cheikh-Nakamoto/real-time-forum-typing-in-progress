package models

import "time"

type TokenData struct {
	SessId     string    `json:"sessId"`
	UserId     string    `json:"userId"`
	Username string `json:"username"`
	Role       string    `json:"role"`
	SessionExp time.Time `json:"sessExp"`
	RemoteAddr string	`json:"remote_addr"`
}
