package ssf_events

// The session revoked event is a CAEP Event, defined here:
// https://openid.net/specs/openid-caep-specification-1_0-ID1.html#rfc.section.3.1
type VerificationEvent struct {
	// Json defines the raw JSON of the CAEP Event. Used if
	// a developer wants greater control over all the attributes
	// of the CAEP Event
	Json map[string]interface{}

	// Event Receivers MAY use the value of this parameter to
	// correlate a verification event with a verification request.
	State string
}

func (event *VerificationEvent) GetEventUri() string {
	return "https://schemas.openid.net/secevent/caep/event-type/verification-event"
}

func (event *VerificationEvent) GetSubjectFormat() SubjectFormat {
	return -1
}

func (event *VerificationEvent) GetSubject() map[string]interface{} {
	return map[string]interface{}{}
}

func (event *VerificationEvent) GetTimestamp() int64 {
	return 0
}

func (event *VerificationEvent) GetType() EventType {
	return VerificationEventType
}

func (event *VerificationEvent) GetState() string {
	return event.State
}
