package dto

import "time"

type UserResponse struct {
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"createdAt"`
}
