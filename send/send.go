package send

import (
	"github.com/Fajurion/pipes"
	"github.com/Fajurion/pipes/adapter"
	"github.com/Fajurion/pipes/receive"
	"github.com/bytedance/sonic"
)

const ProtocolWS = "ws"
const ProtocolUDP = "udp"

func Pipe(protocol string, message pipes.Message) error {

	msg, err := sonic.Marshal(message)
	if err != nil {
		return err
	}

	// Send to sender
	switch protocol {
	case "ws":
		adapter.ReceiveWeb(message.Event.Sender, message.Event, msg)

	case "udp":
		adapter.ReceiveUDP(message.Event.Sender, message.Event, msg)
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
