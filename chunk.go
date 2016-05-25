package main

import (
	"io/ioutil"
	"path/filepath"
	"sort"
	"time"
)

// ChunkedHistoryReader reads Slack history json into a chunk.
// Reading whole history may eventually malloc all computer storage.
// Keep in mind that chunks should be garbage-collectable after use.
type ChunkedHistoryReader struct {
	ChunkSize        int
	ChannelDirectory string
	messageChannel   chan Message
}

func NewChunkedHistoryReader(chunkSize int, channelDirectory string) *ChunkedHistoryReader {
	r := new(ChunkedHistoryReader)
	r.ChunkSize = chunkSize
	r.ChannelDirectory = channelDirectory
	r.messageChannel = make(chan Message, chunkSize)
	go r.startReading()
	return r
}

func (r *ChunkedHistoryReader) startReading() {
	fileInfos, _ := ioutil.ReadDir(r.ChannelDirectory)
	fileNames := make([]string, len(fileInfos))
	for _, fi := range fileInfos {
		fileNames = append(fileNames, fi.Name())
	}
	sort.Strings(fileNames)
	for _, f := range fileNames {
		messageSlice := ReadHistory(filepath.Join(r.ChannelDirectory, f))
		for _, m := range messageSlice {
			r.messageChannel <- m
		}
	}
	close(r.messageChannel)
}

func (r *ChunkedHistoryReader) NextChunk() []Message {
	messages := make([]Message, 0, r.ChunkSize)
	for i := 0; i < r.ChunkSize; i++ {
		m, ok := <-r.messageChannel
		if ok {
			messages = append(messages, m)
		} else {
			break
		}
	}
	return messages
}

// ChunkInfo contains information about a chunk.
// It does not contain messages and its memory usage is O(1).
type ChunkInfo struct {
	ChannelName string
	NumMessages int
	Start       time.Time
	End         time.Time
}

func ToChunkInfo(channelName string, chunk []MessageResolved) ChunkInfo {
	info := new(ChunkInfo)
	info.ChannelName = channelName
	info.NumMessages = len(chunk)
	info.Start = chunk[0].Ts
	info.End = chunk[len(chunk)-1].Ts
	return *info
}
