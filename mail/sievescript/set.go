package sievescript

import "git.sr.ht/~rockorager/go-jmap"

// Create, update, and destroy SieveScript objects
// https://www.ietf.org/archive/id/draft-ietf-jmap-sieve-22.html#section-2.4
type Set struct {
	Account jmap.ID `json:"accountId,omitempty"`

	IfInState string `json:"ifInState,omitempty"`

	Create map[jmap.ID]*SieveScript `json:"create,omitempty"`

	Update map[jmap.ID]jmap.Patch `json:"update,omitempty"`

	Destroy []jmap.ID `json:"destroy,omitempty"`

	OnSuccessActivateScript *jmap.ID `json:"onSuccessActivateScript,omitempty"`

	OnSuccessDeactivateScript bool `json:"onSuccessDeactivateScript,omitempty"`
}

func (m *Set) Name() string { return "SieveScript/set" }

func (m *Set) Requires() []jmap.URI { return []jmap.URI{URI} }

type SetResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	OldState string `json:"oldState,omitempty"`

	NewState string `json:"newState,omitempty"`

	Created map[jmap.ID]*SieveScript `json:"created,omitempty"`

	Updated map[jmap.ID]*SieveScript `json:"updated,omitempty"`

	Destroyed []jmap.ID `json:"destroyed,omitempty"`

	NotCreated map[jmap.ID]*jmap.SetError `json:"notCreated,omitempty"`

	NotUpdated map[jmap.ID]*jmap.SetError `json:"notUpdated,omitempty"`

	NotDestroyed map[jmap.ID]*jmap.SetError `json:"notDestroyed,omitempty"`
}

func newSetResponse() jmap.MethodResponse { return &SetResponse{} }
