package authentication

import "github.com/google/uuid"

type Usecase struct{}

func (u *Usecase) generateRandomSessionToken() string {
	return uuid.New().String()
}
