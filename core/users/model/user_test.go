package model

import "testing"

func TestUpdateEmail(t *testing.T) {
	validEmail := "me@reece-russell.co.uk"
	u := &User{}

	err := u.UpdateEmail(validEmail)
	if err != nil {
		t.Errorf("expected nil but got: %v", err)
		return
	}

	invalidEmails := []string{
		"my email address",
		"hello@ gmail.com",
		"helloWorld@go",
		"",
	}

	for _, e := range invalidEmails {
		err = u.UpdateEmail(e)
		if err == nil {
			t.Errorf("expected error but got nil")
			return
		}
	}
}
