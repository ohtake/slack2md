package main

import (
	"testing"
)

func TestChunkedHistoryReader(t *testing.T) {
	reader := NewChunkedHistoryReader(3, "test_data/channel1")
	expected_chunk_sizes := []int{3, 2}
	actual := make([][]Message, 0)
	for c := reader.NextChunk(); len(c) > 0; c = reader.NextChunk() {
		actual = append(actual, c)
	}
	if len(expected_chunk_sizes) != len(actual) {
		t.Errorf("Wrong length: %d, %d", len(expected_chunk_sizes), len(actual))
		t.FailNow()
	}
	for i := 0; i < len(expected_chunk_sizes); i++ {
		if expected_chunk_sizes[i] != len(actual[i]) {
			t.Errorf("Wrong length: %d, %d", expected_chunk_sizes[i], len(actual[i]))
			t.FailNow()
		}
	}
	first := actual[0][0]
	if first.Text != "<@U00000001|alice> has joined the channel" {
		t.Errorf("Wrong first message: %v", first.Text)
	}
	lastChunk := actual[len(actual)-1]
	last := lastChunk[len(lastChunk)-1]
	if last.Text != "Hello" {
		t.Errorf("Wrong last message: %v", last.Text)
	}
}
