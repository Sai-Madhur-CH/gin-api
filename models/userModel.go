package models

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	tokenSecret = []byte(os.Getenv("TOKEN_SECRET"))
)

type User struct {
	User_id 		int 	`json:"user_id"`
	Name 			string 	`json:"name"`
	Email 			string 	`json:"email"`
	Password_hash 	string 	`json:"_"`
	Password        string  `json:"password"`
	CreatedAt       time.Time `json:"_"`
	UpdatedAt       time.Time `json:"_"`
}


func (u *User) Register(conn *pgx.Conn) error {
	fmt.Println(&u)
	if len(u.Email) < 4 {
		return fmt.Errorf("Email must be at least 4 characters long.")
	}
	if len(u.Name) < 4 {
		return fmt.Errorf("name must be at least 4 characters long.")
	}
	if len(u.Password) < 4 {
		return fmt.Errorf("password must be at least 4 characters long.")
	}
	

	u.Email = strings.ToLower(u.Email)
	row := conn.QueryRow(context.Background(), "SELECT user_id from config.user_details WHERE email = $1", u.Email)
	userLookup := User{}
	err := row.Scan(&userLookup)
	if err != pgx.ErrNoRows {
		fmt.Println("found user")
		fmt.Println(userLookup.Email)
		return fmt.Errorf("A user with that email already exists")
	}

	pwdHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("There was an error creating your account.")
	}
	u.Password_hash = string(pwdHash)

	now := time.Now()

	_, err = conn.Exec(context.Background(), "INSERT INTO config.user_details (created_at, updated_at, email, password_hash, name) VALUES($1, $2, $3, $4, $5)", now, now, u.Email, u.Password_hash, u.Name)

	return err
}


// GetAuthToken returns the auth token to be used
func (u *User) GetAuthToken() (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = u.User_id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	authToken, err := token.SignedString(tokenSecret)
	return authToken, err
}