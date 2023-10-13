package server

import "fmt"

type ErrorAnswersIsNil struct{}

func (ErrorAnswersIsNil) Error() string {
	return "Answer parameter cannot be nil"
}

type ErrorInvalidOption string

func (e ErrorInvalidOption) Error() string {
	return fmt.Sprintf("Invalid option: %s", string(e))
}
