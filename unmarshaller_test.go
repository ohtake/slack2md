package main

import (
	"reflect"
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
		if actual[i].ID != expected[i].ID {
			t.Errorf("Wrong Id: %q, %q", actual[i].ID, expected[i].ID)
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
		{"U00000001", "alice", UserProfile{"alice.doe@example.com", "Alice", "Doe", "Alice Doe", "Alice Doe", "title1", "https://avatars.slack-edge.com/2016-04-27/00000000000_01234567890abcdef012_24.jpg"}},
		{"U00000002", "bob", UserProfile{"bob.doe@example.com", "Bob", "Doe", "Bob Doe", "Bob Doe", "title2", "https://secure.gravatar.com/avatar/0123456789abcdef0123456789abcdef.jpg?s=24&d=https%3A%2F%2Fa.slack-edge.com%2F66f9%2Fimg%2Favatars%2Fava_0002-24.png"}},
	}
	if len(actual) != len(expected) {
		t.Errorf("Wrong length: %d, %d", len(actual), len(expected))
		t.FailNow()
	}
	for i := 0; i < len(expected); i++ {
		if actual[i].ID != expected[i].ID {
			t.Errorf("Wrong Id: %q, %q", actual[i].ID, expected[i].ID)
		}
		if actual[i].Name != expected[i].Name {
			t.Errorf("Wrong Name: %q, %q", actual[i].Name, expected[i].Name)
		}
		if !reflect.DeepEqual(actual[i].Profile, expected[i].Profile) {
			t.Errorf("Wrong Profile: %v, %v", actual[i].Profile, expected[i].Profile)
		}
	}
}

func TestReadHistory(t *testing.T) {
	actual := ReadHistory("test_data/channel1/2016-05-18.json")
	expected := []Message{
		{"U00000002", "", "Hello", "1463564356.000010"},
		{"", "B00000001", "Hello <@U00000002|bob>", "1463564356.595611"},
	}
	if len(actual) != len(expected) {
		t.Errorf("Wrong length: %d, %d", len(actual), len(expected))
		t.FailNow()
	}
	for i := 0; i < len(expected); i++ {
		if actual[i].User != expected[i].User {
			t.Errorf("Wrong User: %q, %q", actual[i].User, expected[i].User)
		}
		if actual[i].BotID != expected[i].BotID {
			t.Errorf("Wrong BotID: %q, %q", actual[i].BotID, expected[i].BotID)
		}
		if actual[i].Text != expected[i].Text {
			t.Errorf("Wrong Text: %q, %q", actual[i].Text, expected[i].Text)
		}
		if actual[i].Ts != expected[i].Ts {
			t.Errorf("Wrong Ts: %q, %q", actual[i].Ts, expected[i].Ts)
		}
	}
}
