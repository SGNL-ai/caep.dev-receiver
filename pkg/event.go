package pkg

type EventType int

const (
	SessionRevoked EventType = iota
	DeviceComplianceChange
)

type CaepEvent interface {
	// Placeholder for now - not sure what common methods
	// we'd want for this interface?
	String() string
}
