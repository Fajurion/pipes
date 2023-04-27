package pipes

import (
	"github.com/cornelk/hashmap"
)

type Node struct {
	ID    string `json:"id"`
	Token string `json:"token"`
	WS    string `json:"ws,omitempty"`  // Websocket ip
	UDP   string `json:"udp,omitempty"` // UDP ip
	SL    string `json:"sl,omitempty"`  // Socketless pipe
}

var nodes = hashmap.New[string, Node]()

var CurrentNode Node

func SetupCurrent(id string, token string) {

	CurrentNode = Node{
		ID:    id,
		Token: token,
		WS:    "",
		UDP:   "",
	}
}

func SetupWS(ws string) {
	CurrentNode.WS = ws
}

func SetupUDP(udp string) {
	CurrentNode.UDP = udp
}

func SetupSocketless(sl string) {
	CurrentNode.SL = sl
}

func GetNode(id string) *Node {

	// Get node
	node, ok := nodes.Get(id)
	if !ok {
		return nil
	}

	return &node
}

func AddNode(node Node) {
	nodes.Insert(node.ID, node)
}

// IterateConnections iterates over all connections. If the callback returns false, the iteration stops.
func IterateNodes(callback func(string, Node) bool) {
	nodes.Range(callback)
}
