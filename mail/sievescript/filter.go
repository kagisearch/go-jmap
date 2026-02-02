package sievescript

import "git.sr.ht/~rockorager/go-jmap"

type Filter interface {
	implementsFilter()
}

type FilterOperator struct {
	Operator jmap.Operator `json:"operator,omitempty"`

	Conditions []Filter `json:"conditions,omitempty"`
}

func (fo *FilterOperator) implementsFilter() {}

// Filter conditions for SieveScript queries
// https://www.ietf.org/archive/id/draft-ietf-jmap-sieve-22.html#section-2.5
type FilterCondition struct {
	Name string `json:"name,omitempty"`

	IsActive *bool `json:"isActive,omitempty"`
}

func (fc *FilterCondition) implementsFilter() {}
