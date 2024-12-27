package myerror

type ErrDuplicate struct{}

func (e ErrDuplicate) Error() string {
	return "duplicate"
}
