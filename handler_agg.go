package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	ctx := context.Background()
	feeds, err := s.db.ListFeeds(ctx)
	if err != nil {
	    return err
	} 

	for _, feed := range feeds {
		rssfeed, err := fetchFeed(ctx, feed.Url)
		if err != nil {
		    return err
		} 
		fmt.Println(rssfeed)
	} 

	return nil
} 
