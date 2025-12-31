package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction() // creates a logger instance with Info level by default and json formatting
	if err != nil {
		log.Println("Error :", err)
	}
	defer logger.Sync() // flushes log buffer, if any

	logger.Info("This is an info message")                                                                                            // logs at Info level
	logger.Info("This is another info message with few extra field", zap.String("user", "Rishav Kumar"), zap.String("role", "Admin")) // logs at Info level with extra fields
}

// The output will be in JSON format like below :

// {"level":"info","ts":1767170845.2499256,"caller":"002_intermediate/zap.go:16","msg":"This is an info message"}
// {"level":"info","ts":1767170845.2499914,"caller":"002_intermediate/zap.go:17","msg":"This is another info message with few extra field","user":"Rishav Kumar","role":"Admin"}
