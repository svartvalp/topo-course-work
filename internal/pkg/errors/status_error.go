package errors

type StatusError struct {
	Status int
	Err    string
}

func (s *StatusError) Error() string {
	return s.Err
}

func Status(err error) int {
	if val, ok := err.(*StatusError); ok {
		return val.Status
	}
	return 0
}

func NewStatusError(status int, err string) *StatusError {
	return &StatusError{
		Status: status,
		Err:    err,
	}
}
