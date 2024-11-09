package repositories

type RepositoryError struct {
	error
}

func newRepositoryError(e error) error {
	return &RepositoryError{error: e}
}
