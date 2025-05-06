package types

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
