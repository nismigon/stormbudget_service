package errors

import "fmt"

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("%s", v.Message)
}
