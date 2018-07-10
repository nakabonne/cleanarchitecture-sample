package usecase

import (
	"github.com/nakabonne/cleanarc-sample/src/domain"
	"github.com/nakabonne/cleanarc-sample/src/usecase/interfaces"
)

type UserInteractor struct {
	UserRepository interfaces.UserRepository
	Logger         Logger
}

func (i *UserInteractor) Add(u domain.User) (int, error) {
	return i.UserRepository.Store(u)
}
