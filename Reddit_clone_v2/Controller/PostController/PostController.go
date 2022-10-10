package postcontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	config "reddit_clone_v2/Config"
	model "reddit_clone_v2/Model"
	"reddit_clone_v2/util"
	"strconv"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	util.SetupResponse(&w, r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "OPTIONS" {
		return
	}
	var post model.PostRequest

	_ = json.NewDecoder(r.Body).Decode(&post)

	collection := config.ConnectDB("post")

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		config.GetError(err, w)
		return
	}
	max := 1
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var temp model.PostRequest
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
	post.Id = max + 1
	result, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		config.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func VoteCount(postid int) int64 {
	collection := config.ConnectDB("vote")
	filter := bson.M{"postid": postid}

	votecount, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		fmt.Println("error in counting")
		return 0
	}
	if votecount == 0 {
		return 0
	}
	votecount = 0
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return 0
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var tmp model.Vote
		err := cur.Decode(&tmp)
		if err != nil {
			log.Fatal(err)
			return 0
		}
		votecount = votecount + int64(tmp.VoteType)
	}

	return votecount
}

func CommentCount(postid int) int64 {
	collection := config.ConnectDB("comment")
	filter := bson.M{"postid": postid}

	commentcount, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		fmt.Println("error in counting")
		return 0
	}
	return commentcount
}

func GetVoteType(postid int, username string) (upvote bool, downvote bool) {
	collection := config.ConnectDB("vote")
	filter := bson.M{"username": username, "postid": postid}
	// filter1 := bson.M{"votetype":bson.M{"$lt":0}}
	var vote model.Vote
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		fmt.Println("error in getting vote type")
	}
	if count == 0 {
		return false, false
	}
	if count > 0 {
		cur, err := collection.Find(context.TODO(), filter)
		if err != nil {
			fmt.Println("error in finding vote")
		}
		defer cur.Close(context.TODO())
		cur.Next(context.TODO())
		cur.Decode(&vote)
		if vote.VoteType == 1 {
			return true, false
		} else if vote.VoteType == -1 {
			return false, true
		}

	}
	return false, false
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	util.SetupResponse(&w, r)
	w.Header().Set("Content-Type", "application/json")

	var postresponses []model.PostResponse
	params := mux.Vars(r)
	username := params["username"]
	fmt.Println(username)
	collection := config.ConnectDB("post")

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		config.GetError(err, w)
		return
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var postrequest model.PostRequest
		var postresponse model.PostResponse
		err := cur.Decode(&postrequest)
		if err != nil {
			log.Fatal(err)

		}
		postresponse.Id = postrequest.Id
		postresponse.PostName = postrequest.PostName
		postresponse.Url = postrequest.Url
		postresponse.SubredditName = postrequest.SubredditName
		postresponse.UserName = postrequest.UserName
		postresponse.Description = postrequest.Description

		postresponse.VoteCount = int(VoteCount(postrequest.Id))
		postresponse.CommentCount = int(CommentCount(postrequest.Id))
		postresponse.Upvote, postresponse.DownVote = GetVoteType(postrequest.Id, username)

		postresponses = append(postresponses, postresponse)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(postresponses)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	util.SetupResponse(&w, r)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	intid, err := strconv.Atoi(id)
	if err != nil {
		config.GetError(err, w)
		fmt.Println("error in conversion")
		return
	}
	username := params["username"]
	fmt.Println(username)
	var post model.PostResponse

	filter := bson.M{"_id": intid}

	collection := config.ConnectDB("post")

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		config.GetError(err, w)
		return
	}
	defer cur.Close(context.TODO())
	cur.Next(context.TODO())
	// var postreq model.PostRequest
	cur.Decode(&post)
	post.VoteCount = int(VoteCount(intid))
	post.CommentCount = int(CommentCount(intid))
	post.Upvote, post.DownVote = GetVoteType(post.Id, username)
	json.NewEncoder(w).Encode(post)
}

func GetPostBySubreddit(w http.ResponseWriter, r *http.Request) {
	util.SetupResponse(&w, r)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	subreddit := params["subreddit"]

	var post []model.PostResponse

	filter := bson.M{"subredditname": subreddit}

	collection := config.ConnectDB("post")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		config.GetError(err, w)
		return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var postresponse model.PostResponse
		err := cur.Decode(&postresponse)
		if err != nil {
			log.Fatal(err)

		}

		postresponse.VoteCount = int(VoteCount(postresponse.Id))
		postresponse.CommentCount = int(CommentCount(postresponse.Id))
		postresponse.Upvote, postresponse.DownVote = GetVoteType(postresponse.Id, postresponse.UserName)

		post = append(post, postresponse)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(post)
}

func GetPostByUser(w http.ResponseWriter, r *http.Request) {
	util.SetupResponse(&w, r)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	user := params["username"]
	fmt.Println(user)

	var post []model.PostResponse

	filter := bson.M{"username": user}

	collection := config.ConnectDB("post")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		config.GetError(err, w)
		return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var postresponse model.PostResponse
		err := cur.Decode(&postresponse)
		if err != nil {
			log.Fatal(err)

		}

		postresponse.VoteCount = int(VoteCount(postresponse.Id))
		postresponse.CommentCount = int(CommentCount(postresponse.Id))
		postresponse.Upvote, postresponse.DownVote = GetVoteType(postresponse.Id, postresponse.UserName)

		post = append(post, postresponse)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(post)
}
