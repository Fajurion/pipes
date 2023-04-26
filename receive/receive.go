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
		log.Printf("[%s] %s: %s: %s", protocol, message.Channel.Channel, message.Event.Sender, message.Event.Name)
	}

	switch message.Channel.Channel {
	case "broadcast":
		receiveBroadcast(protocol, message)

	case "conversation":
		receiveConversation(protocol, message)

	case "p2p":
		receiveP2P(protocol, message)
	}
}
