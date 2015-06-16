package main


// Models
type User struct {
	Username string
	Password string
}

type Like struct {
	Sender string
	Receiver string
}