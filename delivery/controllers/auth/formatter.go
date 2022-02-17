package auth

type LoginRequestFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type LoginResponseFormat struct {
	Id    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Token string `json:"token" form:"token"`
}
