package auth

import (
	"gopkg.in/go-playground/validator.v9"
	"log"
	"sync"

	"github.com/pquerna/otp/totp"
)

type User struct {
	Email     string `json:"email" validate:"required,email,endswith=trustingsocial.com"`
	OtpSecret string `json:"otp_secret" validate:"required"`
}

type secretStore struct {
	sync.RWMutex
	secrets map[string]string
}

var store *secretStore
var lock sync.RWMutex

func initStore() error {
	lock.Lock()
	defer lock.Unlock()

	if store != nil {
		return nil
	}

	store = &secretStore{
		secrets: make(map[string]string),
	}

	data, err := fetchPritunlData()
	if err != nil {
		return err
	}
	store.update(data)

	return nil
}

func (s *secretStore) update(users []User) {
	s.Lock()
	defer s.Unlock()

	if s.secrets == nil {
		s.secrets = make(map[string]string)
	}

	validate := validator.New()

	for _, u := range users {

		err := validate.Struct(u)
		if u.Email != "" && u.OtpSecret != "" && err == nil {
			s.secrets[u.Email] = u.OtpSecret
		}
	}
}

func (s *secretStore) isValid(e, t string) bool {
	if e == "" || t == "" {
		return false
	}

	s.RLock()
	defer s.RUnlock()

	secret, ok := s.secrets[e]
	if !ok {
		log.Printf("otp-validation-error: secret not found for %s", e)
		return false
	}

	return totp.Validate(t, secret)
}
