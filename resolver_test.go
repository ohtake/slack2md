package main

import (
	"testing"
	"time"
)

func newUserResolver() *UserResolver {
	return NewUserResolver(ReadUsers("test_data/users.json"))
}

func TestResolve(t *testing.T) {
	userResolver := newUserResolver()

	expected := []struct{ id, name string }{
		{"U00000001", "alice"},
		{"U00000002", "bob"},
	}

	for _, e := range expected {
		user := userResolver.Resolve(e.id)
		if user == nil {
			t.Errorf("Cannot find user: %q", e.id)
		} else if user.Name != e.name {
			t.Errorf("Found different user: %q, %q", user.Name, e.name)
		}
	}
	user2 := userResolver.Resolve("U99999999")
	if user2 != nil {
		t.Error("Found unexpected user")
	}
}

func TestResolveMessage(t *testing.T) {
	userResolver := newUserResolver()
	messages := ReadHistory("test_data/channel1/2016-05-13.json")

	m1 := messages[0].Resolve(userResolver)
	if m1.User.Name != "alice" {
		t.Errorf("Cannot resolve user: %q, &q", m1.User.Name, "alice")
	}
	if m1.Ts.Before(time.Date(2016, 5, 13, 8, 43, 7, 0, time.UTC)) ||
		m1.Ts.After(time.Date(2016, 5, 13, 8, 43, 8, 0, time.UTC)) {
		t.Errorf("Cannot resolve ts: %v", m1.Ts)
	}
	if m1.Text != "@alice has joined the channel" {
		t.Errorf("Cannot resolve text: %q", m1.Text)
	}

	m2 := messages[3].Resolve(userResolver)
	if m2.Text != "Hello, @bob" {
		t.Errorf("Cannot resolve text: %q", m2.Text)
	}
}
