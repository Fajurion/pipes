package receive

import (
	"github.com/Fajurion/pipes"
	"github.com/Fajurion/pipes/adapter"
	"github.com/Fajurion/pipes/receive/processors"
)

func receiveConversation(protocol string, message pipes.Message) {

	// Send to receivers
	for _, member := range message.Channel.Target {
		if member != message.Event.Sender {

			// Process the message
			msg := processors.ProcessMarshal(&message, member)
			if msg == nil {
				continue
			}

			// Send to correct adapter
			switch protocol {
			case "ws":
				adapter.ReceiveWeb(member, message.Event, msg)

			case "udp":
				adapter.ReceiveUDP(member, message.Event, msg)
			}
		}
	}
}
