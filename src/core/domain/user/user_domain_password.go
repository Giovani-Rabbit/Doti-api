package userdomain

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"regexp"
)

func (ud *userDomain) EncryptPassword() error {
	if err := validatePassword(ud.GetPassword()); err != nil {
		return err
	}

	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.GetPassword()))

	ud.setPassword(hex.EncodeToString(hash.Sum(nil)))

	return nil
}

func validatePassword(password string) error {
	if len(password) < 4 {
		return errors.New("password must be at least 8 characters")
	}

	hasLetter := regexp.MustCompile(`[A-Za-z]`).MatchString
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString

	if !hasLetter(password) || !hasNumber(password) {
		return errors.New("password must contain both letters and numbers")
	}

	return nil
}
