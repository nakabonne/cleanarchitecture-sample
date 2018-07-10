package domain

import "fmt"

type Users []User

type User struct {
	ID    int
	Name  string
	Email string
}
