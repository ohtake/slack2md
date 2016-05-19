package main

import (
	"testing"
)

func TestReadChannels(t *testing.T) {
	actual := ReadChannels("test_data/channels.json")
	expected := []Channel{
		{"C00000001", "channel1", "1463128988"},
		{"C00000002", "channel2", "1436515520"},
	}
	if len(actual) != len(expected) {
		t.Errorf("Wrong length: %d, %d", len(actual), len(expected))
		t.FailNow()
	}
	for i := 0; i < len(expected); i++ {
		if actual[i].Id != expected[i].Id {
			t.Errorf("Wrong Id: %q, %q", actual[i].Id, expected[i].Id)
		}
		if actual[i].Name != expected[i].Name {
			t.Errorf("Wrong Name: %q, %q", actual[i].Name, expected[i].Name)
		}
		if actual[i].Created != expected[i].Created {
			t.Errorf("Wrong Created: %q, %q", actual[i].Created, expected[i].Created)
		}
	}
}

func TestReadUsers(t *testing.T) {
	actual := ReadUsers("test_data/users.json")
	expected := []User{
		{"U00000001", "alice"},
		{"U00000002", "bob"},
	}
	if len(actual) != len(expected) {
		t.Errorf("Wrong length: %d, %d", len(actual), len(expected))
		t.FailNow()
	}
	for i := 0; i < len(expected); i++ {
		if actual[i].Id != expected[i].Id {
			t.Errorf("Wrong Id: %q, %q", actual[i].Id, expected[i].Id)
		}
		if actual[i].Name != expected[i].Name {
			t.Errorf("Wrong Name: %q, %q", actual[i].Name, expected[i].Name)
		}
	}
}

func TestReadHistory(t *testing.T) {
	actual := ReadHistory("test_data/channel1/2016-05-13.json")
	expected := []Message{
		{"U00000001", "<@U00000001|alice> has joined the channel", "1463128987.000002"},
		{"U00000001", "<@U00000001|alice> set the channel purpose: ", "1463128989.000003"},
		{"U00000002", "<@U00000002|bob> has joined the channel", "1463128989.000004"},
		{"U00000001", "Hello, <@U00000002>", "1463129038.000008"},
	}
	if len(actual) != len(expected) {
		t.Errorf("Wrong length: %d, %d", len(actual), len(expected))
		t.FailNow()
	}
	for i := 0; i < len(expected); i++ {
		if actual[i].User != expected[i].User {
			t.Errorf("Wrong User: %q, %q", actual[i].User, expected[i].User)
		}
		if actual[i].Text != expected[i].Text {
			t.Errorf("Wrong Text: %q, %q", actual[i].Text, expected[i].Text)
		}
		if actual[i].Ts != expected[i].Ts {
			t.Errorf("Wrong Ts: %q, %q", actual[i].Ts, expected[i].Ts)
		}
	}
}
