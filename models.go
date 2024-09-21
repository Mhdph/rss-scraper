package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/mhdph/rss-scraper/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}
func databaseFeedtoFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}
func databaseFeedtoFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedtoFeed(dbFeed))
	}

	return feeds

}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedId    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedFollowtoFeedsFollow(dbFeedsFollow database.FeedFollow) FeedFollow {

	return FeedFollow{
		ID:        dbFeedsFollow.ID,
		CreatedAt: dbFeedsFollow.CreatedAt,
		UpdatedAt: dbFeedsFollow.UpdatedAt,
		FeedId:    dbFeedsFollow.FeedID,
		UserID:    dbFeedsFollow.UserID,
	}

}

func databaseFeedFollowstoFeedsFollows(dbFeedsFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}

	for _, dbFeedsFollows := range dbFeedsFollows {
		feedFollows = append(feedFollows, databaseFeedFollowtoFeedsFollow(dbFeedsFollows))
	}
	return feedFollows

}
