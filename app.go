package main

import (
	"fmt"
	"os"
	"path"
	"flag"
)

var chunkSize = flag.Int("messages", 500, "Number of messages per page")
var inputDir = flag.String("input", "slack_export", "Input directory which contains exported json files")
var outputDir = flag.String("output", "output", "Output directory")
var help = flag.Bool("help", false, "Print help")

var translator = NewMarkdownTranslator()

func createIndex(channels []Channel, users []User) {
	f, _ := os.Create(path.Join(*outputDir, translator.FileNameIndex()))
	defer f.Close()
	w := NewTranslatingWriter(translator, f)
	defer w.Flush()

	w.WriteHeading(1, "Exported Slack")
	w.WriteHeading(2, "Channels")
	w.WriteChannelList(channels)
	w.WriteHeading(2, "Users")
	w.WriteUserTable(users)
}

func createChannel(channel Channel, chunks []ChunkInfo) {
	f, _ := os.Create(path.Join(*outputDir, translator.FileNameChannel(channel.Name)))
	defer f.Close()
	w := NewTranslatingWriter(translator, f)
	defer w.Flush()

	w.WriteHeading(1, "Channel #"+channel.Name)
	w.WriteChunkList(chunks)
}

func createHistory(channel Channel, pageNumber int, messages []MessageResolved) {
	f, _ := os.Create(path.Join(*outputDir, translator.FileNameHistory(channel.Name, pageNumber)))
	defer f.Close()
	w := NewTranslatingWriter(translator, f)
	defer w.Flush()

	w.WriteHeading(1, fmt.Sprintf("Channel #%v (%v)", channel.Name, pageNumber))
	w.WriteMessageList(messages)
}

func resolveMessages(messages []Message, resolver *Resolver) []MessageResolved {
	r := make([]MessageResolved, 0, len(messages))
	for _, m := range messages {
		r = append(r, resolver.Resolve(&m))
	}
	return r
}

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		return
	}
	_ = os.Mkdir(*outputDir, 0755)

	channels := ReadChannels(path.Join(*inputDir, "channels.json"))
	users := ReadUsers(path.Join(*inputDir, "users.json"))
	resolver := NewResolver(channels, users)

	createIndex(channels, users)

	for _, ch := range channels {
		pageNumber := 0
		var chunks []ChunkInfo
		reader := NewChunkedHistoryReader(*chunkSize, path.Join(*inputDir, ch.Name))
		for chunk := reader.NextChunk(); len(chunk) > 0; chunk = reader.NextChunk() {
			pageNumber++
			messagesResolved := resolveMessages(chunk, resolver)
			createHistory(ch, pageNumber, messagesResolved)
			chunks = append(chunks, ToChunkInfo(ch.Name, chunk))
		}
		createChannel(ch, chunks)
	}
}
