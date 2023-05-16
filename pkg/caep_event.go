package pkg

import (
	pkg "caep-receiver/pkg/events"
	"errors"
	"strconv"
)

type EventType int

const (
	SessionRevoked EventType = iota
)

// Represents the interface that all CAEP Events should implement
//
// See the SessionRevokedEvent (./events/session_revoked_event.go)
// for an example
type CaepEvent interface {
	// Returns the Event URI of the CAEP Event
	GetEventUri() string

	// Returns the subject of the CAEP Event
	GetSubject() map[string]interface{}

	// Returns the Unix timestamp of the CAEP event
	GetTimestamp() int64
}

var EventUri = map[EventType]string{
	SessionRevoked: "https://schemas.openid.net/secevent/caep/event-type/session-revoked",
}

// Takes an event subject from the JSON of a CAEP Event, and converts it into the matching struct for that event
func EventStructFromEvent(eventUri string, eventSubject interface{}, claimsJson map[string]interface{}) (CaepEvent, error) {
	switch eventUri {
	case "https://schemas.openid.net/secevent/caep/event-type/session-revoked":
		subjectAttributes, ok := eventSubject.(map[string]interface{})
		timestamp, err := strconv.ParseInt(subjectAttributes["timestamp"].(string), 10, 64)
		if !ok || err != nil {
			return nil, errors.New("Unable to parse event subject")
		}

		event := pkg.SessionRevokedEvent{
			Json:           claimsJson,
			Subject:        subjectAttributes["subject"].(map[string]interface{}),
			EventTimestamp: timestamp,
		}
		return &event, nil
	}
	// Add more caep events as desired
	return nil, errors.New("No matching events")
}

// Converts a list of Caep Events to a list of their corresponding Event URI's
func EventTypeArrayToEventUriArray(events []EventType) []string {
	var eventUriArr []string
	for i := 0; i < len(events); i++ {
		eventUriArr = append(eventUriArr, EventUri[events[i]])
	}
	return eventUriArr
}
