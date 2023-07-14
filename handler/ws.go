package handler

import (
	"chat-node/bridge"
	"log"
	"time"

	"github.com/Fajurion/pipes"
	"github.com/Fajurion/pipes/send"
	"github.com/bytedance/sonic"
)

type Message struct {
	Client *bridge.Client         `json:"client"`
	Action string                 `json:"action"` // The action to perform
	Data   map[string]interface{} `json:"data"`
}

// Routes is a map of all the routes
var Routes map[string]func(Message)

func HandleWS(msg []byte) bool {

	var message Message
	err := sonic.Unmarshal(msg, &message)
	if err != nil {
		return false
	}

	defer func() {
		if err := recover(); err != nil {
			ErrorResponse(message, "internal")
		}
	}()

	// Check if the action exists
	if Routes[message.Action] == nil {
		return false
	}

	log.Println("Handling message: " + message.Action)

	go RouteWS(message.Action, message)

	return true
}

func RouteWS(action string, message Message) {
	defer func() {
		if err := recover(); err != nil {
			ErrorResponse(message, "invalid")
		}
	}()

	Routes[message.Action](message)
}

func InitializeWS() {
	Routes = make(map[string]func(Message))
}

func TestConnectionWS() {
	go func() {
		for {
			time.Sleep(time.Second * 5)

			// Send ping
			send.Pipe(send.ProtocolWS, pipes.Message{
				Channel: pipes.BroadcastChannel([]string{"0"}),
				Event: pipes.Event{
					Sender: "0",
					Name:   "ping",
					Data: map[string]interface{}{
						"node": pipes.CurrentNode.ID,
					},
				},
			})
		}
	}()
}
