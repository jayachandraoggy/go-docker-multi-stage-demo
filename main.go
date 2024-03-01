package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	fmt.Println("env loaded")
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("hello, this is demo goapp built with docker multi-stage\n"))
		writer.Write([]byte("This is " + os.Getenv("environment") + "\n"))
	})

	fmt.Println("Server is starting at :3000")
	http.ListenAndServe(":3000", nil)
}
