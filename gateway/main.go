package main

import (
	_ "github.com/joho/godotenv/autoload"
	common "github.com/kweheliye/omsv2/common"
	"log"
	"net/http"
)

var httpAddr = common.EnvString("HTTP_ADDR", ":3000")

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Println("Starting server on " + httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
