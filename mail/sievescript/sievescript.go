package sievescript

import "git.sr.ht/~rockorager/go-jmap"

const URI jmap.URI = "urn:ietf:params:jmap:sieve"

func init() {
	jmap.RegisterCapability(&Capability{})
	jmap.RegisterMethod("SieveScript/get", newGetResponse)
	jmap.RegisterMethod("SieveScript/set", newSetResponse)
	jmap.RegisterMethod("SieveScript/query", newQueryResponse)
	jmap.RegisterMethod("SieveScript/validate", newValidateResponse)
}

// The SieveScript capability
// https://www.ietf.org/archive/id/draft-ietf-jmap-sieve-22.html#section-1.2
type Capability struct {
	Implementation string `json:"implementation,omitempty"`

	MaxSizeScriptName uint64 `json:"maxSizeScriptName,omitempty"`

	MaxSizeScript *uint64 `json:"maxSizeScript,omitempty"`

	MaxNumberScripts *uint64 `json:"maxNumberScripts,omitempty"`

	MaxNumberRedirects *uint64 `json:"maxNumberRedirects,omitempty"`

	SieveExtensions []string `json:"sieveExtensions,omitempty"`

	NotificationMethods []string `json:"notificationMethods,omitempty"`

	ExternalLists []string `json:"externalLists,omitempty"`
}

func (c *Capability) URI() jmap.URI { return URI }

func (c *Capability) New() jmap.Capability { return &Capability{} }

// A Sieve script
// https://www.ietf.org/archive/id/draft-ietf-jmap-sieve-22.html#section-2
type SieveScript struct {
	ID jmap.ID `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	BlobID jmap.ID `json:"blobId,omitempty"`

	IsActive bool `json:"isActive,omitempty"`
}

// SetError types specific to SieveScript
const (
	SetErrorInvalidSieve = "invalidSieve"

	SetErrorSieveIsActive = "sieveIsActive"

	SetErrorAlreadyExists = "alreadyExists"

	SetErrorTooLarge = "tooLarge"

	SetErrorOverQuota = "overQuota"

	SetErrorForbidden = "forbidden"
)
