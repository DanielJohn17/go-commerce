package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Error("expected hash to not be empty")
	}

	if hash == "password" {
		t.Error("expected hashed password")
	}
}

func TestComparePassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !ComparePassword(hash, []byte("password")) {
		t.Error("expected password to match hash")
	}
	if ComparePassword(hash, []byte("notpassword")) {
		t.Errorf("expected password to not match hash")
	}
}
