package main

import (
	"gobarber/cmd/app"
	"log"
)

func main() {
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
