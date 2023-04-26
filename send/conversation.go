package send

import (
	"context"

	"github.com/Fajurion/pipes"
	"github.com/Fajurion/pipes/connection"
	"nhooyr.io/websocket"
)

func sendToConversation(protocol string, message pipes.Message, msg []byte) error {

	for _, node := range message.Channel.Nodes {
		if node == pipes.CurrentNode.ID {
			continue
		}

		connection.GetWS(node).Write(context.Background(), websocket.MessageText, msg)
	}

	return nil
}
