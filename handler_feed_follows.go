package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/mhdph/rss-scraper/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type paramaters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	json.NewDecoder(r.Body)

	decoder := json.NewDecoder(r.Body)

	params := paramaters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	FeedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("coulden't create feed follow: %s", err))
		return
	}
	respondWithJson(w, 201, databaseFeedFollowtoFeedsFollow(FeedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	FeedFollow, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("coulden't get feed follows: %s", err))
		return
	}
	respondWithJson(w, 201, databaseFeedFollowstoFeedsFollows(FeedFollow))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	FeedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(FeedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("coulden't parse feed follow id: %s", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("coulden't delete feed follow with id: %s", err))
		return
	}
	respondWithJson(w, 200, struct{}{})
}
