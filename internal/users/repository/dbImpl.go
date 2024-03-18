package repository

type dbImpl struct{}

func New() Repo {
	return &dbImpl{}
}
