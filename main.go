package main

import (
	"fmt"
	"github.com/momenteam/momentum/configs"
	"github.com/momenteam/momentum/database"
)

func init() {
	configs.Setup()
	database.Setup()
}

func main() {
	client := database.Client

	if client == nil {
		fmt.Println("error")
	}
}