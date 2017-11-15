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
	channels := ReadChannels("testdata/channels.json")
	testLines(t, m.ToChannelList(channels), []string{
		"* [\\#channel1](channel--channel1.md)",
		"* [\\#channel2](channel--channel2.md)",
		"",
	})
}

func TestMarkdownChunkList(t *testing.T) {
	m := NewMarkdownTranslator()
	chunkInfos := ReadAllChunksAsInfo(3, "testdata/channel1")
	actual := m.ToChunkList(chunkInfos)
	testLines(t, actual, []string{
		"* [1 (2016-05-13T08:43:07Z - 2016-05-13T08:43:09Z)](history--channel1--1.md)",
		"* [2 (2016-05-13T08:43:58Z - 2016-05-18T09:39:16Z)](history--channel1--2.md)",
		"",
	})
}

func TestMarkdownMessageList(t *testing.T) {
	channels := ReadChannels("testdata/channels.json")
	users := ReadUsers("testdata/users.json")
	resolver := NewResolver(channels, users)
	chunks := ReadAllChunks(3, "testdata/channel1")
	resolvedMessages := make([]MessageResolved, 0, len(chunks[0]))
	for _, m := range chunks[0] {
		resolvedMessages = append(resolvedMessages, resolver.Resolve(&m))
	}
	m := NewMarkdownTranslator()
	actual := m.ToMessageList(resolvedMessages)
	testLines(t, actual, []string{
		"* 2016-05-13T08:43:07Z @alice: *@alice has joined the channel*",
		"* 2016-05-13T08:43:09Z @alice: *@alice set the channel purpose: *",
		"* 2016-05-13T08:43:09Z @bob: *@bob has joined the channel*",
		"",
	})
}

func TestMarkdownUserTable(t *testing.T) {
	users := ReadUsers("testdata/users.json")
	m := NewMarkdownTranslator()
	actual := m.ToUserTable(users)
	testLines(t, actual, []string{
		"|ID|Icon|Name|Email|FirstName|LastName|Title|",
		"|----|----|----|----|----|----|----|",
		"|U00000001|![](https://avatars.slack-edge.com/2016-04-27/00000000000_01234567890abcdef012_24.jpg)|alice|alice.doe@example.com|Alice|Doe|title1|",
		"|U00000002|![](https://secure.gravatar.com/avatar/0123456789abcdef0123456789abcdef.jpg?s=24&d=https%3A%2F%2Fa.slack-edge.com%2F66f9%2Fimg%2Favatars%2Fava_0002-24.png)|bob|bob.doe@example.com|Bob|Doe|title2|",
		"",
	})
}
