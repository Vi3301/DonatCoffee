package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *sql.DB

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	hasedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, string(hasedPassword))
	if err != nil {
		http.Error(w, "Unable, to register user", http.StatusBadRequest)
		return
	}
	w.Write([]byte("User registered successfully"))
}
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	var dbUser User
	err := db.QueryRow("SELECT * FROM users WHERE username = ?", user.Username).Scan(&dbUser.Username, dbUser.Password)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		http.Error(w, "Bad password", http.StatusBadRequest)
		return
	}
	tokenString, err := createToken(user)
	if err != nil {
		http.Error(w, "Error creating token", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(tokenString))
	w.Write([]byte("Logged is successfully"))
}

var jwtKey = []byte("Secret_key")

func createToken(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return " ", err
	}
	return tokenString, nil
}
func validToken(tokenString string) (string, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) { return jwtKey, nil })
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("invalid Token")
	}
	username := claims["username"].(string)
	return username, nil
}
func protect(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	_, err := validToken(tokenString)
	if err != nil {
		http.Error(w, "Invalid Token", http.StatusUnauthorized)
		return
	}
}
func main() {
	dbUser := os.Getenv("DATABASEUSER")
	dbName := os.Getenv("DATABASENAME")
	dbPass := os.Getenv("DATABASEPASS")
	db, err := sql.Open("mysql", fmt.Sprintf(`%s:%s/@%s`, dbUser, dbPass, dbName))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()
	http.HandleFunc("/register", RegisterUser)
	http.HandleFunc("/logun", LoginUser)
	http.ListenAndServe(":8080", nil)
}
