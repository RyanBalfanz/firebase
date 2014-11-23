package firebase

type firebaseError struct {
	msg string
	err error
}

func NewFirebaseError(msg string, err error) error {
	return &firebaseError{msg: msg, err: err}
}

func (e *firebaseError) Error() string {
	return e.msg
}
