package test

import (
	"testing"
	"user-service/pkg/utils"
)

func TestHashPassword(t *testing.T) {
	password := "testpassword123"

	hash, err := utils.HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword failed: %v", err)
	}

	if hash == "" {
		t.Error("HashPassword returned empty string")
	}

	if hash == password {
		t.Error("HashPassword returned plain text password")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "testpassword123"
	wrongPassword := "wrongpassword"

	hash, err := utils.HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword failed: %v", err)
	}

	if !utils.CheckPasswordHash(password, hash) {
		t.Error("CheckPasswordHash should return true for correct password")
	}

	if utils.CheckPasswordHash(wrongPassword, hash) {
		t.Error("CheckPasswordHash should return false for wrong password")
	}

	if utils.CheckPasswordHash("", hash) {
		t.Error("CheckPasswordHash should return false for empty password")
	}

	if utils.CheckPasswordHash(password, "") {
		t.Error("CheckPasswordHash should return false for empty hash")
	}
}
