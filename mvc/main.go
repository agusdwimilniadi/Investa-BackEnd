package main

import "github.com/go-delve/delve/pkg/config"

func Migration() {
	config.DB.AutoMigrate()
}

type User struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
