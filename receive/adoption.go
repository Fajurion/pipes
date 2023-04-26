package receive

import (
	"errors"
	"log"
	"net"

	"github.com/Fajurion/pipes"
	"github.com/Fajurion/pipes/connection"
	"github.com/bytedance/sonic"
)

func ReceiveWSAdoption(request string) error {

	// Unmarshal
	var adoptionRq connection.AdoptionRequest
	err := sonic.Unmarshal([]byte(request), &adoptionRq)
	if err != nil {
		return err
	}

	// Check token
	if adoptionRq.Token != pipes.CurrentNode.Token {
		return errors.New("invalid token")
	}

	log.Printf("[ws] Incoming event stream from node %s connected.", adoptionRq.Adopting.ID)
	pipes.AddNode(adoptionRq.Adopting)

	// Connect output stream (if not already connected)
	if !connection.ExistsWS(adoptionRq.Adopting.ID) {
		connection.ConnectWS(adoptionRq.Adopting)
	}

	return nil
}

func AdoptUDP(bytes []byte, conn *net.UDPConn) {

	// Close connection by default
	close := true
	defer func() {
		if close {
			conn.Close()
		}
	}()

	// Remove adoption request prefix
	bytes = bytes[2:]

	// Unmarshal
	var adoptionRq connection.AdoptionRequest
	err := sonic.Unmarshal(bytes, &adoptionRq)
	if err != nil {
		return
	}

	// Check token
	if adoptionRq.Token != pipes.CurrentNode.Token {
		return
	}

	// Stop closing connection
	close = false

	log.Printf("[udp] Incoming event stream from node %s connected.", adoptionRq.Adopting.ID)
	pipes.AddNode(adoptionRq.Adopting)

	// Connect output stream (if not already connected)
	if !connection.ExistsUDP(adoptionRq.Adopting.ID) {
		connection.ConnectUDP(adoptionRq.Adopting)
	}
}
