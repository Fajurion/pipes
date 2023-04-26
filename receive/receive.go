package receive

import (
	"log"

	"github.com/Fajurion/pipes"
)

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
