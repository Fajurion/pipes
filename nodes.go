package pipes

import (
	"log"

	"github.com/cornelk/hashmap"
)

type Node struct {
	ID    string `json:"id"`
	Token string `json:"token"`
	Host  string `json:"host"`
}

var nodes = hashmap.New[string, Node]()

var CurrentNode Node

func SetupCurrent(id string, token string) {

	// Set log prefix
	log.SetPrefix("[pipes] ")

	CurrentNode = Node{
		ID:    id,
		Token: token,
		Host:  "",
	}
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
