package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Story stores the components of individual stories.
type Story struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// Option is the struct with stores the option to go to the next story
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func getStoriesFromJSON(jsonFilePath string) (map[string]Story, error) {
	var storyMap map[string]Story
	file, err := os.Open(jsonFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)

	err = json.Unmarshal(byteValue, &storyMap)
	if err != nil {
		return nil, err
	}
	return storyMap, nil
}
