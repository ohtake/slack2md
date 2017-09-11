package main

import (
	"time"
)

type Resolver struct {
	channels map[string]*Channel
	users    map[string]*User
}

type MessageResolved struct {
	User          *User
	BotID         string
	Subtype       string
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
func (l *messageTokenListener) OnChannel(channelID, alt string) {
	channel := l.resolver.ResolveChannel(channelID)
	text := alt
	if "" == text {
		text = channel.Name
	}
	l.add(MessageTokenChannel{channel, text})
}
func (l *messageTokenListener) OnUser(userID, alt string) {
	user := l.resolver.ResolveUser(userID)
	text := alt
	if "" == text {
		text = user.Name
	}
	l.add(MessageTokenUser{user, text})
}
func (l *messageTokenListener) OnVariable(variableID, alt string) {
	text := alt
	if "" == text {
		text = variableID
	}
	l.add(MessageTokenVariable{variableID, text})
}

func newMessageTokenListener(r *Resolver) *messageTokenListener {
	return &messageTokenListener{r, nil}
}

func NewResolver(channels []Channel, users []User) *Resolver {
	r := new(Resolver)
	r.channels = make(map[string]*Channel)
	for i, c := range channels {
		r.channels[c.ID] = &channels[i]
	}
	r.users = make(map[string]*User)
	for i, u := range users {
		r.users[u.ID] = &users[i]
	}
	return r
}

func (r *Resolver) ResolveUser(userID string) *User {
	return r.users[userID]
}
func (r *Resolver) ResolveChannel(channelID string) *Channel {
	return r.channels[channelID]
}

func (r *Resolver) Resolve(m *Message) MessageResolved {
	res := new(MessageResolved)

	if "" != m.User {
		res.User = r.ResolveUser(m.User)
	}
	res.BotID = m.BotID

	res.Ts = SlackTsToTime(m.Ts)
	res.Subtype = m.Subtype

	messageListner := newMessageTokenListener(r)
	NewSlackMessageParser(messageListner).Parse(m.Text)
	res.MessageTokens = messageListner.Tokens

	return *res
}
