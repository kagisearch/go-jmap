package contactcard

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// Copy ContactCards from one account to another
// https://jmap.io/spec-contacts.html#contactcard-copy
type Copy struct {
	FromAccount jmap.ID `json:"fromAccountId,omitempty"`

	IfFromInState string `json:"ifFromInState,omitempty"`

	Account jmap.ID `json:"accountId,omitempty"`

	IfInState string `json:"ifInState,omitempty"`

	Create map[jmap.ID]*ContactCard `json:"create,omitempty"`

	OnSuccessDestroyOriginal bool `json:"onSuccessDestroyOriginal,omitempty"`

	DestroyFromIfInState string `json:"destroyFromIfInState,omitempty"`
}

func (m *Copy) Name() string { return "ContactCard/copy" }

func (m *Copy) Requires() []jmap.URI { return []jmap.URI{mail.ContactsURI} }

type CopyResponse struct {
	FromAccount jmap.ID `json:"fromAccountId,omitempty"`

	Account jmap.ID `json:"accountId,omitempty"`

	OldState string `json:"oldState,omitempty"`

	NewState string `json:"newState,omitempty"`

	Created map[jmap.ID]*ContactCard `json:"created,omitempty"`

	NotCreated map[jmap.ID]*jmap.SetError `json:"notCreated,omitempty"`
}

func newCopyResponse() jmap.MethodResponse { return &CopyResponse{} }
