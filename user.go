package auth

type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"pass"`
}
