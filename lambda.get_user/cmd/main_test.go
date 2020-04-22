package main

import (
	"os"
	"testing"

	"github.com/reecerussell/tw-management-system/core/users"
)

func TestGetUser(t *testing.T) {
	os.Setenv("AWS_ACCESS_KEY_ID", "AKIAIIBGLO3OFXXP7OMA")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "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")
	os.Setenv("AWS_REGION", "eu-west-2")

	repo := users.NewRepository()

	u, err := repo.Get("00000300-0000-0000-0000-000000000000")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("| id | username | email | password hash |\n")
	t.Logf("| %s | %s | %s | %s |\n", u.ID, u.Username, u.Email, u.PasswordHash)
}
