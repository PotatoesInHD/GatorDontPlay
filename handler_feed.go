package main

import (
	"context"
	"fmt"
	"github.com/PotatoesInHD/GatorDontPlay/internal/database"
)


func handlerAddFeed(s *state, cmd command) error {
	argReqCount := 2
	if len(cmd.args) < argReqCount {
		return fmt.Errorf("addfeed command requires %d arguments", argReqCount)
	}

	//get current user information to pass into CreateFeed
	dbUser, err := s.dbPtr.GetUser(context.Background(), s.configPtr.CurrentUserName)
	if err != nil {
		return fmt.Errorf("db get list of users failed: %w", err)
	}

	//load up param struct
	dBFeedParams := database.CreateFeedParams {
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: cmd.args[0],
		Url: cmd.args[1],
		UserID: dbUser.ID,
	}

	_, err = s.dbPtr.CreateFeed(context.Background(), dBFeedParams)
	if err != nil {
		return fmt.Errorf("db create feed failed: %w", err)
	}

	return nil
}


func handlerFeeds(s *state, cmd command) error {
	dbFeeds, err := s.dbPtr.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("db feeds failed: %w", err)
	}

	for _, feed := range dbFeeds {
		fmt.Printf("%v\n%v\n%v\n", feed.Name, feed.Url, feed.Uname)
	}

	return nil
}
