package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/JoStMc/gator/internal/database"
	"github.com/google/uuid"
)

func handlerListUserFeeds(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	feedFollows, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("unable to get feed follows: %w", err)
	} 

	if len(feedFollows) == 0 {
	    fmt.Println("You are not following any feeds.")
		return nil
	} 

	fmt.Printf("Feeds followed by %s\n", feedFollows[0].Username)
	for _, feed := range feedFollows {
		fmt.Printf("- %s\n    URL: %s\n\n", feed.Feedname, feed.Url)
	} 

	return nil
} 

func handlerFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
	    return fmt.Errorf("insufficient number of args")
	} 

	ctx := context.Background()

	feed, err := s.db.GetFeed(ctx, cmd.args[0])
	if err != nil {
		return fmt.Errorf("feed not found")
	} 

	currentTime := time.Now()
	params := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID: user.ID,
		FeedID: feed.ID,
	} 

	followedFeed, err := s.db.CreateFeedFollow(ctx, params)
	if err != nil {
		return fmt.Errorf("unable to create feed follow: %w", err)
	} 

	fmt.Printf("%s followed feed %s\n", s.cfg.CurrentUserName, followedFeed.Feedname)
	return nil
} 

func handlerUnfollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
	    return fmt.Errorf("insufficient number of arguments")
	} 
	ctx := context.Background()

	feed, err := s.db.GetFeed(ctx, cmd.args[0])
	if err != nil {
		return fmt.Errorf("unable to get feed: %w", err)
	} 

	err = s.db.UnfollowFeed(ctx, database.UnfollowFeedParams{
												FeedID: feed.ID, 
												UserID: user.ID,})
	if err != nil {
	    return err
	} 
	fmt.Printf("%s successfully unfollowed from feed %s\n", user.Name, feed.Name)
	return nil
} 


func handlerBrowse(s* state, cmd command) error {
	limit := 2
	if len(cmd.args) >= 1 {
		var err error
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
		    return err
		} 
	} 

	userPosts, err := s.db.GetPostsForUser(context.Background())
	if err != nil {
	    return err
	} 
	limit = min(len(userPosts), limit)

	fmt.Println("Posts fetched:")
	fmt.Println()
	for i := range limit {
	    fmt.Println(userPosts[i].Title)
	    fmt.Println(userPosts[i].Description)
		fmt.Println()
		fmt.Println()
	} 

	return nil
} 
