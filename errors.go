package gonextcloud

import (
	"fmt"
	"strings"
)

// APIError contains the returned error code and message from the Nextcloud's API
type APIError struct {
	Code    int
	Message string
}

//errorFromMeta return a types.APIError from the Response's types.meta
func errorFromMeta(meta meta) *APIError {
	return &APIError{
		meta.Statuscode,
		meta.Message,
	}
}

// Error return the types.APIError string
func (e *APIError) Error() string {
	return fmt.Sprintf("%d : %s", e.Code, e.Message)
}

// UpdateError contains the user's field and corresponding error
type UpdateError struct {
	Field string
	Error error
}

// UserUpdateError contains the errors resulting from a UserUpdate or a UserCreateFull call
type UserUpdateError struct {
	Errors map[string]error
}

func (e *UserUpdateError) Error() string {
	var errors []string
	for k, e := range e.Errors {
		errors = append(errors, fmt.Sprintf("%s: %v", k, e))
	}
	return strings.Join(errors, ", ")
}

//newUpdateError returns an UpdateError based on an UpdateError channel
func newUpdateError(errors chan *UpdateError) *UserUpdateError {
	ue := UserUpdateError{map[string]error{}}
	for e := range errors {
		if e != nil {
			ue.Errors[e.Field] = e.Error
		}
	}
	if len(ue.Errors) > 0 {
		return &ue
	}
	return nil
}
