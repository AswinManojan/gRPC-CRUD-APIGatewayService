package models

type User struct {
	Username string
	Password string
}

type UserSignUp struct {
	Username string
	Password string
	Email    string
}
