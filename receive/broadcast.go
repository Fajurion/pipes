package receive

import (
	"log"

	"github.com/Fajurion/pipes"
	"github.com/Fajurion/pipes/adapter"
	"github.com/Fajurion/pipes/receive/processors"
)

func receiveBroadcast(protocol string, message pipes.Message) {

	if message.Event.Name == "ping" {
		log.Println("Received ping from node", message.Event.Data["node"])
	}

	// Send to all receivers
	for _, tg := range message.Channel.Target {

		// Process the message
		msg := processors.ProcessMarshal(&message, tg)
		if msg == nil {
			continue
		}

		// Send to correct adapter
		switch protocol {
		case "ws":
			adapter.ReceiveWeb(tg, message.Event, msg)

		case "udp":
			adapter.ReceiveUDP(tg, message.Event, msg)
		}
	}
}
