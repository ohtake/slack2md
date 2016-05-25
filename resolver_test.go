package main

import (
	"testing"
	"time"
)

func newResolver() *Resolver {
	return NewResolver(ReadChannels("test_data/channels.json"), ReadUsers("test_data/users.json"))
}

func TestResolve(t *testing.T) {
	resolver := newResolver()

	expected := []struct{ id, name string }{
		{"U00000001", "alice"},
		{"U00000002", "bob"},
	}

	for _, e := range expected {
		user := resolver.ResolveUser(e.id)
		if user == nil {
			t.Errorf("Cannot find user: %q", e.id)
		} else if user.Name != e.name {
			t.Errorf("Found different user: %q, %q", user.Name, e.name)
		}
	}
	user2 := resolver.ResolveUser("U99999999")
	if user2 != nil {
		t.Error("Found unexpected user")
	}
}

func TestResolveMessage(t *testing.T) {
	resolver := newResolver()
	messages := ReadHistory("test_data/channel1/2016-05-13.json")

	m1 := resolver.Resolve(&messages[0])
	if m1.User.Name != "alice" {
		t.Errorf("Cannot resolve user: %q, &q", m1.User.Name, "alice")
	}
	if m1.Ts.Before(time.Date(2016, 5, 13, 8, 43, 7, 0, time.UTC)) ||
		m1.Ts.After(time.Date(2016, 5, 13, 8, 43, 8, 0, time.UTC)) {
		t.Errorf("Cannot resolve ts: %v", m1.Ts)
	}
	if len(m1.MessageTokens) != 2 {
		t.Error("Cannot resolve first message text")
	} else {
		token1, _ := m1.MessageTokens[0].(MessageTokenUser)
		if token1.User == nil {
			t.Error("Cannot resolve first message user")
		}
	}

	m2 := resolver.Resolve(&messages[3])
	if m2.Ts.Before(time.Date(2016, 5, 13, 8, 43, 58, 0, time.UTC)) ||
		m2.Ts.After(time.Date(2016, 5, 13, 8, 43, 59, 0, time.UTC)) {
		t.Errorf("Cannot resolve ts: %v", m2.Ts)
	}
}
