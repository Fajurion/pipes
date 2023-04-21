package send

import (
	"github.com/Fajurion/pipes"
	"github.com/Fajurion/pipes/adapter"
	"github.com/Fajurion/pipes/receive"
	"github.com/bytedance/sonic"
)

func WS(message pipes.Message) error {

	msg, err := sonic.Marshal(message)
	if err != nil {
		return err
	}

	// Send to own client(s)
	adapter.ReceiveWeb(message.Event.Sender, message.Event, msg)
	receive.HandleWS(message)

	switch message.Channel.Channel {
	case "conversation":
		return sendToConversation(message, msg)

	case "broadcast":
		return sendBroadcast(message, msg)

	case "p2p":
		return sendP2P(message, msg)
	}

	return nil
}

func UDP(message pipes.Message) error {
	// TODO
	return nil
}
