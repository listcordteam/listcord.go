package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type ErrorObject struct {
	Message interface{} `json:"message"`
}

type Bot struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Description struct {
		Short string `json:"short"`
		Long  string `json:"long"`
	} `json:"description"`
	Developers    []string    `json:"developers"`
	Permissions   int32       `json:"required_permissions"`
	Upvotes       int32       `json:"upvotes"`
	SupportServer string      `json:"support_server"`
	Website       string      `json:"website"`
	Tags          []string    `json:"tags"`
	Prefix        string      `json:"prefix"`
	SubmittedAt   interface{} `json:"submitted_at"`
	Approved      bool        `json:"approved"`
}

type BotReview struct {
	AuthorID string `json:"author_id"`
	Stars    int    `json:"stars"`
	Content  string `json:"content"`
	SentAt   int64  `json:"sent_at"`
}

type VoteData struct {
	Voted        bool  `json:"voted"`
	UpvotedAt    int64 `json:"upvoted_at"`
	NextUpvoteAt int64 `json:"next_at"`
}

type FetchOptions struct {
	url    string
	params string
}

type ClientConstructor struct {
	Token   string
	baseURL string
}

func Client(token string) ClientConstructor {

	return ClientConstructor{
		Token:   token,
		baseURL: "https://listcord.xyz/api",
	}

}

var ErrorMessages map[int]string = map[int]string{
	401: "401! Invalid token!",
	404: "404! Not Found!",
	429: "429! You have been rate limited by the api!",
	500: "500! Something went wrong in the api server!",
}

func (self ClientConstructor) Fetch(options FetchOptions, structure interface{}) error {
	reqClient := &http.Client{}

	req, err := http.NewRequest("GET", (self.baseURL + options.url + "?" + options.params), nil)
	req.Header.Add("token", self.Token)

	res, err := reqClient.Do(req)
	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		return readErr
	}

	if err != nil {
		return errors.New("UnexpectedError: Request failed while accessing Listcord Api!")
	}

	if res.StatusCode != 200 {
		errorMsg, found := ErrorMessages[res.StatusCode]

		if found {
			return errors.New("ListcordApiError: " + errorMsg)
		}
	}

	jsonErr := json.Unmarshal(body, structure)

	if jsonErr != nil {
		return jsonErr
	}

	return nil
}

func (self ClientConstructor) GetBot(ID string) (Bot, error) {
	var bot Bot
	err := self.Fetch(FetchOptions{"/bot/" + ID, ""}, &bot)
	return bot, err
}

func (self ClientConstructor) GetBotReviews(ID string) ([]BotReview, error) {
	var reviews []BotReview
	err := self.Fetch(FetchOptions{"/bot/" + ID + "/reviews", ""}, &reviews)
	return reviews, err
}

func (self ClientConstructor) GetReview(UserID string, BotID string) (BotReview, bool) {
	reviews, err := self.GetBotReviews(BotID)

	if err != nil {
		return BotReview{}, false
	}

	for _, review := range reviews {
		if review.AuthorID == UserID {
			return review, true
		}
	}

	return BotReview{}, false
}

func (self ClientConstructor) Search(Query string) ([]Bot, error) {
	var bots []Bot
	err := self.Fetch(FetchOptions{"/bots", "q=" + Query}, &bots)
	return bots, err
}

func (self ClientConstructor) HasVoted(UserID string, BotID string) (VoteData, error) {
	var voteData VoteData
	err := self.Fetch(FetchOptions{"/bot/" + BotID + "/voted", "user_id=" + UserID}, &voteData)
	return voteData, err
}
