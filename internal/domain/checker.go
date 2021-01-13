package domain

type ExistenceChecker interface {
	CheckExistence(id string) bool
}
