package models

type UserProfile struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
