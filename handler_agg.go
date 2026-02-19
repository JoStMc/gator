package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/JoStMc/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
	    return fmt.Errorf("insufficient number of commands")
	} 

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
	    return err
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

	fmt.Println("Feed: ", feed.Channel.Title)
	for _, item := range feed.Channel.Item {
	    fmt.Println(item.Title)
	}
	fmt.Println()
	fmt.Println()

	return nil
} 
