package receive

import (
	"github.com/Fajurion/pipes"
	"github.com/Fajurion/pipes/adapter"
	"github.com/Fajurion/pipes/receive/processors"
)

func receiveP2P(message pipes.Message) {

	// Process the message
	msg := processors.ProcessMarshal(&message, message.Channel.Target[0])
	if msg == nil {
		return
	}

	// Send to receiver
	adapter.ReceiveWeb(message.Channel.Target[0], message.Event, msg)
}
