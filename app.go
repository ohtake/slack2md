package main

import (
	"os"
	"path"
	"strconv"
)

const chunkSize = 500
const inputDir = "slack_export"
const outputDir = "output"

var translator = NewMarkdownTranslator()

func createIndex(channels []Channel) {
	f, _ := os.Create(path.Join(outputDir, translator.FileNameIndex()))
	defer f.Close()
	w := NewTranslatingWriter(translator, f)
	defer w.Flush()

	w.WriteHeading(1, "Exported Slack")
	w.WriteChannelList(channels)
}

func createChannel(channel Channel, chunks []ChunkInfo) {
	f, _ := os.Create(path.Join(outputDir, translator.FileNameChannel(channel.Name)))
	defer f.Close()
	w := NewTranslatingWriter(translator, f)
	defer w.Flush()

	w.WriteHeading(1, "Channel #"+channel.Name)
	w.WriteChunkList(chunks)
}

func createHistory(channel Channel, pageNumber int, messages []MessageResolved) {
	f, _ := os.Create(path.Join(outputDir, translator.FileNameHistory(channel.Name, pageNumber)))
	defer f.Close()
	w := NewTranslatingWriter(translator, f)
	defer w.Flush()

	w.WriteHeading(1, "Channel #"+channel.Name+" ("+strconv.FormatInt(int64(pageNumber), 10)+")")
	w.WriteMessageList(messages)
}

func resolveMessages(messages []Message, resolver *UserResolver) []MessageResolved {
	r := make([]MessageResolved, 0, len(messages))
	for _, m := range messages {
		r = append(r, m.Resolve(resolver))
	}
	return r
}

func main() {
	_ = os.Mkdir(outputDir, os.ModeDir)

	channels := ReadChannels(path.Join(inputDir, "channels.json"))
	userResolver := NewUserResolver(ReadUsers(path.Join(inputDir, "users.json")))

	createIndex(channels)

	for _, ch := range channels {
		pageNumber := 0
		chunks := make([]ChunkInfo, 0)
		reader := NewChunkedHistoryReader(chunkSize, path.Join(inputDir, ch.Name))
		for chunk := reader.NextChunk(); len(chunk) > 0; chunk = reader.NextChunk() {
			pageNumber++
			messagesResolved := resolveMessages(chunk, userResolver)
			createHistory(ch, pageNumber, messagesResolved)
			chunks = append(chunks, ToChunkInfo(ch.Name, messagesResolved))
		}
		createChannel(ch, chunks)
	}
}
