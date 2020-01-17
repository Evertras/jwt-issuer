package token

import (
	"strings"
	"testing"
)

const testUser = "testUser"

func TestPrivateKeyInitializes(t *testing.T) {
	if privateKey == nil {
		t.Error("Private key was nil")
	}
}

func TestTokenCreates(t *testing.T) {
	_, err := New(testUser)

	if err != nil {
		t.Fatal(err)
	}
}

func TestTokenCreatesAndParses(t *testing.T) {
	token, err := New(testUser)

	if err != nil {
		t.Fatal(err)
	}

	parsed, err := Parse(strings.NewReader(token))

	if err != nil {
		t.Fatal(err)
	}

	if userID, exists := parsed.Get("UserID"); !exists || userID != testUser {
		t.Fatalf("Expected user ID %q but got %q", testUser, userID)
	}
}
