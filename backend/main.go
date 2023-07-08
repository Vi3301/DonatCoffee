package main

import (
	"fmt"
	"net/http"
)

func handleHttpRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, World!")
}

func main() {
	http.HandleFunc("/api/greeting", handleHttpRequest)
	http.ListenAndServe(":8080", nil)

}
