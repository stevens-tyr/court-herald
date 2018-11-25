package main

import (
	"court-herald/api"
)

func main() {
	server := api.SetUp()
	server.Run(":4444")
}
