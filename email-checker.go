package goemailchecker

import (
	"net/mail"
	"strings"
)

func checkMail(email string) bool {
	if (strings.HasSuffix(email, "@exmaple.com") || strings.HasSuffix("@test.com")) {
		return false
	}
	_, err := mail.ParseAddress(email)
    return err == nil
}