package adapter

import (
	"log"

	"github.com/Fajurion/pipes"
	"github.com/cornelk/hashmap"
)

type Adapter struct {
	ID string // Identifier of the client

	// Functions
	Receive func(pipes.Event, []byte) error
}

var websocketAdapters = hashmap.New[string, Adapter]()

// Register a new adapter for websocket
func AdaptWeb(adapter Adapter) {
	websocketAdapters.Insert(adapter.ID, adapter)
}

// Handles receiving messages from the target and passes them to the adapter
func ReceiveWeb(ID string, event pipes.Event, msg []byte) {

	adapter, ok := websocketAdapters.Get(ID)
	if !ok {
		return
	}

	err := adapter.Receive(event, msg)

	if err != nil {
		log.Printf("Error receiving message from target %s: %s \n", ID, err)
	}
}
