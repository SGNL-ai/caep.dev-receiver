package ssf_events

type CredentialType uint64

const (
	Password CredentialType = iota
	Pin
	X509
	Fido2_platform
	Fido2_roaming
	Fido_u2f
	Verifiable_credential
	Phone_voice
	Phone_sms
	App
)

type ChangeType uint64

const (
	Create ChangeType = iota
	Revoked
	Update
	Delete
)

var CredentialTypeEnumMap = map[uint64]CredentialType{
	0: Password,
	1: Pin,
	2: X509,
	3: Fido2_platform,
	4: Fido2_roaming,
	5: Fido_u2f,
	6: Verifiable_credential,
	7: Phone_voice,
	8: Phone_sms,
	9: App,
}

var ChangeTypeEnumMap = map[uint64]ChangeType{
	0: Create,
	1: Revoked,
	2: Update,
	3: Delete,
}

// The credential change event is a CAEP Event, defined here:
// https://openid.net/specs/openid-caep-specification-1_0-ID1.html#rfc.section.3.1
type CredentialChangeEvent struct {
	// Json defines the raw JSON of the CAEP Event. Used if
	// a developer wants greater control over all the attributes
	// of the CAEP Event
	Json map[string]interface{}

	// SubjectFormat defines the format of the subject event.
	//
	// See: https://datatracker.ietf.org/doc/html/draft-ietf-secevent-subject-identifiers
	// for an overview of available subject formats for CAEP Events
	Format SubjectFormat

	// Subject defines the subject that the CAEP Event applies to.
	//
	// See your transmitter's specification for the exact format
	// of the Subject
	Subject map[string]interface{}

	// EventTimestamp defines the timestamp of the CAEP Event in
	// Unix time (seconds since January 1, 1970 UTC)
	EventTimestamp int64

	// CredentialType defines the type of credential of the CAEP Event that has been modified/removed.
	// See https://openid.net/specs/openid-caep-specification-1_0.html#rfc.section.3.3.1 for the options for this field
	CredentialType CredentialType

	// CredentialType defines the type of modification/deletion towards the credential of the CAEP Event.
	// See https://openid.net/specs/openid-caep-specification-1_0.html#rfc.section.3.3.1 for the options for this field
	ChangeType ChangeType
}

func (event *CredentialChangeEvent) GetEventUri() string {
	return "https://schemas.openid.net/secevent/caep/event-type/credential-change"
}

func (event *CredentialChangeEvent) GetSubjectFormat() SubjectFormat {
	return event.Format
}

func (event *CredentialChangeEvent) GetSubject() map[string]interface{} {
	return event.Subject
}

func (event *CredentialChangeEvent) GetTimestamp() int64 {
	return event.EventTimestamp
}

func (event *CredentialChangeEvent) GetCredentialType() CredentialType {
	return event.CredentialType
}

func (event *CredentialChangeEvent) GetChangeType() ChangeType {
	return event.ChangeType
}
