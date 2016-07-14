package main

import (
	"testing"
)

func readAllChunks(chunkSize int, channelDirectory string) [][]Message {
	reader := NewChunkedHistoryReader(chunkSize, channelDirectory)
	var chunks [][]Message
	for c := reader.NextChunk(); len(c) > 0; c = reader.NextChunk() {
		chunks = append(chunks, c)
	}
	return chunks
}

func TestChunkedHistoryReader(t *testing.T) {
	actual := readAllChunks(4, "test_data/channel1")
	expectedChunkSizes := []int{4, 2}
	if len(expectedChunkSizes) != len(actual) {
		t.Errorf("Wrong length: %d, %d", len(expectedChunkSizes), len(actual))
		t.FailNow()
	}
	for i := 0; i < len(expectedChunkSizes); i++ {
		if expectedChunkSizes[i] != len(actual[i]) {
			t.Errorf("Wrong length: %d, %d", expectedChunkSizes[i], len(actual[i]))
			t.FailNow()
		}
	}
	first := actual[0][0]
	if first.Text != "<@U00000001|alice> has joined the channel" {
		t.Errorf("Wrong first message: %v", first.Text)
	}
	lastChunk := actual[len(actual)-1]
	last := lastChunk[len(lastChunk)-1]
	if last.Text != "Hello <@U00000002|bob>" {
		t.Errorf("Wrong last message: %v", last.Text)
	}
}

func TestToChunkInfo(t *testing.T) {
	chunks := readAllChunks(3, "test_data/channel1")
	expectedMessageSizes := []int{3, 3}
	expectedYear := 2016
	actual := make([]ChunkInfo, 0, len(chunks))
	for _, c := range chunks {
		actual = append(actual, ToChunkInfo("channel1", c))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i].NumMessages != expectedMessageSizes[i] {
			t.Errorf("Wrong number of messages: %v, %v", actual[i].NumMessages, expectedMessageSizes[i])
		}
		if actual[i].Start.Year() != expectedYear {
			t.Errorf("Wrong start: %v", actual[i].Start)
		}
		if actual[i].End.Year() != expectedYear {
			t.Errorf("Wrong end: %v", actual[i].Start)
		}
	}
}
