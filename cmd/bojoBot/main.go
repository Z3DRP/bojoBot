package main

import (
	"fmt"
	"github.com/Z3DRP/bojoBot/internal/routes"
	"net/http"
	"time"
)

func main() {
	router := routes.NewRouter()
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)

	if err != nil {
		panic(err)
	}
}
