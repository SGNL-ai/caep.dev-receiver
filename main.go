package main

import (
	"caep-receiver/pkg"
	"fmt"
	"time"
)

func main() {
	receiverConfig := pkg.ReceiverConfig{
		TransmitterUrl:     "https://ssf.stg.caep.dev",
		TransmitterPollUrl: "https://ssf.stg.caep.dev/ssf/streams/poll",
		EventsRequested:    []pkg.EventType{0},
		AuthorizationToken: "f843a2ce-4e94-48d4-aed6-c1617024b245",
		PushCallback:       PrintEvents,
		PushInterval:       20,
	}
	receiver, err := pkg.ConfigureReceiver(receiverConfig)
	if err != nil {
		print(err)
	}

	time.Sleep(time.Duration(90) * time.Second)
	receiver.DeleteReceiver()
}

func PrintEvents(events []pkg.CaepEvent) {
	fmt.Printf("Number of events: %d\n", len(events))
	for _, event := range events {
		fmt.Printf("Subject: %v\n", event.GetSubject())
		fmt.Printf("Timestamp: %d\n", event.GetTimestamp())
		fmt.Println("--------------------")
	}
	fmt.Print("\n\n")
}
