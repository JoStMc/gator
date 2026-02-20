package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/JoStMc/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
	    return fmt.Errorf("insufficient number of commands")
	} 

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
	    return err
	} 
	if timeBetweenRequests < time.Second {
		fmt.Println("Time between requests too small. Minimum: 1 second")
		return nil
	} 
	fmt.Println("Collecting feeds every ", timeBetweenRequests)
	fmt.Println()
	ticker := time.NewTicker(timeBetweenRequests)

	for ;; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
		    return err
		} 
	} 
} 

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	nextFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
	    return err
	} 

	params := database.MarkFeedFetchedParams{
		ID: nextFeed.ID, 
		LastFetchedAt: sql.NullTime{
			Time: time.Now(),
			Valid: true,
		},
	}
	err = s.db.MarkFeedFetched(ctx, params)
	if err != nil {
	    return err
	} 

	feed, err := fetchFeed(ctx, nextFeed.Url)
	if err != nil {
	    return err
	} 

	for _, item := range feed.Channel.Item {
		post, _ := s.db.GetPost(ctx, item.Link)
		if post != (database.Post{}) {
		    continue
		} 
		postParams := database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: item.Title,
			Url: item.Link,
			Description: item.Description,
			PublishedAt: item.PubDate,
			FeedID: nextFeed.ID,
		} 
		err = s.db.CreatePost(ctx, postParams)
		if err != nil {
		    return err
		} 
		fmt.Printf("Post %s saved for feed %s\n", item.Title, nextFeed.Name)
	}
	fmt.Println()

	return nil
} 
