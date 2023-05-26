package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "HELLO!")
	})

	port := os.Getenv("SERVER_PORT")
	fmt.Printf("Starting server at post %v\n", port)
	formattedPort := fmt.Sprintf(":%v", port)
	if err := http.ListenAndServe(formattedPort, nil); err != nil {
		log.Fatal(err)
	}
}
