package send

import (
	"github.com/Fajurion/pipes"
	"github.com/Fajurion/pipes/util"
)

func Socketless(nodeEntity pipes.Node, message pipes.Message) error {

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
