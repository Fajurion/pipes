package adapter

import (
	"log"

	"github.com/Fajurion/pipes"
	"github.com/cornelk/hashmap"
)

type Adapter struct {
	ID      string // Identifier of the client
	Address string // IP and port of the client (UDP only)

	// Functions
	Receive func(Context) error
}

type Context struct {
	Event   *pipes.Event
	Message []byte
	Adapter *Adapter
}

var websocketAdapters = hashmap.New[string, Adapter]()
var udpAdapters = hashmap.New[string, Adapter]()

// Register a new adapter for websocket/sl (all safe protocols)
func AdaptWS(adapter Adapter) {

	if websocketAdapters.Del(adapter.ID) {
		log.Printf("[ws] Replacing adapter for target %s \n", adapter.ID)
	}

	websocketAdapters.Insert(adapter.ID, adapter)
}

// Register a new adapter for UDP
func AdaptUDP(adapter Adapter) {

	if udpAdapters.Del(adapter.ID) {
		log.Printf("[udp] Replacing adapter for target %s \n", adapter.ID)
	}

	udpAdapters.Insert(adapter.ID, adapter)
}

// Remove a websocket/sl adapter
func RemoveWS(ID string) {
	websocketAdapters.Del(ID)
}

// Remove a UDP adapter
func RemoveUDP(ID string) {
	udpAdapters.Del(ID)
}

// Handles receiving messages from the target and passes them to the adapter
func ReceiveWeb(ID string, event pipes.Event, msg []byte) {

	adapter, ok := websocketAdapters.Get(ID)
	if !ok {
		return
	}

	err := adapter.Receive(Context{
		Event:   &event,
		Message: msg,
		Adapter: &adapter,
	})

	if err != nil {
		log.Printf("[ws] Error receiving message from target %s: %s \n", ID, err)
	}
}

// Handles receiving messages from the target and passes them to the adapter
func ReceiveUDP(ID string, event pipes.Event, msg []byte) {

	adapter, ok := udpAdapters.Get(ID)
	if !ok {
		return
	}

	err := adapter.Receive(Context{
		Event:   &event,
		Message: msg,
		Adapter: &adapter,
	})

	if err != nil {
		log.Printf("[udp] Error receiving message from target %s: %s \n", ID, err)
	}
}
