package jmap

import "fmt"

// A RequestError occurs when there is an error with the HTTP request
type RequestError struct {
	// The type of request error, eg "urn:ietf:params:jmap:error:limit"
	Type string `json:"type"`

	// The HTTP status code of the response
	Status int `json:"status"`

	// The description of the error
	Detail string `json:"detail"`

	// If the error is of type ErrLimit, Limit will contain the name of the
	// limit the request would have exceeded
	Limit *string `json:"limit,omitempty"`
}

func (e *RequestError) Error() string {
	// A server may send any subset of these; fall back through them so the error
	// is never empty and always names its status code when one is known.
	msg := e.Detail
	if msg == "" {
		msg = e.Type
	}
	if e.Limit != nil {
		msg = fmt.Sprintf("%s: %s", msg, *e.Limit)
	}
	if e.Status == 0 {
		return msg
	}
	if msg == "" {
		return fmt.Sprintf("HTTP %d", e.Status)
	}
	return fmt.Sprintf("HTTP %d: %s", e.Status, msg)
}

// A MethodError is returned when an error occurred while the server was
// processing a method. Instead of the Response of that method, a MethodError
// invocation will be in it's place
type MethodError struct {
	// The type of error that occurred. Always present
	Type string `json:"type,omitempty"`

	// Description is available on some method errors (notably,
	// invalidArguments)
	Description *string `json:"description,omitempty"`
}

func (m *MethodError) Error() string {
	if m.Description != nil {
		return fmt.Sprintf("%s: %s", m.Type, *m.Description)
	}
	return m.Type
}

func newMethodError() MethodResponse { return &MethodError{} }

// A SetError is returned in set calls for individual record changes
type SetError struct {
	// The type of SetError
	Type string `json:"type,omitempty"`

	// A description of the error to help with debugging that includes an
	// explanation of what the problem was. This is a non-localised string
	// and is not intended to be shown directly to end users.
	Description *string `json:"description,omitempty"`

	// Properties is available on InvalidProperties SetErrors and lists the
	// individual properties were
	Properties *[]string `json:"properties,omitempty"`
}
