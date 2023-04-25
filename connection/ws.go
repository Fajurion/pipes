package connection

import (
	"context"
	"log"

	"github.com/Fajurion/pipes"
	"github.com/bytedance/sonic"
	"github.com/cornelk/hashmap"
	"nhooyr.io/websocket"
)

var nodeWSConnections = hashmap.New[string, *websocket.Conn]()

type AdoptionRequest struct {
	Token    string     `json:"tk"`
	Adopting pipes.Node `json:"adpt"`
}

func ConnectWS(node pipes.Node) {

	// Marshal current node
	nodeBytes, err := sonic.Marshal(AdoptionRequest{
		Token:    node.Token,
		Adopting: pipes.CurrentNode,
	})
	if err != nil {
		return
	}

	// Connect to node
	c, _, err := websocket.Dial(context.Background(), node.WS, &websocket.DialOptions{
		Subprotocols: []string{string(nodeBytes)},
	})

	if err != nil {
		return
	}

	// Add connection to map
	nodeWSConnections.Insert(node.ID, c)

	log.Printf("[ws] Outgoing event stream to node %s connected.", node.ID)

	// TODO: Connect to UDP
}

func ExistsWS(node string) bool {

	// Check if connection exists
	_, ok := nodeWSConnections.Get(node)
	return ok
}

func GetWS(node string) *websocket.Conn {

	// Check if connection exists
	connection, ok := nodeWSConnections.Get(node)
	if !ok {
		return nil
	}

	return connection
}

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
func IterateWS(f func(key string, value *websocket.Conn) bool) {
	nodeWSConnections.Range(f)
}
