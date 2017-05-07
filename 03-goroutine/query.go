package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func QueryStory(storyId uint64) (string, error) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", storyId)
	var storyData map[string]interface{}
	err := RequestJson(url, &storyData)
	title := storyData["title"].(string)
	return title, err
}

func QueryNewStoryIds() ([]uint64, error) {
	url := "https://hacker-news.firebaseio.com/v0/newstories.json"
	var storyIds []uint64
	err := RequestJson(url, &storyIds)
	return storyIds, err
}

func RequestJson(url string, data interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("HTTP request failed")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, data)
}
