package send

import (
	"context"

	"github.com/Fajurion/pipes"
	"github.com/Fajurion/pipes/connection"
	"nhooyr.io/websocket"
)

func sendBroadcast(message pipes.Message, msg []byte) error {

	// Send to other nodes
	connection.IterateWS(func(_ string, node *websocket.Conn) bool {
		node.Write(context.Background(), websocket.MessageText, msg)
		return true
	})

	return nil
}
