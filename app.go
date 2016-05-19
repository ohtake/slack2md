package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	fmt.Printf("hello, world\n")

	channels := ReadChannels("slack_export/channels.json")
	fmt.Println(channels[0].Name)
	fmt.Println(channels[1].Name)

	users := ReadUsers("slack_export/users.json")
	fmt.Println(users[0].Name)
	fmt.Println(users[1].Name)
}
