package router

import (
	commentcontroller "reddit_clone_v2/Controller/CommentController"
	postcontroller "reddit_clone_v2/Controller/PostController"
	signcontoller "reddit_clone_v2/Controller/SignupContoller"
	redditcontroller "reddit_clone_v2/Controller/SubredditController"
	votecontroller "reddit_clone_v2/Controller/VoteController"
	middleware "reddit_clone_v2/Middleware"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	apiRouter := mux.NewRouter()
	SignHandler(apiRouter)
	SubredditHandler(apiRouter)
	PostHandler(apiRouter)
	CommentHandler(apiRouter)
	VoteHandler(apiRouter)
	TokenHandler(apiRouter)
	return apiRouter
}

func SignHandler(r *mux.Router) {
	r.HandleFunc("/signup", signcontoller.Signup).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", signcontoller.Login).Methods("POST", "OPTIONS")
}

func SubredditHandler(r *mux.Router) {
	r.HandleFunc("/subreddit", redditcontroller.CreateSubreddit).Methods("POST", "OPTIONS")
	r.HandleFunc("/subreddit", redditcontroller.GetAllSubreddit).Methods("GET")
	r.HandleFunc("/subreddit/{id}", redditcontroller.GetSubredditById).Methods("GET")
}

func PostHandler(r *mux.Router) {
	r.HandleFunc("/post/by-user/{username}", postcontroller.GetPostByUser).Methods("GET")
	r.HandleFunc("/post/by-subreddit/{subreddit}", postcontroller.GetPostBySubreddit).Methods("GET")
	r.HandleFunc("/post", postcontroller.CreatePost).Methods("POST", "OPTIONS")
	r.HandleFunc("/post-all/{username}", postcontroller.GetAllPosts).Methods("GET")
	r.HandleFunc("/post/{id}/{username}", postcontroller.GetPostById).Methods("GET")
}

func CommentHandler(r *mux.Router) {
	r.HandleFunc("/comment", commentcontroller.CreateComment).Methods("POST", "OPTIONS")
	r.HandleFunc("/comment/by-post/{postid}", commentcontroller.GetAllCommentsForPost).Methods("GET")
	r.HandleFunc("/comment/by-user/{username}", commentcontroller.GetAllCommentsByUserName).Methods("GET")
}

func VoteHandler(r *mux.Router) {
	r.HandleFunc("/vote", votecontroller.ManageVote).Methods("POST", "OPTIONS")
}

func TokenHandler(r *mux.Router) {
	r.HandleFunc("/refresh", middleware.RefreshToken).Methods("POST", "OPTIONS")
}
