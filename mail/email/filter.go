package email

import (
	"encoding/json"
	"time"

	"git.sr.ht/~rockorager/go-jmap"
)

type Filter interface {
	implementsFilter()
}

type FilterOperator struct {
	Operator jmap.Operator `json:"operator,omitempty"`

	Conditions []Filter `json:"conditions,omitempty"`
}

func (fo *FilterOperator) implementsFilter() {}

// Email query condition that can be compounded with FilterOperator
// https://www.rfc-editor.org/rfc/rfc8621.html#section-4.4.1
type FilterCondition struct {
	InMailbox jmap.ID `json:"inMailbox,omitempty"`

	InMailboxOtherThan []jmap.ID `json:"inMailboxOtherThan,omitempty"`

	Before *time.Time `json:"before,omitempty"`

	After *time.Time `json:"after,omitempty"`

	MinSize uint64 `json:"minSize,omitempty"`

	MaxSize uint64 `json:"maxSize,omitempty"`

	AllInThreadHaveKeyword Keyword `json:"allInThreadHaveKeyword,omitempty"`

	SomeInThreadHaveKeyword Keyword `json:"someInThreadHaveKeyword,omitempty"`

	NoneInThreadHaveKeyword Keyword `json:"noneInThreadHaveKeyword,omitempty"`

	HasKeyword Keyword `json:"hasKeyword,omitempty"`

	NotKeyword Keyword `json:"notKeyword,omitempty"`

	HasAttachment bool `json:"hasAttachment,omitempty"`

	Text string `json:"text,omitempty"`

	From string `json:"from,omitempty"`

	To string `json:"to,omitempty"`

	Cc string `json:"cc,omitempty"`

	Bcc string `json:"bcc,omitempty"`

	Subject string `json:"subject,omitempty"`

	Body string `json:"body,omitempty"`

	Header []string `json:"header,omitempty"`

	HasSMIME bool `json:"hasSmime,omitempty"`

	HasVerifiedSMIME bool `json:"hasVerifiedSmime,omitempty"`

	HasVerifiedSMIMEAtDelivery bool `json:"hasVerifiedSmimeAtDelivery,omitempty"`
}

func (fc *FilterCondition) implementsFilter() {}

func (fc *FilterCondition) MarshalJSON() ([]byte, error) {
	if fc.Before != nil && fc.Before.Location() != time.UTC {
		utc := fc.Before.UTC()
		fc.Before = &utc
	}
	if fc.After != nil && fc.After.Location() != time.UTC {
		utc := fc.After.UTC()
		fc.After = &utc
	}
	// create a type alias to avoid infinite recursion
	type Alias FilterCondition
	return json.Marshal((*Alias)(fc))
}
