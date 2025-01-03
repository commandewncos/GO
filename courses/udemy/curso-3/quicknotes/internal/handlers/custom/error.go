package custom

type StatusError struct {
	error
	status int
}

func (s StatusError) StatusCustomHttpError() int {
	return s.status
}

func WithStatusError(e error, s int) error {
	return StatusError{
		error:  e,
		status: s,
	}
}
