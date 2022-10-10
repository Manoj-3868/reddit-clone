package subredditcontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"

	config "reddit_clone_v2/Config"
	model "reddit_clone_v2/Model"
)

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func postCount(name string) int {
	collection := config.ConnectDB("post")
	filter := bson.M{"subredditname": name}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0
	}
	return int(count)
}

func CreateSubreddit(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "OPTIONS" {
		return
	}

	var subreddit model.Subreddit

	collection := config.ConnectDB("subreddit")
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		config.GetError(err, w)
		return
	}
	max := 1
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var temp model.Subreddit
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
	_ = json.NewDecoder(r.Body).Decode(&subreddit)
	subreddit.Id = max + 1
	fmt.Println(subreddit)

	result, err := collection.InsertOne(context.TODO(), subreddit)

	if err != nil {
		config.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func GetAllSubreddit(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	w.Header().Set("Content-Type", "application/json")

	var subreddits []model.Subreddit

	collection := config.ConnectDB("subreddit")

	filter := bson.M{}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		config.GetError(err, w)
		return
	}
	for cur.Next(context.TODO()) {
		var subreddit model.Subreddit
		err := cur.Decode(&subreddit)
		if err != nil {
			config.GetError(err, w)
		}
		subreddit.Postcount = postCount(subreddit.Name)
		subreddits = append(subreddits, subreddit)
	}
	defer cur.Close(context.TODO())

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(subreddits)
}

func GetSubredditById(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	id := params["id"]
	intid, err := strconv.Atoi(id)
	if err != nil {
		config.GetError(err, w)
		return
	}

	filter := bson.M{"_id": intid}

	var subreddit model.Subreddit

	collection := config.ConnectDB("subreddit")

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		config.GetError(err, w)
		return
	}
	cur.Next(context.TODO())
	defer cur.Close(context.TODO())
	cur.Decode(&subreddit)

	json.NewEncoder(w).Encode(subreddit)
}
