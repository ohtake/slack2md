package main

import (
	"regexp"
	"strconv"
	"strings"
)

type MarkdownTranslator struct {
}

func NewMarkdownTranslator() *MarkdownTranslator {
	r := new(MarkdownTranslator)
	return r
}

func (f *MarkdownTranslator) FileNameIndex() string {
	return "index.md"
}
func (f *MarkdownTranslator) FileNameChannel(channelName string) string {
	return "channel--" + channelName + ".md"
}
func (f *MarkdownTranslator) FileNameHistory(channelName string, chunkNumber int) string {
	return "history--" + channelName + "--" + strconv.FormatInt(int64(chunkNumber), 10) + ".md"
}

var mdSpecialCharsRegexp = regexp.MustCompile("[\\\\\\[\\]#*!<>`]")
var mdSpecialCharsReplacer = func(matched string) string {
	return "\\" + matched
}

func (_ *MarkdownTranslator) Escape(text string) string {
	result := text
	result = mdSpecialCharsRegexp.ReplaceAllStringFunc(result, mdSpecialCharsReplacer)
	// TODO needs more for escaping
	return result
}

func (t *MarkdownTranslator) ToHeading(level int, text string) []string {
	return []string{
		strings.Repeat("#", level) + " " + t.Escape(text),
		"",
	}
}

func (t *MarkdownTranslator) ToChannelList(channels []Channel) []string {
	result := make([]string, 0, len(channels)+1)
	for _, ch := range channels {
		text := "#" + ch.Name
		link := t.FileNameChannel(ch.Name)
		result = append(result, "* ["+t.Escape(text)+"]("+link+")")
	}
	result = append(result, "")
	return result
}

func (t *MarkdownTranslator) ToChunkList(chunks []ChunkInfo) []string {
	result := make([]string, 0, len(chunks)+1)
	for i, ch := range chunks {
		text := strconv.FormatInt(int64(i+1), 10) + " (" + ch.Start.String() + " - " + ch.End.String() + ")"
		link := t.FileNameHistory(ch.ChannelName, i+1)
		result = append(result, "* ["+t.Escape(text)+"]("+link+")")
	}
	result = append(result, "")
	return result
}

func (t *MarkdownTranslator) ToMessageList(chunk []MessageResolved) []string {
	result := make([]string, 0, len(chunk)+1)
	for _, m := range chunk {
		// TODO
		result = append(result, "* "+m.Text)
	}
	result = append(result, "")
	return result
}
