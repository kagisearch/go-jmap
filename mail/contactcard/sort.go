package contactcard

import "git.sr.ht/~rockorager/go-jmap"

// ContactCard sort criteria
// https://jmap.io/spec-contacts.html#contactcard-query-sorting
type SortComparator struct {
	Property string `json:"property,omitempty"`

	IsAscending bool `json:"isAscending"`

	Collation jmap.CollationAlgo `json:"collation,omitempty"`
}
