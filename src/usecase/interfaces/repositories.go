package interfaces

import "domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindByName(string) ([]domain.User, error)
	FindAll() ([]domain.User, error)
}
