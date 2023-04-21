package send

import (
	"errors"

	"github.com/Fajurion/pipes"
)

func Socketless(nodeEntity pipes.Node, message pipes.Message) error {

	/*
		_, err := util.PostRaw(nodeEntity., map[string]interface{}{
			"this":    util.NODE_ID,
			"token":   nodeEntity.Token,
			"message": message,
		}) */

	return errors.New("not implemented")
}
