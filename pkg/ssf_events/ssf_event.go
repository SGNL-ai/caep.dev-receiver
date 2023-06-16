package ssf_events

import (
	"errors"
	"fmt"
	"strconv"
)

type EventType int

const (
	SessionRevoked EventType = iota
)

type SubjectFormat int

const (
	Account SubjectFormat = iota
	Email
	IssuerAndSubject
	Opaque
	PhoneNumber
	DecentralizedIdentifier
	UniqueResourceIdentifier
	Aliases
	ComplexSubject
)

const AccountSubjectFormat = "account"
const EmailSubjectFormat = "email"
const IssuerAndSubjectFormat = "iss_sub"
const OpaqueSubjectFormat = "opaque"
const PhoneNumberSubjectFormat = "phone_number"
const DecentralizedIdentifierSubjectFormat = "did"
const UniqueResourceIdentifierSubjectFormat = "uri"
const AliasesSubjectFormat = "aliases"

// Represents the interface that all SSF Events should implement
//
// See the SessionRevokedEvent (./events/session_revoked_event.go)
// for an example
type SsfEvent interface {
	// Returns the Event URI for the given event
	GetEventUri() string

	// Returns the format of the event's subject
	GetSubjectFormat() SubjectFormat

	// Returns the subject of the event
	GetSubject() map[string]interface{}

	// Returns the Unix timestamp of the event
	GetTimestamp() int64
}

var EventUri = map[EventType]string{
	SessionRevoked: "https://schemas.openid.net/secevent/caep/event-type/session-revoked",
}

var EventEnum = map[string]EventType{
	"https://schemas.openid.net/secevent/caep/event-type/session-revoked": SessionRevoked,
}

// Takes an event subject from the JSON of an SSF Event, and converts it into the matching struct for that event
func EventStructFromEvent(eventUri string, eventSubject interface{}, claimsJson map[string]interface{}) (SsfEvent, error) {
	eventEnum := EventEnum[eventUri]

	// Add more Ssf Events as desired
	switch eventEnum {
	case SessionRevoked:
		subjectAttributes, ok := eventSubject.(map[string]interface{})
		timestamp, err := strconv.ParseInt(subjectAttributes["timestamp"].(string), 10, 64)
		if !ok || err != nil {
			return nil, errors.New("Unable to parse event subject")
		}

		format, err := GetSubjectFormat(subjectAttributes["subject"].(map[string]interface{}))
		if err != nil {
			return nil, err
		}

		event := SessionRevokedEvent{
			Json:           claimsJson,
			Format:         format,
			Subject:        subjectAttributes["subject"].(map[string]interface{}),
			EventTimestamp: timestamp,
		}
		return &event, nil
	default:
		return nil, errors.New("No matching events")
	}
}

func GetSubjectFormat(subject map[string]interface{}) (SubjectFormat, error) {
	format, formatFound := subject["format"]
	formatString := fmt.Sprintf("%v", format)
	if !formatFound {
		return ComplexSubject, nil
	}

	switch formatString {
	case AccountSubjectFormat:
		return Account, nil
	case EmailSubjectFormat:
		return Email, nil
	case IssuerAndSubjectFormat:
		return IssuerAndSubject, nil
	case OpaqueSubjectFormat:
		return Opaque, nil
	case PhoneNumberSubjectFormat:
		return PhoneNumber, nil
	case DecentralizedIdentifierSubjectFormat:
		return DecentralizedIdentifier, nil
	case UniqueResourceIdentifierSubjectFormat:
		return UniqueResourceIdentifier, nil
	case AliasesSubjectFormat:
		return Aliases, nil
	default:
		return -1, errors.New("Unable to determine subject format")
	}
}

// Converts a list of Ssf Events to a list of their corresponding Event URI's
func EventTypeArrayToEventUriArray(events []EventType) []string {
	var eventUriArr []string
	for i := 0; i < len(events); i++ {
		eventUriArr = append(eventUriArr, EventUri[events[i]])
	}
	return eventUriArr
}