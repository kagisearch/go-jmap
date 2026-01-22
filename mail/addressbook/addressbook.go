package addressbook

import "git.sr.ht/~rockorager/go-jmap"

func init() {
	jmap.RegisterMethod("AddressBook/get", newGetResponse)
	jmap.RegisterMethod("AddressBook/changes", newChangesResponse)
	jmap.RegisterMethod("AddressBook/set", newSetResponse)
}

// A named collection of ContactCards
// https://jmap.io/spec-contacts.html#addressbooks
type AddressBook struct {
	ID jmap.ID `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	SortOrder uint64 `json:"sortOrder,omitempty"`

	IsDefault bool `json:"isDefault,omitempty"`

	IsSubscribed bool `json:"isSubscribed,omitempty"`

	ShareWith map[jmap.ID]*AddressBookRights `json:"shareWith,omitempty"`

	MyRights *AddressBookRights `json:"myRights,omitempty"`
}

// Access Control Lists (ACLs) for AddressBooks
type AddressBookRights struct {
	MayRead bool `json:"mayRead,omitempty"`

	MayWrite bool `json:"mayWrite,omitempty"`

	MayAdmin bool `json:"mayAdmin,omitempty"`

	MayDelete bool `json:"mayDelete,omitempty"`
}
