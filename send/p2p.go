package send

import (
	"context"

	"github.com/Fajurion/pipes"
	"github.com/Fajurion/pipes/adapter"
	"github.com/Fajurion/pipes/connection"
	"nhooyr.io/websocket"
)

func sendP2P(protocol string, message pipes.Message, msg []byte) error {

	// Check if receiver is on this node
	if message.Channel.Target[0] == pipes.CurrentNode.ID {
		adapter.ReceiveWeb(message.Channel.Target[1], message.Event, msg)
		return nil
	}

	// Send to correct node
	switch protocol {
	case "ws":
		connection.GetWS(message.Channel.Target[1]).Write(context.Background(), websocket.MessageText, msg)

	case "udp":
		connection.GetUDP(message.Channel.Target[1]).Write(msg)
	}

	return nil
}
