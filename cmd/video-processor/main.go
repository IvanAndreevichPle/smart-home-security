// cmd/video-processor/main.go
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Настройка логгера
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Канал для грациозного завершения
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Ожидание сигнала завершения
	<-sigChan
	log.Println("Shutting down...")
}
