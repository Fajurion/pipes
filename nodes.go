package pipes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"log"

	"github.com/Fajurion/pipes/util"
	"github.com/cornelk/hashmap"
)

type Node struct {
	ID    string `json:"id"`
	Token string `json:"token"`
	WS    string `json:"ws,omitempty"`  // Websocket ip
	UDP   string `json:"udp,omitempty"` // UDP ip
	SL    string `json:"sl,omitempty"`  // Socketless pipe

	// Encryption
	Cipher cipher.Block
}

var nodes = hashmap.New[string, Node]()

var CurrentNode Node

func SetupCurrent(id string, token string) {

	if len(token) < 32 {
		panic("Token is too short (must be longer than 32 characters for AES-256 encryption)")
	}

	// Create encryption cipher
	tokenHash := sha256.Sum256([]byte(token))
	encryptionKey := tokenHash[:]

	log.Println("Encryption key:", base64.StdEncoding.EncodeToString(encryptionKey))

	cipher, err := aes.NewCipher(encryptionKey)
	if err != nil {
		panic(err)
	}

	CurrentNode = Node{
		ID:     id,
		Token:  token,
		WS:     "",
		UDP:    "",
		SL:     "",
		Cipher: cipher,
	}
}

func Encrypt(node string, msg []byte) ([]byte, error) {

	// Get key and encrypt message
	key := GetNode(node).Cipher
	encrypted, err := util.EncryptAES(key, msg)
	if err != nil {
		return nil, err
	}

	return encrypted, nil
}

func Decrypt(node string, msg []byte) ([]byte, error) {

	// Get key and decrypt message
	key := GetNode(node).Cipher
	decrypted, err := util.DecryptAES(key, msg)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
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
