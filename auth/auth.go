package auth

import (
	"errors"
)

type Credential struct {
	Email    string
	Password string
}

var (
	DemoCred  = Credential{"nguyen@open-vn.org", "12345678"}
	DemoToken = "zaq123456789"
)

func Login(c Credential) (string, error) {
	if c.Email == DemoCred.Email && c.Password == DemoCred.Password {
		return DemoToken, nil
	}

	return "", errors.New("Invalid credential")
}

func ValidToken(token string) error {
	if token == DemoToken {
		return nil
	}

	return errors.New("Invalid token")
}
