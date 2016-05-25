package main

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Slack message format is described at https://api.slack.com/docs/formatting

type SlackMessageListener interface {
	OnNewLine()
	OnText(text string)
	OnUser(userId string, alt string)
	OnChannel(channelId string, alt string)
	OnVariable(variableId string, alt string)
	OnLink(href string, text string)
}

type SlackMessageParser struct {
	listener SlackMessageListener
}

var reLink = regexp.MustCompile("<([^|>][^>]*)>")

func NewSlackMessageParser(listener SlackMessageListener) *SlackMessageParser {
	p := new(SlackMessageParser)
	p.listener = listener
	return p
}

func (p *SlackMessageParser) Parse(text string) {
	lines := strings.Split(text, "\n")
	for i, l := range lines {
		p.parseLine(l)
		if i+1 < len(lines) {
			p.listener.OnNewLine()
		}
	}
}

func (p *SlackMessageParser) parseLine(line string) {
	for remainingLine := line; len(remainingLine) > 0; {
		submatch := reLink.FindStringSubmatchIndex(remainingLine)
		if submatch == nil {
			p.listener.OnText(remainingLine)
			return
		}
		if submatch[0] > 0 {
			p.listener.OnText(remainingLine[:submatch[0]])
		}
		p.parseLink(remainingLine[submatch[2]:submatch[3]])
		remainingLine = remainingLine[submatch[1]:]
	}
}

func (p *SlackMessageParser) parseLink(linkText string) {
	splits := strings.SplitN(linkText, "|", 2)
	alt := ""
	if len(splits) == 2 {
		alt = splits[1]
	}
	switch splits[0][0] {
	case '@':
		p.listener.OnUser(splits[0][1:], alt)
	case '#':
		p.listener.OnChannel(splits[0][1:], alt)
	case '!':
		p.listener.OnVariable(splits[0][1:], alt)
	default:
		p.listener.OnLink(splits[0], alt)
	}
}

func SlackTsToTime(timestamp string) time.Time {
	f, _ := strconv.ParseFloat(timestamp, 64)
	sec := int64(f)
	nsec := int64((f - float64(sec)) * 1000000)
	return time.Unix(sec, nsec)
}
