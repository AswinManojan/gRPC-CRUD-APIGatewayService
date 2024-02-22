package models

type Provider struct {
	Username string
	Password string
}

type ProviderSignUp struct {
	Username string
	Password string
	Email    string
}
