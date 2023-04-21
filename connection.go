package pipes

import (
	"context"
	"log"

	"github.com/bytedance/sonic"
	"github.com/cornelk/hashmap"
	"nhooyr.io/websocket"
)

var nodeConnections = hashmap.New[string, *websocket.Conn]()

type AdoptionRequest struct {
	Token    string `json:"tk"`
	Adopting Node   `json:"adpt"`
}

func ConnectToNode(node Node) {

	// Marshal current node
	nodeBytes, err := sonic.Marshal(AdoptionRequest{
		Token:    node.Token,
		Adopting: CurrentNode,
	})
	if err != nil {
		return
	}

	// Connect to node
	c, _, err := websocket.Dial(context.Background(), node.Host, &websocket.DialOptions{
		Subprotocols: []string{string(nodeBytes)},
	})

	if err != nil {
		return
	}

	// Add connection to map
	nodeConnections.Insert(node.ID, c)

	log.Printf("Outgoing event stream to node %s connected.", node.ID)

	// TODO: Connect to UDP
}

func ConnectionExists(node string) bool {

	// Check if connection exists
	_, ok := nodeConnections.Get(node)
	return ok
}

func GetConnection(node string) *websocket.Conn {

	// Check if connection exists
	connection, ok := nodeConnections.Get(node)
	if !ok {
		return nil
	}

	return connection
}

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
func IterateConnections(f func(key string, value *websocket.Conn) bool) {
	nodeConnections.Range(f)
}
