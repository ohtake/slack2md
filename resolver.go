package main

import (
	"strconv"
	"time"
)

type Resolver struct {
	channels map[string]*Channel
	users    map[string]*User
}

type MessageResolved struct {
	User          *User
	BotId         string
	MessageTokens []MessageToken
	Ts            time.Time
}

type MessageToken interface{}
type MessageTokenNewLine struct {
}
type MessageTokenText struct {
	Text string
}
type MessageTokenLink struct {
	Href string
	Text string
}
type MessageTokenChannel struct {
	Channel *Channel
	Text    string
}
type MessageTokenUser struct {
	User *User
	Text string
}
type MessageTokenVariable struct {
	Variable string
	Text     string
}

type messageTokenListener struct {
	resolver *Resolver
	Tokens   []MessageToken
}

func (l *messageTokenListener) add(token MessageToken) {
	l.Tokens = append(l.Tokens, token)
}
func (l *messageTokenListener) OnNewLine() {
	l.add(MessageTokenNewLine{})
}
func (l *messageTokenListener) OnText(text string) {
	l.add(MessageTokenText{text})
}
func (l *messageTokenListener) OnLink(href, text string) {
	text2 := text
	if "" == text2 {
		text2 = href
	}
	l.add(MessageTokenLink{href, text2})
}
func (l *messageTokenListener) OnChannel(channelId, alt string) {
	channel := l.resolver.ResolveChannel(channelId)
	text := alt
	if "" == text {
		text = channel.Name
	}
	l.add(MessageTokenChannel{channel, text})
}
func (l *messageTokenListener) OnUser(userId, alt string) {
	user := l.resolver.ResolveUser(userId)
	text := alt
	if "" == text {
		text = user.Name
	}
	l.add(MessageTokenUser{user, text})
}
func (l *messageTokenListener) OnVariable(variableId, alt string) {
	text := alt
	if "" == text {
		text = variableId
	}
	l.add(MessageTokenVariable{variableId, text})
}

func newMessageTokenListener(r *Resolver) *messageTokenListener {
	return &messageTokenListener{r, nil}
}

type ChunkInfo struct {
	ChannelName string
	NumMessages int
	Start       time.Time
	End         time.Time
}

func NewResolver(channels []Channel, users []User) *Resolver {
	r := new(Resolver)
	r.channels = make(map[string]*Channel)
	for i, c := range channels {
		r.channels[c.Id] = &channels[i]
	}
	r.users = make(map[string]*User)
	for i, u := range users {
		r.users[u.Id] = &users[i]
	}
	return r
}

func (r *Resolver) ResolveUser(userId string) *User {
	return r.users[userId]
}
func (r *Resolver) ResolveChannel(channelId string) *Channel {
	return r.channels[channelId]
}

func (r *Resolver) Resolve(m *Message) MessageResolved {
	res := new(MessageResolved)

	if "" != m.User {
		res.User = r.ResolveUser(m.User)
	}
	res.BotId = m.BotId

	ts, _ := strconv.ParseFloat(m.Ts, 64)
	sec := int64(ts)
	nsec := int64((ts - float64(sec)) * 1000000)
	res.Ts = time.Unix(sec, nsec)

	messageListner := newMessageTokenListener(r)
	NewSlackMessageParser(messageListner).Parse(m.Text)
	res.MessageTokens = messageListner.Tokens

	return *res
}

func ToChunkInfo(channelName string, chunk []MessageResolved) ChunkInfo {
	info := new(ChunkInfo)
	info.ChannelName = channelName
	info.NumMessages = len(chunk)
	info.Start = chunk[0].Ts
	info.End = chunk[len(chunk)-1].Ts
	return *info
}
