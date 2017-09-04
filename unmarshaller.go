package main

import (
	"encoding/json"
	"io/ioutil"
)

type Channel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created"`
}

type User struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Profile UserProfile `json:"profile"`
}
type UserProfile struct {
	Email              string `json:"email"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	ReamName           string `json:"real_name"`
	ReamNameNormalized string `json:"real_name_normalized"`
	Title              string `json:"title"`
	Image24            string `json:"image_24"`
}

type Message struct {
	User    string `json:"user"`
	BotID   string `json:"bot_id"`
	Text    string `json:"text"`
	Subtype string `json:"subtype"`
	Ts      string `json:"ts"`
}

func ReadChannels(channelJSONFilename string) []Channel {
	var channels []Channel
	body, _ := ioutil.ReadFile(channelJSONFilename)
	_ = json.Unmarshal(body, &channels)
	return channels
}

func ReadUsers(userJSONFilename string) []User {
	var users []User
	body, _ := ioutil.ReadFile(userJSONFilename)
	_ = json.Unmarshal(body, &users)
	return users
}

func ReadHistory(historyJSONFilename string) []Message {
	var messages []Message
	body, _ := ioutil.ReadFile(historyJSONFilename)
	_ = json.Unmarshal(body, &messages)
	return messages
}
