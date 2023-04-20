package send

import (
	"context"

	"github.com/Fajurion/pipes"
	"nhooyr.io/websocket"
)

func sendBroadcast(message pipes.Message, msg []byte) error {

	// Send to other nodes
	pipes.IterateConnections(func(_ string, node *websocket.Conn) bool {
		node.Write(context.Background(), websocket.MessageText, msg)
		return true
	})

	return nil
}
