// Package contacts implements JMAP for Contacts (https://jmap.io/spec-contacts.html)
package mail

import (
	"git.sr.ht/~rockorager/go-jmap"
)

// urn:ietf:params:jmap:contacts represents support for the AddressBook and
// ContactCard data types and associated API methods
const ContactsURI jmap.URI = "urn:ietf:params:jmap:contacts"

const (
	// The AddressBook event type
	AddressBookEvent jmap.EventType = "AddressBook"

	// The ContactCard event type
	ContactCardEvent jmap.EventType = "ContactCard"
)

func init() {
	jmap.RegisterCapability(&Contacts{})
}

type Contacts struct {
	MaxAddressBooksPerCard uint64 `json:"maxAddressBooksPerCard"`
	MayCreateAddressBook   bool   `json:"mayCreateAddressBook"`
}

func (c *Contacts) URI() jmap.URI { return ContactsURI }

func (c *Contacts) New() jmap.Capability { return &Contacts{} }
