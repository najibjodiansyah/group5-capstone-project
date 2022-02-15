package auth

type LoginRequestFormat struct {
	Email string
	Password string
}
type LoginResponseFormat struct {
	Id int
	Name string
	Token string
}