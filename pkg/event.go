package pkg

type EventType int

const (
	SessionRevoked EventType = iota
	DeviceComplianceChange
)

type CaepEvent interface {
	GetSubject() string
}

var eventUri = map[EventType]string{
	SessionRevoked:         "https://schemas.openid.net/secevent/caep/event-type/session-revoked",
	DeviceComplianceChange: "https://schemas.openid.net/secevent/caep/event-type/device-compliance-change",
}
