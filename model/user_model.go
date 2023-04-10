package model

type User struct {
	ID       int
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
