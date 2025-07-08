package userdomain

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"regexp"
)

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()

	defer hash.Reset()
	hash.Write([]byte(ud.GetPassword()))
	ud.setPassword(hex.EncodeToString(hash.Sum(nil)))
}

func (ud *userDomain) ValidatePassword() error {
	if len(ud.password) < 4 {
		return errors.New("password must be at least 4 characters")
	}

	hasLetter := regexp.MustCompile(`[A-Za-z]`).MatchString
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString

	if !hasLetter(ud.password) || !hasNumber(ud.password) {
		return errors.New("password must contain both letters and numbers")
	}

	return nil
}
