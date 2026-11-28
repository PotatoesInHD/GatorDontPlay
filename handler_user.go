package main

import (
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/PotatoesInHD/GatorDontPlay/internal/database"
	"github.com/google/uuid"

)


func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("command has no arguments")
	}

	userName := cmd.args[0]
	dbUser, err := s.dbPtr.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("db query get user failed: %w", err)
	}

	if err := s.configPtr.SetUser(dbUser.Name); err != nil {
		return fmt.Errorf("%w", err)
	}

	fmt.Printf("User: %s has been set\n", dbUser.Name)

	return nil
}


func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("command has no argument")
	}

	userName := cmd.args[0]
	timeNow := time.Now()
	dBUserParams := database.CreateUserParams {
		ID:	uuid.New(),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Name: userName,
	}

	dbUser, err := s.dbPtr.CreateUser(context.Background(), dBUserParams)
	if err != nil {
		return fmt.Errorf("db create user failed: %w", err)
	}

	if err := s.configPtr.SetUser(dbUser.Name); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}


func handlerReset(s *state, cmd command) error {
	err := s.dbPtr.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("db delete users failed: %w", err)
	}

	fmt.Println("Database: Users table was reset successfully!")

	return nil
}


func handlerUsers(s *state, cmd command) error {
	dbUsers, err := s.dbPtr.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("db get list of users failed: %w", err)
	}

	for _, user := range dbUsers {
		if user.Name == s.configPtr.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}


func handlerAgg(s *state, cmd command) error {
  	if len(cmd.args) == 0 {
		return errors.New("command has no argument")
		}
	_, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
		if err != nil {
			return fmt.Errorf("fetch feed failed: %w", err)
		}

	return nil
}


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
