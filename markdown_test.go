package main

import (
	"testing"
)

func testLines(t *testing.T, actual []string, expected []string) bool {
	if len(actual) != len(expected) {
		t.Errorf("Wrong number of lines: %v, %v", len(actual), len(expected))
		return false
	}
	hasError := false
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("Wrong line(%v): %q, %q", i, actual[i], expected[i])
			hasError = true
		}
	}
	return hasError
}

func TestMarkdownEscape(t *testing.T) {
	m := NewMarkdownTranslator()
	cases := []struct {
		expected string
		input    string
	}{
		{"plain text", "plain text"},
		{"\\#\\[\\]\\<\\>\\\\\\!\\*\\`\\|", "#[]<>\\!*`|"},
	}
	for _, c := range cases {
		actual := m.Escape(c.input)
		if actual != c.expected {
			t.Errorf("Wrong escape: %q, %q", actual, c.expected)
		}
	}
}

func TestMarkdownHeadings(t *testing.T) {
	m := NewMarkdownTranslator()
	cases := []struct {
		expected string
		level    int
		text     string
	}{
		{"# heading 1", 1, "heading 1"},
		{"## heading 2", 2, "heading 2"},
		{"## \\#channel-name", 2, "#channel-name"},
	}
	for _, c := range cases {
		testLines(t, m.ToHeading(c.level, c.text), []string{c.expected, ""})
	}
}

func TestMarkdownChannelList(t *testing.T) {
	m := NewMarkdownTranslator()
	channels := ReadChannels("test_data/channels.json")
	testLines(t, m.ToChannelList(channels), []string{
		"* [\\#channel1](channel--channel1.md)",
		"* [\\#channel2](channel--channel2.md)",
		"",
	})
}

func TestMarkdownChunkList(t *testing.T) {
	// TODO
}

func TestMarkdownMessageList(t *testing.T) {
	// TODO
}

func TestMarkdownUserTable(t *testing.T) {
	// TODO
}
