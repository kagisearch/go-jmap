package addressbook

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// Create, update or destroy AddressBooks
// https://jmap.io/spec-contacts.html#addressbook-set
type Set struct {
	Account jmap.ID `json:"accountId,omitempty"`

	IfInState string `json:"ifInState,omitempty"`

	Create map[jmap.ID]*AddressBook `json:"create,omitempty"`

	Update map[jmap.ID]jmap.Patch `json:"update,omitempty"`

	Destroy []jmap.ID `json:"destroy,omitempty"`

	OnDestroyRemoveContents bool `json:"onDestroyRemoveContents,omitempty"`

	OnSuccessSetIsDefault jmap.ID `json:"onSuccessSetIsDefault,omitempty"`
}

func (m *Set) Name() string { return "AddressBook/set" }

func (m *Set) Requires() []jmap.URI { return []jmap.URI{mail.ContactsURI} }

type SetResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	OldState string `json:"oldState,omitempty"`

	NewState string `json:"newState,omitempty"`

	Created map[jmap.ID]*AddressBook `json:"created,omitempty"`

	Updated map[jmap.ID]*AddressBook `json:"updated,omitempty"`

	Destroyed []jmap.ID `json:"destroyed,omitempty"`

	NotCreated map[jmap.ID]*jmap.SetError `json:"notCreated,omitempty"`

	NotUpdated map[jmap.ID]*jmap.SetError `json:"notUpdated,omitempty"`

	NotDestroyed map[jmap.ID]*jmap.SetError `json:"notDestroyed,omitempty"`
}

func newSetResponse() jmap.MethodResponse { return &SetResponse{} }
