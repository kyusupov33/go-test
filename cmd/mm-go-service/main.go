package main

import (
	"github.com/kyusupov33/mm-go-service/internal/app"
)

func main() {
	application := app.New()

	application.Start()
}
