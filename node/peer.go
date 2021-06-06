package node

type PeerNode struct {
	Host        string `json:"host"`
	IsBootstrap bool   `json:"isBootstrap"`

	connected bool
}

func NewPeerNode(host string, isBootstrap bool, connected bool) PeerNode {
	return PeerNode{
		Host:        host,
		IsBootstrap: isBootstrap,
		connected:   connected,
	}
}
