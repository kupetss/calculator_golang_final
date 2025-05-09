package auth

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"-"`
}

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
