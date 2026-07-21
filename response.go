package jmap

import "fmt"

type Response struct {
	// Responses are the responses to the request, in the same order that
	// the request was made
	Responses []*Invocation `json:"methodResponses"`

	// A map of client-specified ID to server-assigned ID
	CreatedIDs map[ID]ID `json:"createdIds,omitempty"`

	// SessionState is the current state of the Session
	SessionState string `json:"sessionState"`
}

type NotFoundError[T any] struct {
	ID CallID
}

func (e NotFoundError[T]) Error() string {
	var v, _ = interface{}(nil).(T)
	return fmt.Sprintf("missing response of type %T for id=%q", v, e.ID)
}

func ResponseByCallID[T any](resp *Response, id CallID) (T, error) {
	var zero T
	for _, r := range resp.Responses {
		if r.CallID == id {
			if args, ok := r.Args.(T); ok {
				return args, nil
			}
			// A method-level failure replaces the expected response with an
			// error invocation under the same call id. Return the server's
			// actual error rather than the opaque NotFoundError below.
			if methodErr, ok := r.Args.(*MethodError); ok {
				return zero, methodErr
			}
		}
	}
	return zero, NotFoundError[T]{ID: id}
}
