package main

func Migration() {
	DB.AutoMigrate()
}

type User struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
