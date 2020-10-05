package models

// SignInRequest is request model for /text API
type SignInRequest struct {
	Text string `json:"text"`
}

// SignInResponse is response model for /text API
type SignInResponse struct {
	JwtToken string `json:"jwt-token"`
}
