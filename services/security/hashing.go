package security

import (
	"crypto/rand"
	"golang.org/x/crypto/scrypt"
	env "github.com/Julia-Marcal/reusable-api/services/env"
)

func DeriveKey(password string, salt []byte) ([]byte, []byte, error) {
    if salt == nil {
        salt = make([]byte, 32)
        if _, err := rand.Read(salt); err != nil {
            return nil, nil, err
        }
    }

    password_byte := []byte(password)
    key, err := scrypt.Key(password_byte, salt, 1048576, 8, 1, 32)
    if err != nil {
        return nil, nil, err
    }

    return key, salt, nil
}



func LoginSystem(password string, password_db string) (bool, error) {
	salt := env.SetSalt()

	password_compare, _,err := DeriveKey(password, salt)
	if err != nil {
		return false, err
	}

	password_compare_str := string(password_compare)
	if (password_compare_str == password_db) {
		return true, nil
	}

	return false, nil
}
