package commentcontroller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	config "reddit_clone_v2/Config"
	model "reddit_clone_v2/Model"

	"github.com/gorilla/mux"
)

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "OPTIONS" {
		return
	}
	var comment model.Comment

	_ = json.NewDecoder(r.Body).Decode(&comment)

	collection := config.ConnectDB("comment")

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		config.GetError(err, w)
		return
	}
	max := 1
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var temp model.Comment
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
	comment.Id = max + 1
	comment.Instant = time.Now().Format(time.RFC3339)

	result, err := collection.InsertOne(context.TODO(), comment)
	if err != nil {
		config.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func GetAllCommentsForPost(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id := params["postid"]
	intid, err := strconv.Atoi(id)
	if err != nil {
		config.GetError(err, w)
		return
	}
	var comments []model.Comment
	filter := bson.M{"postid": intid}

	collection := config.ConnectDB("comment")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		config.GetError(err, w)
		return
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var comment model.Comment
		err := cur.Decode(&comment)
		if err != nil {
			config.GetError(err, w)
		}
		comments = append(comments, comment)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(comments)
}

func GetAllCommentsByUserName(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	name := params["username"]

	var comments []model.Comment
	filter := bson.M{"username": name}

	collection := config.ConnectDB("comment")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		config.GetError(err, w)
		return
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var comment model.Comment
		err := cur.Decode(&comment)
		if err != nil {
			config.GetError(err, w)
		}
		comments = append(comments, comment)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(comments)
}
