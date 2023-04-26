package connection

import (
	"log"
	"net"

	"github.com/Fajurion/pipes"
	"github.com/bytedance/sonic"
	"github.com/cornelk/hashmap"
)

var nodeUDPConnections = hashmap.New[string, *net.UDPConn]()

/* Eventually implement custom UDP protocol
type AdoptionRequest struct {
	Token    string     `json:"tk"`
	Adopting pipes.Node `json:"adpt"`
}
*/

func ConnectUDP(node pipes.Node) {

	// Marshal current node
	adoptionRq, err := sonic.Marshal(AdoptionRequest{
		Token:    node.Token,
		Adopting: pipes.CurrentNode,
	})
	if err != nil {
		return
	}

	// Add prefix
	adoptionRq = append([]byte("a:"), adoptionRq...)

	// Resolve udp address
	udpAddr, err := net.ResolveUDPAddr("udp", node.UDP)
	if err != nil {
		return
	}

	// Connect to node
	c, err := net.DialUDP("udp", nil, udpAddr)

	if err != nil {
		c.Close()
		return
	}

	// Send adoption request
	_, err = c.Write(adoptionRq)
	if err != nil {
		c.Close()
		return
	}

	// Add connection to map
	nodeUDPConnections.Insert(node.ID, c)

	log.Printf("[udp] Outgoing event stream to node %s connected.", node.ID)

	// TODO: Connect to UDP
}

func ExistsUDP(node string) bool {

	// Check if connection exists
	_, ok := nodeUDPConnections.Get(node)
	return ok
}

func GetUDP(node string) *net.UDPConn {

	// Check if connection exists
	connection, ok := nodeUDPConnections.Get(node)
	if !ok {
		return nil
	}

	return connection
}

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
func IterateUDP(f func(key string, value *net.UDPConn) bool) {
	nodeUDPConnections.Range(f)
}