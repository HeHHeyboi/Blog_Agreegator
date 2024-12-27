package main

import (
	"blog_agreegator/internal/database"
	myerror "blog_agreegator/internal/myError"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func scrapeFeeds(s *state, ctx context.Context) (database.Feed, error) {
	nextFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return nextFeed, fmt.Errorf("Error Get feed: %v", err)
	}

	err = MarkFeed(s, ctx, nextFeed.ID)
	if err != nil {
		return nextFeed, err
	}

	rss, err := fetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return nextFeed, fmt.Errorf("Error Fetch feed: %v", err)
	}
	fmt.Println("Feed Item:")
	for _, item := range rss.Channel.Item {
		post, err := createPost(s, ctx, item, nextFeed)
		if err != nil {
			return nextFeed, err
		}
		if !post.Title.Valid {
			continue
		}
		fmt.Println("Title:", post.Title.String)
		fmt.Println("Url:", post.Url)
	}

	return nextFeed, nil
}
func createPost(s *state, ctx context.Context, item RSSItem, feed database.Feed) (database.Post, error) {
	// FIX: Test Multiple Format from Rss, if get error try different format
	// Might need to create Array of time.Format

	publish, err := time.Parse(time.RFC1123, item.PubDate)
	if err != nil {
		return database.Post{}, fmt.Errorf("Parse Time error: %v", err)
	}
	post, err := s.db.CreatePost(ctx, database.CreatePostParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title: sql.NullString{
			String: item.Title,
			Valid:  item.Title != "",
		},
		Url: item.Link,
		Description: sql.NullString{
			String: item.Description,
			Valid:  item.Description != "",
		},
		PublishedAt: publish,
		FeedID:      feed.ID,
	})
	if errors.Is(err, &myerror.ErrDuplicate{}) {
		return database.Post{}, nil
	}
	return post, nil
}

func MarkFeed(s *state, ctx context.Context, id uuid.UUID) error {
	err := s.db.MarkFeedFetched(ctx, id)
	if err != nil {
		return fmt.Errorf("Error mark feed: %v", err)
	}
	err = s.db.MarkFollowFetched(ctx, id)
	if err != nil {
		return fmt.Errorf("Error mark follow: %v", err)
	}
	return nil
}
