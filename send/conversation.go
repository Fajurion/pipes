package send

import (
	"context"

	"github.com/Fajurion/pipes"
	"nhooyr.io/websocket"
)

func sendToConversation(message pipes.Message, msg []byte) error {

	for _, node := range message.Channel.Nodes {
		if node == pipes.CurrentNode.ID {
			continue
		}

		pipes.GetConnection(node).Write(context.Background(), websocket.MessageText, msg)
	}

	return nil
}
