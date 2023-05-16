package pkg

// Represents the interface for the CAEP receiver with user facing
// methods
type CaepReceiver interface {
	ConfigureCallback(callback func(events []CaepEvent), pushInterval int) error

	// Polls the configured receiver a returns a list of the available CAEP
	// Events
	PollEvents() ([]CaepEvent, error)

	// Cleans up the Receiver's resources and deletes it from the transmitter
	DeleteReceiver()
}

// The struct that contains all the necessary fields and methods for the
// CAEP Receiver's implementation
type CaepReceiverImplementation struct {
	// transmitterUrl stores the base url of the transmitter the
	// receiver will make request to
	transmitterUrl string

	// transmitterPollUrl defines the url that the receiver
	// should hit to receiver CAEP Events
	transmitterPollUrl string

	// eventsRequested contains a list of the CAEP Event URI's requested
	// by the receiver
	eventsRequested []string

	// authorizationToken defines the Auth Token used to authorize the
	// receiver with the transmitter
	authorizationToken string

	// pushCallback defines the method the receiver will call to pass
	// events into when the push interval is triggered
	pushCallback func(events []CaepEvent)

	// pushInterval defines the interval, in seconds, between every
	// CAEP Event push delivery to the callback method
	pushInterval int

	// configurationUrl defines the transmitter's configuration url
	configurationUrl string

	// streamId defines the Id of the stream that corresponds to the
	// transmitter
	streamId string

	// terminate is used to stop the push interval routine
	terminate chan bool
}

// Struct used to read a Transmitter's configuration
type TransmitterConfig struct {
	Issuer                   string   `json:"issuer"`
	JwksUri                  string   `json:"jwks_uri,omitempty"`
	DeliveryMethodsSupported []string `json:"delivery_methods_supported,omitempty"`
	ConfigurationEndpoint    string   `json:"configuration_endpoint,omitempty"`
}

// Struct used to make a Create Stream request for the receiver
type CreateStreamReq struct {
	Delivery        CaepDelivery `json:"delivery"`
	EventsRequested []string     `json:"events_requested"`
}

// Struct that defines the deliver method for the Create Stream Request
type CaepDelivery struct {
	DeliveryMethod string `json:"delivery_method"`
}

// Struct to make a request to poll CAEP Events to the
// configured transmitter
type PollTransmitterRequest struct {
	Acknowledgements  []string `json:"ack"`
	MaxEvents         int      `json:"maxEvents,omitempty"`
	ReturnImmediately bool     `json:"returnImmediately"`
}
