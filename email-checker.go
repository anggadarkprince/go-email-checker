package goemailchecker

import (
	"net/mail"
	"strings"
)

func CheckMail(email string, isStrict bool) bool {
	if (isStrict) {
		if (strings.HasSuffix(email, "@exmaple.com") || strings.HasSuffix(email, "@test.com")) {
			return false
		}
	}
	_, err := mail.ParseAddress(email)
    return err == nil
}