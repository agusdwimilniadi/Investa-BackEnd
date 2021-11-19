package main

import (
	configs "investaBackend/configs"
	routes "investaBackend/routes"
)

func main() {
	configs.ConnectDB()

	e := routes.New()
	e.Start(":8000")
}
