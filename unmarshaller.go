package main

import (
	"encoding/json"
	"io/ioutil"
)

type Channel struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created"`
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	User string `json:"user"`
	Text string `json:"text"`
	Ts   string `json:"ts"`
}

func ReadChannels(channel_json_filename string) []Channel {
	var channels []Channel
	body, _ := ioutil.ReadFile(channel_json_filename)
	_ = json.Unmarshal(body, &channels)
	return channels
}

func ReadUsers(user_json_filename string) []User {
	var users []User
	body, _ := ioutil.ReadFile(user_json_filename)
	_ = json.Unmarshal(body, &users)
	return users
}

func ReadHistory(history_json_filename string) []Message {
	var messages []Message
	body, _ := ioutil.ReadFile(history_json_filename)
	_ = json.Unmarshal(body, &messages)
	return messages
}
