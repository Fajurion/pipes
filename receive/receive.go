package receive

import (
	"log"

	"github.com/Fajurion/pipes"
	"github.com/bytedance/sonic"
)

func ReceiveWS(bytes []byte) {

	// Unmarshal
	var message pipes.Message
	err := sonic.Unmarshal(bytes, &message)
	if err != nil {
		return
	}

	// Handle message
	HandleMessage("ws", message)
}

func ReceiveUDP(bytes []byte) {

	// Check for adoption request
	if bytes[0] == 'a' {

		// Adopt node
		AdoptUDP(bytes)
		return
	}

	// Unmarshal
	var message pipes.Message
	err := sonic.Unmarshal(bytes, &message)
	if err != nil {
		return
	}

	// Handle message
	HandleMessage("udp", message)
}

func HandleMessage(protocol string, message pipes.Message) {

	if pipes.DebugLogs {
		log.Printf("received on [%s] %s: %s: %s to %s", protocol, message.Channel.Channel, message.Event.Sender, message.Event.Name, message.Channel.Target)
	}

	switch message.Channel.Channel {
	case pipes.ChannelBroadcast:
		receiveBroadcast(protocol, message)

	case pipes.ChannelConversation:
		receiveConversation(protocol, message)

	case pipes.ChannelP2P:
		receiveP2P(protocol, message)
	}
}
