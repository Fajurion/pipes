package send

import (
	"context"
	"net"

	"github.com/Fajurion/pipes"
	"github.com/Fajurion/pipes/connection"
	"nhooyr.io/websocket"
)

func sendBroadcast(protocol string, message pipes.Message, msg []byte) error {

	// Send to other nodes
	switch protocol {
	case "ws":
		connection.IterateWS(func(_ string, node *websocket.Conn) bool {
			node.Write(context.Background(), websocket.MessageText, msg)
			return true
		})

	case "udp":
		connection.IterateUDP(func(_ string, node *net.UDPConn) bool {
			node.Write(msg)
			return true
		})
	}

	return nil
}
