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

	// Encrypt message for node
	encryptedMsg, err := pipes.Encrypt(message.Channel.Target[1], msg)
	if err != nil {
		return err
	}

	// Send to correct node
	switch protocol {
	case "ws":
		connection.GetWS(message.Channel.Target[1]).Write(context.Background(), websocket.MessageText, encryptedMsg)

	case "udp":
		connection.GetUDP(message.Channel.Target[1]).Write(encryptedMsg)
	}

	return nil
}
