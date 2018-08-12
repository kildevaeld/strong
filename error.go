package strong

import (
	"encoding/json"
	"errors"
	"fmt"
)

// HttpError Represents a http error
type HttpError struct {
	status  int
	err     string
	message string
}

func (self *HttpError) Error() string {
	return self.err
}

func (self *HttpError) Message() string {
	return self.message
}

func (self *HttpError) StatusCode() int {
	return self.status
}

func (self *HttpError) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.ToMap())
}

func (self *HttpError) ToMap() map[string]interface{} {
	m := map[string]interface{}{
		"status": self.status,
		"error":  self.err,
	}

	if self.message != "" {
		m["message"] = self.message
	}
	return m
}

func NewHTTPError(code int, msg ...interface{}) error {
	m := StatusText(code)
	var msgStr string
	if len(msg) != 0 {
		var ok bool
		msgStr, ok = msg[0].(string)
		if !ok {
			panic("first argument must be a string")
		}
		if len(msg) > 1 {
			msgStr = fmt.Sprintf(msgStr, msg[1:]...)
		}

	}

	return &HttpError{code, m, msgStr}
}

var (
	ErrUnsupportedMediaType        = NewHTTPError(StatusUnsupportedMediaType)
	ErrNotFound                    = NewHTTPError(StatusNotFound)
	ErrUnauthorized                = NewHTTPError(StatusUnauthorized)
	ErrMethodNotAllowed            = NewHTTPError(StatusMethodNotAllowed)
	ErrStatusRequestEntityTooLarge = NewHTTPError(StatusRequestEntityTooLarge)
	ErrValidatorNotRegistered      = errors.New("validator not registered")
	ErrRendererNotRegistered       = errors.New("renderer not registered")
	ErrInvalidRedirectCode         = errors.New("invalid redirect status code")
	ErrCookieNotFound              = errors.New("cookie not found")
)
