package main

import (
	"log"
	"github.com/tantrungse/go-ecommerce-backend-api/internal/routers"
)

func main() {
	r := routers.NewRouter()

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
