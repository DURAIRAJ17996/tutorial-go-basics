package mock

import "fmt"

type CustomError struct {
	ErrorMessage string
	ErrorCode    int
}

func NewCustomError(errorMesssage string, errorCode int) CustomError {
	return CustomError{errorMesssage, errorCode}
}

func (errorObj CustomError) Error() string {
	return fmt.Sprintf("Error code: %d,: %s", errorObj.ErrorCode, errorObj.ErrorMessage)
}
