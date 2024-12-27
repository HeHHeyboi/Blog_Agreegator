package main

import (
	"context"
	"fmt"
	"html"
	"strconv"
)

func handlerBrowse(s *state, ctx context.Context, cmd command) error {
	var limit int32 = 2
	if len(cmd.arg) > 0 {
		parseInt, err := strconv.Atoi(cmd.arg[0])
		if err != nil {
			return fmt.Errorf("Please input number")
		}
		limit = int32(parseInt)
	}

	posts, err := s.db.GetPosts(ctx, limit)
	if err != nil {
		return fmt.Errorf("Error when get Posts: %v", err)
	}

	for i, post := range posts {
		fmt.Println("No.", i+1)
		title := html.UnescapeString(post.Title.String)
		fmt.Println("Title:", title)
		fmt.Println("Url:", post.Url)
		desc := html.UnescapeString(post.Description.String)
		fmt.Println("Descripton:", desc)
		fmt.Println("Publish:", post.PublishedAt)
		fmt.Println()
	}

	return nil
}
