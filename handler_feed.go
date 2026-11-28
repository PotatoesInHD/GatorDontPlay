package main

import (
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/PotatoesInHD/GatorDontPlay/internal/database"
	"github.com/google/uuid"
)


func handlerAddFeed(s *state, cmd command, dbUser database.User) error {
	argReqCount := 2
	if len(cmd.args) < argReqCount {
		return fmt.Errorf("addfeed command requires %d arguments", argReqCount)
	}

	//load up param struct
	timeNow := time.Now()
	dBFeedParams := database.CreateFeedParams {
		ID: uuid.New(),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Name: cmd.args[0],
		Url: cmd.args[1],
		UserID: dbUser.ID,
	}
	//add feed
	dbFeed, err := s.dbPtr.CreateFeed(context.Background(), dBFeedParams)
	if err != nil {
		return fmt.Errorf("db create feed failed: %w", err)
	}

	//follow feed after its added
	//load up db params struct
	dbFeedFollowParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		UserID:    dbUser.ID,
		FeedID:    dbFeed.ID,
	}

	_, err = s.dbPtr.CreateFeedFollow(context.Background(), dbFeedFollowParams)
	if err != nil {
		return fmt.Errorf("create feed follow failed: %w", err)
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


func handlerFollow(s *state, cmd command, dbUser database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("command has no arguments")
	}

	//get feed by url
	dbFeed, err := s.dbPtr.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("db get feed failed: %w", err)
	}

	//load up db params struct
	timeNow := time.Now()
	dbFeedFollowParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		UserID:    dbUser.ID,
		FeedID:    dbFeed.ID,
	}

	dbFeedFollow, err := s.dbPtr.CreateFeedFollow(context.Background(), dbFeedFollowParams)
	if err != nil {
		return fmt.Errorf("create feed follow failed: %w", err)
	}
	fmt.Printf("%s followed by %s", dbFeedFollow.FeedName, dbFeedFollow.UserName)

	return nil
}


func handlerFollowing(s *state, cmd command, dbUser database.User) error {
	//get feed follows for current user
	dbFeedFollows, err := s.dbPtr.GetFeedFollowsForUser(context.Background(), dbUser.ID)
	if err != nil {
		return fmt.Errorf("db feed follows failed: %w", err)
	}

	for _, feedFollow := range dbFeedFollows {
		fmt.Printf("Feed: %v\nName: %v\n", feedFollow.Fname, feedFollow.Uname)
	}

	return nil
}


func handlerUnfollow(s *state, cmd command, dbUser database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("command has no arguments")
	}

	//get feed by url
	dbFeed, err := s.dbPtr.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("db get feed failed: %w", err)
	}

	//load up db params
	dbFeedUnfollowParams := database.FeedUnfollowParams{
		UserID: dbUser.ID,
		FeedID: dbFeed.ID,
	}

	if err := s.dbPtr.FeedUnfollow(context.Background(), dbFeedUnfollowParams); err != nil {
		return fmt.Errorf("db feed unfollow failed: %w", err)
	}

	return nil
}
