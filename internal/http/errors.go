package http

import "fmt"

type UnexpectedError struct {
	err error
}

func (e *UnexpectedError) Error() string {
	return fmt.Sprintf("Unexpected error: %s", e.err)
}
