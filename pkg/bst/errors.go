package bst

type ErrInvalidTraversalType struct {
	msg string
}

func (e ErrInvalidTraversalType) Error() string {
	return e.msg
}
