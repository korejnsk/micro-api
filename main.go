package main

import (
	"context"
	"fmt"

	"github.com/korejnsk/micro-api/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Erro ao iniciar o aplicativo:", err)
	}
}
