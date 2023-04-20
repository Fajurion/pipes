package receive

import (
	"log"

	"github.com/Fajurion/pipes"
)

func Handle(message pipes.Message) {

	log.Printf("%s: %s: %s", message.Channel.Channel, message.Event.Sender, message.Event.Name)

	switch message.Channel.Channel {
	case "broadcast":
		receiveBroadcast(message)

	case "conversation":
		receiveConversation(message)

	case "p2p":
		receiveP2P(message)

	}
}
