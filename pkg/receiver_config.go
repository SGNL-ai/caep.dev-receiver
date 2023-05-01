package pkg

type ReceiverConfig struct {
	// TransmitterUrl defines the URL for the transmitter that
	// the configured receiver will create a stream with and receive
	// events from.
	//
	// Required
	TransmitterUrl string

	// EventsRequested specified the CAEP events you want to receiver
	// from the transmitter.
	//
	// Required
	EventsRequested []EventNums

	// AuthorizationToken is the authorization token used to authorize
	// your receiver with the specified transmitter
	//
	// Note - all transmitter's will require an authorization token
	//
	// Required
	AuthorizationToken string

	// PushCallback is used to configure the method that you want the
	// receiver to call when it's received CAEP events. Each time the
	// receiver runs the push method, it will call the PushCallback and
	// pass as a parameter the list of received CAEP events since the
	// last push call.
	//
	// Note - The PushCallback and PushInterval can also be configured
	// after initial receiver construction
	//
	// Optional
	PushCallback func(events []CaepEvent)

	// PushInterval defines, in seconds how often you want the receiver to
	// push any CAEP events to your callback function.
	//
	// Note - This field will not be used if the PushCallback isn't configured
	//
	// Optional, defaults to __
	PushInterval int
}
