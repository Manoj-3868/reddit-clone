package votecontroller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	config "reddit_clone_v2/Config"
	middleware "reddit_clone_v2/Middleware"
	model "reddit_clone_v2/Model"
	"reddit_clone_v2/util"

	bson "go.mongodb.org/mongo-driver/bson"
)

func ManageVote(w http.ResponseWriter, r *http.Request) {
	util.SetupResponse(&w, r)
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "OPTIONS" {
		return
	}
	var vote model.Vote

	_ = json.NewDecoder(r.Body).Decode(&vote)
	// authorization start
	err := middleware.VerifyToken(w, r, vote.Token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	// authorization end

	filter := bson.M{"username": vote.UserName, "postid": vote.PostId}

	collection := config.ConnectDB("vote")

	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		config.GetError(err, w)
		return
	}
	// fmt.Println(count)
	if count > 0 {
		var vote1 model.Vote
		cur, err := collection.Find(context.TODO(), filter)
		if err != nil {
			config.GetError(err, w)
			return
		}
		defer cur.Close(context.TODO())
		cur.Next(context.TODO())
		cur.Decode(&vote1)
		if vote.VoteType == vote1.VoteType {
			result, err := collection.DeleteOne(context.TODO(), filter)
			if err != nil {
				config.GetError(err, w)
				return
			}
			json.NewEncoder(w).Encode(result)
		} else {
			update := bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "votetype", Value: vote.VoteType},
				},
				},
			}

			result, err := collection.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				config.GetError(err, w)
				return
			}
			json.NewEncoder(w).Encode(result)
		}
	} else {
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
		vote.Id = max + 1
		vote.Token = ""
		result, err := collection.InsertOne(context.TODO(), vote)
		if err != nil {
			config.GetError(err, w)
			return
		}

		json.NewEncoder(w).Encode(result)
	}
}
