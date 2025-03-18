package models

type Message struct {
	Country string `json:"Country"`
	StoryID uint   `json:"storyId"`
	ID      uint   `json:"id"`
	Content string `json:content`
}
