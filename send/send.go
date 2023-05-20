package send

import (
	"github.com/Fajurion/pipes"
	"github.com/Fajurion/pipes/adapter"
	"github.com/Fajurion/pipes/receive"
	"github.com/Fajurion/pipes/receive/processors"
	"github.com/bytedance/sonic"
)

const ProtocolWS = "ws"
const ProtocolUDP = "udp"

func Pipe(protocol string, message pipes.Message) error {

	// Marshal message for sending to other nodes
	msg, err := sonic.Marshal(message)
	if err != nil {
		return err
	}

	// Marshal event for sender
	event := processors.ProcessMarshal(&message, message.Event.Sender)

	// Send to sender
	switch protocol {
	case "ws":
		adapter.ReceiveWeb(message.Event.Sender, message.Event, event)

	case "udp":
		adapter.ReceiveUDP(message.Event.Sender, message.Event, event)
	}

	// Send to receivers on current node
	receive.HandleMessage(protocol, message)

	switch message.Channel.Channel {
	case "conversation":
		return sendToConversation(protocol, message, msg)

	case "broadcast":
		return sendBroadcast(protocol, message, msg)

	case "p2p":
		return sendP2P(protocol, message, msg)
	}

	return nil
}
