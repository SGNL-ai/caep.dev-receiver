package main

import (
	"fmt"
	"time"

	"caep.dev-receiver/pkg"
	events "caep.dev-receiver/pkg/ssf_events"
)

func main() {
	// Configure the receiver (do not specify poll callback if polling is not required yet)
	receiverConfig := pkg.ReceiverConfig{
		TransmitterUrl:     "https://ssf.stg.caep.dev",
		TransmitterPollUrl: "https://ssf.stg.caep.dev/ssf/streams/poll",
		EventsRequested:    []events.EventType{0},
		AuthorizationToken: "f843a2ce-4e94-48d4-aed6-c1617024b245",
		PollCallback:       nil,
	}
	// Initialize the receiver but does not start polling
	receiver, err := pkg.ConfigureSsfReceiver(receiverConfig)
	if err != nil {
		print(err)
	}

	time.Sleep(time.Duration(60) * time.Second)
	// Manually call the PollEvents() to poll events at desired time
	events, err := receiver.PollEvents()
	if err != nil {
		print(err)
	}

	PrintEvents(events)
	// Delete the receiver after done
	receiver.DeleteReceiver()

}

func PrintEvents(events []events.SsfEvent) {
	fmt.Printf("Number of events: %d\n", len(events))
	for _, event := range events {
		fmt.Println("--------EVENT-------")
		fmt.Printf("Subject Format: %v\n", event.GetSubjectFormat())
		fmt.Printf("Subject: %v\n", event.GetSubject())
		fmt.Printf("Timestamp: %d\n", event.GetTimestamp())
		fmt.Println("--------------------")
	}
	fmt.Print("\n\n")
}
