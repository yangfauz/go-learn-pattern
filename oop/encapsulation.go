package oop

type Account struct {
	Email    string
	username string
	password string
}

func (a *Account) getUsernameAndPassword() (string, string) {
	return a.username, a.password
}

func (a *Account) GetEmail() string {
	return a.Email
}
