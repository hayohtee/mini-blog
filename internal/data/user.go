package data

import (
	"errors"
	"time"

	"github.com/hayohtee/mini-blog/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  password  `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type password struct {
	plainText string
	hash      []byte
}

// Set calculates the bcrypt hash of a plaintext password, and stores both the
// hash and the plaintext versions in the struct.
func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}
	p.plainText = plaintextPassword
	p.hash = hash
	return nil
}

// Matches checks whether the provided plaintext password matches the hashed
// password stored in the struct, returning true if it matches and false otherwise.
func (p *password) Matches(plaintextPassword string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword)); err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, nil
		}
	}
	return true, nil
}

// ValidatePasswordPlaintext ensures that the provided password is not empty,
// and it is between 8 and 72 bytes long.
func ValidatePasswordPlainText(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}

// ValidateUser perform validation on the user struct fields,
// ensuring that values are specified and are within range.
func ValidateUser(v *validator.Validator, user User) {
	v.Check(user.Username != "", "username", "must be provided")
	v.Check(len(user.Username) <= 500, "username", "must not be more than 500 bytes long")

	if user.Password.plainText != "" {
		ValidatePasswordPlainText(v, user.Password.plainText)
	}

	if user.Password.hash == nil {
		panic("missing password hash for user")
	}
}
