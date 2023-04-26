package pipes

var DebugLogs = false

const ChannelP2P = "p"
const ChannelConversation = "c"
const ChannelBroadcast = "br"

type Event struct {
	Sender string                 `json:"sender"` // Sender identifier (0 for system)
	Name   string                 `json:"name"`
	Data   map[string]interface{} `json:"data"`
}

type Channel struct {
	Channel string   `json:"channel"` // "p2p", "conversation", "broadcast"
	Target  []string `json:"target"`  // User IDs to send to (node and user ID for p2p channel)
	Nodes   []string `json:"-"`       // Nodes to send to (only for conversation channel)
}

type Message struct {
	Channel Channel `json:"channel"`
	Event   Event   `json:"event"`
}

func (c Channel) IsP2P() bool {
	return c.Channel == "p"
}

func (c Channel) IsConversation() bool {
	return c.Channel == "c"
}

func (c Channel) IsBroadcast() bool {
	return c.Channel == "br"
}

func P2PChannel(receiver string, receiverNode string) Channel {
	return Channel{
		Channel: "p",
		Target:  []string{receiver, receiverNode},
	}
}

func Conversation(receivers []string, nodes []string) Channel {
	return Channel{
		Channel: "c",
		Target:  receivers,
		Nodes:   nodes,
	}
}

func BroadcastChannel(receivers []string) Channel {
	return Channel{
		Channel: "br",
		Target:  receivers,
	}
}
