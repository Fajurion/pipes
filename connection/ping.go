package connection

import (
	"strings"
	"time"

	"github.com/Fajurion/pipes"
	"github.com/Fajurion/pipes/util"
)

var DisconnectHandler = func(node pipes.Node) {}

func SetupDisconnections() {

	// Start ping routine
	go func() {

		for {

			// Ping all nodes every second
			time.Sleep(1 * time.Second)

			pipes.IterateNodes(func(_ string, node pipes.Node) bool {

				err := util.PostRaw(strings.Replace(node.SL, "socketless", "ping", -1), map[string]interface{}{})
				if err != nil {
					DisconnectHandler(node)
					RemoveUDP(node.ID)
					RemoveWS(node.ID)
				}

				return true
			})

		}

	}()

}
