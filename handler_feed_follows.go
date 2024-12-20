package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/ringtho/rssagg/internal/database"
)


func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	 type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	 }

	 decoder := json.NewDecoder(r.Body)
	 params := parameters{}
	 err := decoder.Decode(&params)
	 if err != nil {
		respondWithErrors(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	 }

	 feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:       	uuid.New(),
		CreatedAt: 	time.Now().UTC(),
		UpdatedAt: 	time.Now().UTC(),
		UserID:   	user.ID,
		FeedID:     params.FeedID,
	 }) 

	if err != nil {
		respondWithErrors(w, 400, fmt.Sprintf("Couldn't create a feed follow: %v", err))
		return
	}
	respondWithJson(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollow, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID) 

	if err != nil {
		respondWithErrors(w, 400, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}
	respondWithJson(w, 200, databaseFeedFollowsToFeedFollows(feedFollow))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollowIDstr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDstr)

	if err != nil {
		respondWithErrors(w, 400, fmt.Sprintf("Couldn't parse feed follow: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID: feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		respondWithErrors(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}

	respondWithJson(w, 200, struct{}{})
}