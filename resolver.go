package main

import (
	"regexp"
	"strconv"
	"time"
)

var userRegexInMessage = regexp.MustCompile("(?:<@)U[0-9A-Z]{8}(?:\\|[-_A-Za-z0-9]+)?(?:>)")

type UserResolver struct {
	idMap map[string]*User
}

type MessageResolved struct {
	User *User
	Text string
	Ts   time.Time
}

func NewUserResolver(users []User) *UserResolver {
	r := new(UserResolver)
	r.idMap = make(map[string]*User)
	for i, u := range users {
		r.idMap[u.Id] = &users[i]
	}
	return r
}

func (r *UserResolver) Resolve(userId string) *User {
	return r.idMap[userId]
}

func (m *Message) Resolve(r *UserResolver) *MessageResolved {
	res := new(MessageResolved)

	res.User = r.Resolve(m.User)

	ts, _ := strconv.ParseFloat(m.Ts, 64)
	sec := int64(ts)
	nsec := int64((ts - float64(sec)) * 1000000)
	res.Ts = time.Unix(sec, nsec)

	replacer := func(matched string) string {
		return "@" + r.Resolve(matched[2:11]).Name
	}
	res.Text = userRegexInMessage.ReplaceAllStringFunc(m.Text, replacer)

	return res
}
