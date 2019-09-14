package interfaces

import "github.com/nakabonne/cleanarchitecture-sample/pkg/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindByName(string) ([]domain.User, error)
	FindAll() ([]domain.User, error)
}
