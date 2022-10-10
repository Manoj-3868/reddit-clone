package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	model "reddit_clone_v2/Model"
	"reddit_clone_v2/util"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func VerifyToken(w http.ResponseWriter, r *http.Request, token string) (err error) {
	err = nil
	// c, err := r.Cookie("token")
	// if err != nil {
	// 	return err
	// }

	tknStr := token

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return err
	}
	if !tkn.Valid {
		err = errors.New("token not valid")
		return err
	}
	return err
}

func GenerateToken(w http.ResponseWriter, r *http.Request, signup model.Signup) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: signup.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	return tokenString, nil
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	util.SetupResponse(&w, r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "OPTIONS" {
		return
	}

	claims := &Claims{}
	var refreshToken model.RefreshToken
	_ = json.NewDecoder(r.Body).Decode(&refreshToken)
	fmt.Println(refreshToken)
	err1 := VerifyToken(w, r, refreshToken.Token)
	if err1 != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err1.Error()))
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	refreshToken.Token = tokenString

	json.NewEncoder(w).Encode(refreshToken)
}
