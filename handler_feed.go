package main

import (
	"context"
	"fmt"
	"time"

	"github.com/JoStMc/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	ctx := context.Background()
	if len(cmd.args) < 2 {
	    return fmt.Errorf("insufficent number of arguments")
	} 

	currentTime := time.Now()
	params := database.CreateFeedParams{ 
		ID: uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name: cmd.args[0],
		Url: cmd.args[1],
		UserID: user.ID,
	}

	feed, err := s.db.CreateFeed(ctx, params)
	if err != nil {
	    return err
	} 

	followParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID: user.ID,
		FeedID: feed.ID,
	}
	_, err = s.db.CreateFeedFollow(ctx, followParams)
	if err != nil {
		return fmt.Errorf("unable to add to feed follows: %w", err)
	} 

	fmt.Printf("Feed added for %s: %s\n", s.cfg.CurrentUserName, feed.Name)
	return nil
} 

func handlerListFeeds(s *state, cmd command) error {
	rows, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return err
	} 

	if len(rows) == 0 {
		fmt.Println("No feeds found.")
		return nil
	} 

	fmt.Println()
	for _, row := range rows {
		fmt.Printf("- %s\n    URL: %s\n    Created by: %s\n\n", row.Feedname, row.Url, row.Username)
	}
	return nil
} 
