package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func httpRegisterUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	hasedPassword := bcrypt
	fmt.Println("Hello, World!")
}

func main() {
	dbUser := os.Getenv("DATABASEUSER")
	dbName := os.Getenv("DATABASENAME")
	dbPass := os.Getenv("DATABASEPASS")
	db, err := sql.Open("mysql", fmt.Sprintf(`%s:%s/@%s`, dbUser, dbPass, dbName))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
	http.HandleFunc("/api/greeting", httpRegisterUser)
	http.ListenAndServe(":8080", nil)

}
