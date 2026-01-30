package sievescript

import "git.sr.ht/~rockorager/go-jmap"

// Validate a Sieve script without storing it
// https://www.ietf.org/archive/id/draft-ietf-jmap-sieve-22.html#section-2.6
type Validate struct {
	Account jmap.ID `json:"accountId,omitempty"`

	BlobID jmap.ID `json:"blobId,omitempty"`
}

func (m *Validate) Name() string { return "SieveScript/validate" }

func (m *Validate) Requires() []jmap.URI { return []jmap.URI{URI} }

type ValidateResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	Error *jmap.SetError `json:"error"`
}

func newValidateResponse() jmap.MethodResponse { return &ValidateResponse{} }
