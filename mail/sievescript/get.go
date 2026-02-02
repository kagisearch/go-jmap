package sievescript

import "git.sr.ht/~rockorager/go-jmap"

// Get SieveScript objects
// https://www.ietf.org/archive/id/draft-ietf-jmap-sieve-22.html#section-2.3
type Get struct {
	Account jmap.ID `json:"accountId,omitempty"`

	IDs []jmap.ID `json:"ids,omitempty"`

	Properties []string `json:"properties,omitempty"`

	ReferenceIDs *jmap.ResultReference `json:"#ids,omitempty"`

	ReferenceProperties *jmap.ResultReference `json:"#properties,omitempty"`
}

func (m *Get) Name() string { return "SieveScript/get" }

func (m *Get) Requires() []jmap.URI { return []jmap.URI{URI} }

type GetResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	State string `json:"state,omitempty"`

	List []*SieveScript `json:"list,omitempty"`

	NotFound []jmap.ID `json:"notFound,omitempty"`
}

func newGetResponse() jmap.MethodResponse { return &GetResponse{} }
