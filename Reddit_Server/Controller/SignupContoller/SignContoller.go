package signupcontoller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	config "reddit_clone_v2/Config"
	middleware "reddit_clone_v2/Middleware"
	model "reddit_clone_v2/Model"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func Signup(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		return
	}
	var signup model.Signup

	_ = json.NewDecoder(r.Body).Decode(&signup)
	collection := config.ConnectDB("user")
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		config.GetError(err, w)
		return
	}
	max := 1
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var temp model.Signup
		err := cur.Decode(&temp)
		if err != nil {
			log.Fatal(err)
		}
		if max < temp.Id {
			max = temp.Id
		}
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	signup.Id = max + 1
	result, err := collection.InsertOne(context.TODO(), signup)
	if err != nil {
		config.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func Login(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	w.Header().Set("Content-Type", "application/json")

	var signup model.Signup
	var userLogin model.Login

	if r.Method == "OPTIONS" {
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&userLogin)
	filter := bson.M{"username": userLogin.UserName}
	collection := config.ConnectDB("user")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		config.GetError(err, w)
		return
	}
	defer cur.Close(context.TODO())
	cur.Next(context.TODO())
	cur.Decode(&signup)
	if userLogin.Password != signup.Password {
		config.GetError(err, w)
		return
	}
	token, err1 := middleware.GenerateToken(w, r, signup)
	if err1 != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err1.Error()))
		return
	}
	fmt.Println("loggedin", token)
	signup.Token = token
	json.NewEncoder(w).Encode(signup)
}
