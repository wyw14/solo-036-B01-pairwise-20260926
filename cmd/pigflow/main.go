package main

import (
	"context"
	"github.com/wyw14/solo-036/internal/api"
	"github.com/wyw14/solo-036/internal/service"
	"log"
	"net/http"
	"os"
)

func main() {
	runtime := service.NewRuntime()
	if dataDirectory := os.Getenv("PIGFLOW_DATA_DIR"); dataDirectory != "" {
		if err := runtime.ConfigurePersistence(dataDirectory); err != nil {
			log.Fatal(err)
		}
		if os.Getenv("PIGFLOW_RECOVER") == "1" {
			if _, err := runtime.RecoverPersistent(); err != nil {
				log.Fatal(err)
			}
		}
	}
	runtime.Bootstrap()
	address := os.Getenv("PIGFLOW_ADDR")
	if address == "" {
		address = "127.0.0.1:21243"
	}
	server := &http.Server{Addr: address, Handler: api.NewRouter(runtime)}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	_ = runtime.Shutdown(context.Background())
}
