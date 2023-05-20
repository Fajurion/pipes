package send

import (
	"log"

	"github.com/Fajurion/pipes"
	"github.com/Fajurion/pipes/receive"
	"github.com/Fajurion/pipes/util"
)

func Socketless(nodeEntity pipes.Node, message pipes.Message) error {

	if pipes.DebugLogs {
		log.Printf("sent on [socketless] %s: %s: %s", message.Channel.Channel, message.Event.Sender, message.Event.Name)
	}

	if nodeEntity.ID == pipes.CurrentNode.ID {

		receive.HandleMessage("ws", message)
		return nil
	}

	err := util.PostRaw(nodeEntity.SL, map[string]interface{}{
		"this":    pipes.CurrentNode.ID,
		"token":   nodeEntity.Token,
		"message": message,
	})

	if err != nil {
		return err
	}

	return nil
}
