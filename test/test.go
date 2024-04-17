package main

import "github.com/nusa-exchange/pkg/services"

func main() {
	logger := services.NewLoggerService("Finex")

	logger.Info("This is an info message")
}
