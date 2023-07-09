package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func handleHttpRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, World!")
}

func main() {
	http.HandleFunc("/api/greeting", handleHttpRequest)
	http.ListenAndServe(":8080", nil)

}
