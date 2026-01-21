package contactcard

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

// ContactCard query condition that can be compounded with FilterOperator
// https://jmap.io/spec-contacts.html#contactcard-query-filtering
type FilterCondition struct {
	InAddressBook jmap.ID `json:"inAddressBook,omitempty"`

	UID string `json:"uid,omitempty"`

	HasMember string `json:"hasMember,omitempty"`

	Kind string `json:"kind,omitempty"`

	CreatedBefore *time.Time `json:"createdBefore,omitempty"`

	CreatedAfter *time.Time `json:"createdAfter,omitempty"`

	UpdatedBefore *time.Time `json:"updatedBefore,omitempty"`

	UpdatedAfter *time.Time `json:"updatedAfter,omitempty"`

	Text string `json:"text,omitempty"`

	Name string `json:"name,omitempty"`

	NameGiven string `json:"name/given,omitempty"`

	NameSurname string `json:"name/surname,omitempty"`

	NameSurname2 string `json:"name/surname2,omitempty"`

	NickName string `json:"nickName,omitempty"`

	Organization string `json:"organization,omitempty"`

	Email string `json:"email,omitempty"`

	Phone string `json:"phone,omitempty"`

	OnlineService string `json:"onlineService,omitempty"`

	Address string `json:"address,omitempty"`

	Note string `json:"note,omitempty"`
}

func (fc *FilterCondition) implementsFilter() {}

func (fc *FilterCondition) MarshalJSON() ([]byte, error) {
	if fc.CreatedBefore != nil && fc.CreatedBefore.Location() != time.UTC {
		utc := fc.CreatedBefore.UTC()
		fc.CreatedBefore = &utc
	}
	if fc.CreatedAfter != nil && fc.CreatedAfter.Location() != time.UTC {
		utc := fc.CreatedAfter.UTC()
		fc.CreatedAfter = &utc
	}
	if fc.UpdatedBefore != nil && fc.UpdatedBefore.Location() != time.UTC {
		utc := fc.UpdatedBefore.UTC()
		fc.UpdatedBefore = &utc
	}
	if fc.UpdatedAfter != nil && fc.UpdatedAfter.Location() != time.UTC {
		utc := fc.UpdatedAfter.UTC()
		fc.UpdatedAfter = &utc
	}
	// create a type alias to avoid infinite recursion
	type Alias FilterCondition
	return json.Marshal((*Alias)(fc))
}
