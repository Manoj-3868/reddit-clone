package model

type Signup struct {
	Id       int    `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName string `json:"username,omitempty" bson:"username,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	Token    string `json:"token,omitempty"`
}

type Login struct {
	UserName string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

type Subreddit struct {
	Id          int    `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Postcount   int    `json:"numberofpost,omitempty" bson:"numberofpost,omitempty"`
}

type PostRequest struct {
	Id            int    `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName      string `json:"username,omitempty" bson:"username,omitempty"`
	SubredditName string `json:"subredditname,omitempty" bson:"subredditname,omitempty"`
	PostName      string `json:"postname,omitempty" bson:"postname,omitempty"`
	Url           string `json:"url,omitempty" bson:"url,omitempty"`
	Description   string `json:"description,omitempty" bson:"description,omitempty"`
}

type PostResponse struct {
	Id            int    `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName      string `json:"username,omitempty" bson:"username,omitempty"`
	SubredditName string `json:"subredditname,omitempty" bson:"subredditname,omitempty"`
	PostName      string `json:"postname,omitempty" bson:"postname,omitempty"`
	Url           string `json:"url,omitempty" bson:"url,omitempty"`
	Description   string `json:"description,omitempty" bson:"description,omitempty"`
	VoteCount     int    `json:"votecount" bson:"votecount,omitempty"`
	CommentCount  int    `json:"commentcount" bson:"commentcount,omitempty"`
	Upvote        bool   `json:"upvote" bson:"upvote,omitempty"`
	DownVote      bool   `json:"downvote" bson:"downvote,omitempty"`
}

type Comment struct {
	Id       int    `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName string `json:"username,omitempty" bson:"username,omitempty"`
	PostId   int    `json:"postid,omitempty" bson:"postid,omitempty"`
	Instant  string `json:"instant,omitempty" bson:"instant,omitempty"`
	Text     string `json:"text,omitempty" bson:"text,omitempty"`
}

type Vote struct {
	Id       int    `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName string `json:"username,omitempty" bson:"username,omitempty"`
	PostId   int    `json:"postid,omitempty" bson:"postid,omitempty"`
	VoteType int    `json:"votetype,omitempty" bson:"votetype,omitempty"`
	Token    string `json:"token,omitempty"`
}

type RefreshToken struct {
	Token string `json:"token,omitempty"`
}
