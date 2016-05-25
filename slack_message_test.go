package main

import (
	"testing"
)

type recordableListner struct {
	records []record
}

type record struct {
	Type   recordType
	Text   string
	Target string
}

type recordType int

const (
	rtNewLine recordType = iota
	rtText
	rtUser
	rtChannel
	rtVariable
	rtLink
)

func (l *recordableListner) addRecord(rt recordType, text string, target string) {
	l.records = append(l.records, record{rt, text, target})
}

func (l *recordableListner) testEquals(expected []record, t *testing.T) {
	if len(l.records) != len(expected) {
		t.Errorf("Wrong length of records: %v, %v", len(l.records), len(expected))
		return
	}
	for i := 0; i < len(l.records); i++ {
		ra := l.records[i]
		re := expected[i]
		if ra.Type != re.Type || ra.Text != re.Text || ra.Target != re.Target {
			t.Errorf("Wrong record: %v, %v", ra, re)
		}
	}
}

func (l *recordableListner) OnNewLine() {
	l.addRecord(rtNewLine, "", "")
}
func (l *recordableListner) OnText(text string) {
	l.addRecord(rtText, text, "")
}
func (l *recordableListner) OnUser(userId, alt string) {
	l.addRecord(rtUser, alt, userId)
}
func (l *recordableListner) OnChannel(channelId, alt string) {
	l.addRecord(rtChannel, alt, channelId)
}
func (l *recordableListner) OnVariable(variableId, alt string) {
	l.addRecord(rtVariable, alt, variableId)
}
func (l *recordableListner) OnLink(href, text string) {
	l.addRecord(rtLink, text, href)
}

func TestSlackMessageParser(t *testing.T) {
	cases := []struct {
		input    string
		expected []record
	}{
		{
			"",
			nil,
		},
		{
			"single line",
			[]record{
				{rtText, "single line", ""},
			},
		},
		{
			"multi\nline",
			[]record{
				{rtText, "multi", ""},
				{rtNewLine, "", ""},
				{rtText, "line", ""},
			},
		},
		{
			"links\n<http://example.com/> <http://example.com/|example | example>",
			[]record{
				{rtText, "links", ""},
				{rtNewLine, "", ""},
				{rtLink, "", "http://example.com/"},
				{rtText, " ", ""},
				{rtLink, "example | example", "http://example.com/"},
			},
		},
		{
			"user, variable, channel\n<@U11111111>, <@U22222222|bob>, and <!here|here>\n\nPlease check <#C11111111> and <#C22222222|channel2> channels",
			[]record{
				{rtText, "user, variable, channel", ""},
				{rtNewLine, "", ""},
				{rtUser, "", "U11111111"},
				{rtText, ", ", ""},
				{rtUser, "bob", "U22222222"},
				{rtText, ", and ", ""},
				{rtVariable, "here", "here"},
				{rtNewLine, "", ""},
				{rtNewLine, "", ""},
				{rtText, "Please check ", ""},
				{rtChannel, "", "C11111111"},
				{rtText, " and ", ""},
				{rtChannel, "channel2", "C22222222"},
				{rtText, " channels", ""},
			},
		},
	}

	for _, c := range cases {
		actual := new(recordableListner)
		parser := NewSlackMessageParser(actual)
		parser.Parse(c.input)

		actual.testEquals(c.expected, t)
	}
}
