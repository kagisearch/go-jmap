package contactcard

import (
	"time"

	"git.sr.ht/~rockorager/go-jmap"
)

func init() {
	jmap.RegisterMethod("ContactCard/get", newGetResponse)
	jmap.RegisterMethod("ContactCard/changes", newChangesResponse)
	jmap.RegisterMethod("ContactCard/query", newQueryResponse)
	jmap.RegisterMethod("ContactCard/queryChanges", newQueryChangesResponse)
	jmap.RegisterMethod("ContactCard/set", newSetResponse)
	jmap.RegisterMethod("ContactCard/copy", newCopyResponse)
}

// A ContactCard contains information about a person, company, or other entity,
// or represents a group of such entities. It is a JSCard object with additional
// JMAP-specific properties.
// https://jmap.io/spec-contacts.html#contactcards
type ContactCard struct {
	ID jmap.ID `json:"id,omitempty"`

	AddressBookIDs map[jmap.ID]bool `json:"addressBookIds,omitempty"`

	// JSCard properties - core fields
	UID string `json:"uid,omitempty"`

	ProdID string `json:"prodId,omitempty"`

	Created *time.Time `json:"created,omitempty"`

	Updated *time.Time `json:"updated,omitempty"`

	Kind string `json:"kind,omitempty"`

	RelatedTo map[string]*Relation `json:"relatedTo,omitempty"`

	Name *Name `json:"name,omitempty"`

	FullName *FullName `json:"fullName,omitempty"`

	NickNames map[string]*NickName `json:"nickNames,omitempty"`

	Organizations map[string]*Organization `json:"organizations,omitempty"`

	Titles map[string]*Title `json:"titles,omitempty"`

	Emails map[string]*EmailAddress `json:"emails,omitempty"`

	Phones map[string]*Phone `json:"phones,omitempty"`

	OnlineServices map[string]*OnlineService `json:"onlineServices,omitempty"`

	Addresses map[string]*Address `json:"addresses,omitempty"`

	PreferredContactLanguages map[string]*ContactLanguage `json:"preferredContactLanguages,omitempty"`

	Anniversaries map[string]*Anniversary `json:"anniversaries,omitempty"`

	PersonalInfo map[string]*PersonalInfo `json:"personalInfo,omitempty"`

	Notes map[string]*Note `json:"notes,omitempty"`

	Categories map[string]bool `json:"categories,omitempty"`

	Photos map[string]*Media `json:"photos,omitempty"`

	Links map[string]*Link `json:"links,omitempty"`

	Keywords map[string]bool `json:"keywords,omitempty"`

	LocalizedStrings map[string]*LocalizedString `json:"localizedStrings,omitempty"`

	Members map[string]bool `json:"members,omitempty"`
}

// Relation represents a relation to another entity
type Relation struct {
	Relation map[string]bool `json:"relation,omitempty"`
}

// Name represents structured name information
type Name struct {
	Components []*NameComponent `json:"components,omitempty"`

	IsOrdered bool `json:"isOrdered,omitempty"`

	DefaultSeparator string `json:"defaultSeparator,omitempty"`

	SortAs map[string]string `json:"sortAs,omitempty"`
}

// NameComponent represents a component of a name
type NameComponent struct {
	Kind string `json:"kind,omitempty"`

	Value string `json:"value,omitempty"`

	Phonetic string `json:"phonetic,omitempty"`
}

// FullName represents the full name as a single string
type FullName struct {
	Value string `json:"value,omitempty"`

	LocalizedValues map[string]string `json:"localizedValues,omitempty"`
}

// NickName represents a nickname
type NickName struct {
	Name string `json:"name,omitempty"`

	Contexts map[string]bool `json:"contexts,omitempty"`

	Pref uint64 `json:"pref,omitempty"`
}

// Organization represents an organization
type Organization struct {
	Name string `json:"name,omitempty"`

	Units []*OrgUnit `json:"units,omitempty"`

	SortAs string `json:"sortAs,omitempty"`

	Contexts map[string]bool `json:"contexts,omitempty"`

	Pref uint64 `json:"pref,omitempty"`
}

// OrgUnit represents an organizational unit
type OrgUnit struct {
	Name string `json:"name,omitempty"`

	SortAs string `json:"sortAs,omitempty"`
}

// Title represents a job title or role
type Title struct {
	Title string `json:"title,omitempty"`

	Organization string `json:"organization,omitempty"`

	Contexts map[string]bool `json:"contexts,omitempty"`

	Pref uint64 `json:"pref,omitempty"`
}

// EmailAddress represents an email address
type EmailAddress struct {
	Address string `json:"address,omitempty"`

	Contexts map[string]bool `json:"contexts,omitempty"`

	Pref uint64 `json:"pref,omitempty"`

	Label string `json:"label,omitempty"`
}

// Phone represents a phone number
type Phone struct {
	Number string `json:"number,omitempty"`

	Features map[string]bool `json:"features,omitempty"`

	Contexts map[string]bool `json:"contexts,omitempty"`

	Pref uint64 `json:"pref,omitempty"`

	Label string `json:"label,omitempty"`
}

// OnlineService represents an online service or instant messaging address
type OnlineService struct {
	Service string `json:"service,omitempty"`

	URI string `json:"uri,omitempty"`

	User string `json:"user,omitempty"`

	Contexts map[string]bool `json:"contexts,omitempty"`

	Pref uint64 `json:"pref,omitempty"`

	Label string `json:"label,omitempty"`
}

// Address represents a physical address
type Address struct {
	Contexts map[string]bool `json:"contexts,omitempty"`

	Pref uint64 `json:"pref,omitempty"`

	Label string `json:"label,omitempty"`

	FullAddress *FullAddress `json:"fullAddress,omitempty"`

	Street []*StreetComponent `json:"street,omitempty"`

	Locality string `json:"locality,omitempty"`

	Region string `json:"region,omitempty"`

	Country string `json:"country,omitempty"`

	Postcode string `json:"postcode,omitempty"`

	CountryCode string `json:"countryCode,omitempty"`

	Coordinates string `json:"coordinates,omitempty"`

	TimeZone string `json:"timeZone,omitempty"`
}

// FullAddress represents the full address as a single string
type FullAddress struct {
	Value string `json:"value,omitempty"`

	LocalizedValues map[string]string `json:"localizedValues,omitempty"`
}

// StreetComponent represents a component of a street address
type StreetComponent struct {
	Kind string `json:"kind,omitempty"`

	Value string `json:"value,omitempty"`
}

// ContactLanguage represents a preferred language for contact
type ContactLanguage struct {
	Language string `json:"language,omitempty"`

	Contexts map[string]bool `json:"contexts,omitempty"`

	Pref uint64 `json:"pref,omitempty"`
}

// Anniversary represents an anniversary or special date
type Anniversary struct {
	Kind string `json:"kind,omitempty"`

	Date string `json:"date,omitempty"`

	Place *Address `json:"place,omitempty"`
}

// PersonalInfo represents personal information
type PersonalInfo struct {
	Kind string `json:"kind,omitempty"`

	Value string `json:"value,omitempty"`

	Level string `json:"level,omitempty"`

	ListAs string `json:"listAs,omitempty"`
}

// Note represents a note or comment
type Note struct {
	Note string `json:"note,omitempty"`

	Created *time.Time `json:"created,omitempty"`

	Author *Author `json:"author,omitempty"`
}

// Author represents the author of a note
type Author struct {
	Name string `json:"name,omitempty"`

	URI string `json:"uri,omitempty"`
}

// Media represents a media file (photo, logo, sound)
type Media struct {
	URI string `json:"uri,omitempty"`

	BlobID jmap.ID `json:"blobId,omitempty"`

	MediaType string `json:"mediaType,omitempty"`

	Contexts map[string]bool `json:"contexts,omitempty"`

	Pref uint64 `json:"pref,omitempty"`
}

// Link represents a link or URL
type Link struct {
	URI string `json:"uri,omitempty"`

	Contexts map[string]bool `json:"contexts,omitempty"`

	Pref uint64 `json:"pref,omitempty"`

	Label string `json:"label,omitempty"`

	MediaType string `json:"mediaType,omitempty"`
}

// LocalizedString represents a localized string value
type LocalizedString struct {
	Value string `json:"value,omitempty"`

	Language string `json:"language,omitempty"`

	Localizations map[string]string `json:"localizations,omitempty"`
}
