package main

import (
	"fmt"
)

func main() {
	fmt.Printf("hello, world\n")

	channels := ReadChannels("slack_export/channels.json")
	fmt.Println(channels[0].Name)
	fmt.Println(channels[1].Name)

	users := ReadUsers("slack_export/users.json")
	fmt.Println(users[0].Name)
	fmt.Println(users[1].Name)

	messages := ReadHistory("test_data/channel1/2016-05-13.json")
	fmt.Println(messages[0].Text)
	fmt.Println(messages[1].Text)

	r := NewChunkedHistoryReader(50, "slack_export/random")
	for c := r.NextChunk(); len(c) > 0; c = r.NextChunk() {
		fmt.Println(len(c), cap(c))
	}
}
