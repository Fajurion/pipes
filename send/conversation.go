package send

import (
	"context"
	"errors"

	"github.com/Fajurion/pipes"
	"github.com/Fajurion/pipes/connection"
	"nhooyr.io/websocket"
)

func sendToConversation(protocol string, message pipes.Message, msg []byte) error {

	if protocol == "udp" {
		return errors.New("udp not supported for conversation")
	}

	for _, node := range message.Channel.Nodes {
		if node == pipes.CurrentNode.ID {
			continue
		}

		// Encrypt message for node
		encryptedMsg, err := pipes.Encrypt(node, msg)
		if err != nil {
			return err
		}

		connection.GetWS(node).Write(context.Background(), websocket.MessageText, encryptedMsg)
	}

	return nil
}
