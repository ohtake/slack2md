package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type MarkdownTranslator struct {
}

func NewMarkdownTranslator() *MarkdownTranslator {
	r := new(MarkdownTranslator)
	return r
}

func (t *MarkdownTranslator) FileNameIndex() string {
	return "index.md"
}
func (t *MarkdownTranslator) FileNameChannel(channelName string) string {
	return fmt.Sprintf("channel--%v.md", channelName)
}
func (t *MarkdownTranslator) FileNameHistory(channelName string, chunkNumber int) string {
	return fmt.Sprintf("history--%v--%v.md", channelName, chunkNumber)
}

var mdSpecialCharsRegexp = regexp.MustCompile(`[\\\[\]#*!<>` + "`" + `|]`)
var mdSpecialCharsReplacer = func(matched string) string {
	return "\\" + matched
}

func (t *MarkdownTranslator) Escape(text string) string {
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
		result = append(result, fmt.Sprintf("* [%v](%v)", t.Escape(text), link))
	}
	result = append(result, "")
	return result
}

func (t *MarkdownTranslator) ToChunkList(chunks []ChunkInfo) []string {
	result := make([]string, 0, len(chunks)+1)
	for i, ch := range chunks {
		text := fmt.Sprintf("%v (%v - %v)", i+1, toTimeStampString(ch.Start), toTimeStampString(ch.End))
		link := t.FileNameHistory(ch.ChannelName, i+1)
		result = append(result, fmt.Sprintf("* [%v](%v)", t.Escape(text), link))
	}
	result = append(result, "")
	return result
}

func (t *MarkdownTranslator) ToMessageList(chunk []MessageResolved) []string {
	result := make([]string, 0, len(chunk)+1)
	for _, m := range chunk {
		md := make([]string, 0, 1+len(m.MessageTokens))
		header := toTimeStampString(m.Ts)
		if nil != m.User {
			header += " @" + m.User.Name
		} else if "" != m.BotID {
			header += " (APP)" + m.BotID
		} else {
			// `subtype=file_comment` does not have `user` or `bot_id`
		}
		header += ":"
		md = append(md, "* *"+header+"* ")
		for _, token := range m.MessageTokens {
			switch token := token.(type) {
			case MessageTokenEndSubtype:
				md = append(md, "*")
			case MessageTokenStartSubtype:
				md = append(md, "*")
			case MessageTokenNewLine:
				md = append(md, "\n    ")
			case MessageTokenText:
				md = append(md, token.Text)
			case MessageTokenLink:
				md = append(md, fmt.Sprintf("[%v](%v)", t.Escape(token.Text), token.Href))
			case MessageTokenChannel:
				md = append(md, "#"+token.Text)
			case MessageTokenUser:
				md = append(md, "@"+token.Text)
			case MessageTokenVariable:
				md = append(md, "@"+token.Text)
			default:
				panic("Unknown message type")
			}
		}
		result = append(result, strings.Join(md, ""))
	}
	result = append(result, "")
	return result
}

func (t *MarkdownTranslator) ToUserTable(users []User) []string {
	result := make([]string, 0, len(users)+2+1)
	result = append(result, "|ID|Icon|Name|Email|FirstName|LastName|Title|")
	result = append(result, strings.Repeat("|----", 7)+"|")
	for _, u := range users {
		vals := []string{u.ID, "![](" + u.Profile.Image24 + ")", u.Name, u.Profile.Email, t.Escape(u.Profile.FirstName), t.Escape(u.Profile.LastName), t.Escape(u.Profile.Title)}
		result = append(result, "|"+strings.Join(vals, "|")+"|")
	}
	result = append(result, "")
	return result
}

func toTimeStampString(ts time.Time) string {
	return ts.UTC().Format(time.RFC3339)
}
