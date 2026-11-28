package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

	type RSSFeed struct {
		Channel struct {
			Title       string    `xml:"title"`
			Link        string    `xml:"link"`
			Description string    `xml:"description"`
			Item        []RSSItem `xml:"item"`
		} `xml:"channel"`
	}

	type RSSItem struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		PubDate     string `xml:"pubDate"`
	}


func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	//setup the request and header
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
		if err != nil {
			return nil, err
		}
	req.Header.Set("User-Agent", "gator")

	//create client and Do request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	//changes io.reader into []bytes
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll resonse error: %w", err)
	}

	//unmarshal into RSSFeed struct
	rssFeed := RSSFeed{}
	if err = xml.Unmarshal(body, &rssFeed); err != nil {
	    return nil, err
	}

	//from docs: UnescapeString unescapes entities like "&lt;" to become "<".
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i, item := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(item.Title)
		rssFeed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}

	return &rssFeed, nil
}
