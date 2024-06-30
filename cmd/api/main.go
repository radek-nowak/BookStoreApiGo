package main

import (
	"book_store_api_go/api"
)

func main() {
	app := api.NewApp()

	app.Listen(":3000")
}
