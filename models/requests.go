package models

// SignInRequest is request model for /text API
type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
