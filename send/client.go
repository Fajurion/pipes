package send

import (
	"github.com/Fajurion/pipes"
	"github.com/Fajurion/pipes/adapter"
	"github.com/bytedance/sonic"
)

func Client(id string, event pipes.Event) {

	msg, err := sonic.Marshal(event)
	if err != nil {
		return
	}

	adapter.ReceiveWeb(id, event, msg)
}
