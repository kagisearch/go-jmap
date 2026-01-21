package contactcard

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// Get changes to ContactCards since a given state
// https://jmap.io/spec-contacts.html#contactcard-changes
type Changes struct {
	Account jmap.ID `json:"accountId,omitempty"`

	SinceState string `json:"sinceState,omitempty"`

	MaxChanges uint64 `json:"maxChanges,omitempty"`
}

func (m *Changes) Name() string { return "ContactCard/changes" }

func (m *Changes) Requires() []jmap.URI { return []jmap.URI{mail.ContactsURI} }

type ChangesResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	OldState string `json:"oldState,omitempty"`

	NewState string `json:"newState,omitempty"`

	HasMoreChanges bool `json:"hasMoreChanges,omitempty"`

	Created []jmap.ID `json:"created,omitempty"`

	Updated []jmap.ID `json:"updated,omitempty"`

	Destroyed []jmap.ID `json:"destroyed,omitempty"`
}

func newChangesResponse() jmap.MethodResponse { return &ChangesResponse{} }
