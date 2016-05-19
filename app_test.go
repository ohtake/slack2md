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
