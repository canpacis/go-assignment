package errs

import "encoding/json"

// Error represents an HTTP error with a status code and user friendly message.
type Error struct {
	Status  int
	Message string
	// Optional underlying error.
	Err error
}

// Error method implements the error interface.
func (e *Error) Error() string {
	return e.Message
}

// Unwrap method implements the errors package's unwrapper interface.
func (e *Error) Unwrap() error {
	return e.Err
}

// MarshalJSON method implements the json package's Marshaler interface.
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{"error": e.Message})
}
