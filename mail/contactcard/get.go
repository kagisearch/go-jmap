package contactcard

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// Get ContactCard details
// https://jmap.io/spec-contacts.html#contactcard-get
type Get struct {
	Account jmap.ID `json:"accountId,omitempty"`

	IDs []jmap.ID `json:"ids,omitempty"`

	Properties []string `json:"properties,omitempty"`

	ReferenceIDs *jmap.ResultReference `json:"#ids,omitempty"`

	ReferenceProperties *jmap.ResultReference `json:"#properties,omitempty"`
}

func (m *Get) Name() string { return "ContactCard/get" }

func (m *Get) Requires() []jmap.URI { return []jmap.URI{mail.ContactsURI} }

type GetResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	State string `json:"state,omitempty"`

	List []*ContactCard `json:"list,omitempty"`

	NotFound []jmap.ID `json:"notFound,omitempty"`
}

func newGetResponse() jmap.MethodResponse { return &GetResponse{} }
