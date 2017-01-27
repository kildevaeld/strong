package strong

import (
	"encoding/json"
	"errors"

	"github.com/kildevaeld/dict"
)

type HttpError struct {
	code    int
	message string
}

func (self *HttpError) Error() string {
	return self.message
}

func (self *HttpError) Message() string {
	return self.message
}

func (self *HttpError) Code() int {
	return self.code
}

func (self *HttpError) MarshalJSON() ([]byte, error) {
	return json.Marshal(dict.Map{
		"code":    self.code,
		"message": self.message,
	})
}

func NewHTTPError(code int, msg ...string) error {
	m := StatusText(code)
	if len(msg) != 0 {
		m = msg[0]
	}

	return &HttpError{code, m}
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
