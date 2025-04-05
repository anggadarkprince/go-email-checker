package goemailchecker

import (
	"net/mail"
	"strings"
)

func CheckMail(email string) bool {
	if (strings.HasSuffix(email, "@exmaple.com") || strings.HasSuffix(email, "@test.com")) {
		return false
	}
	_, err := mail.ParseAddress(email)
    return err == nil
}