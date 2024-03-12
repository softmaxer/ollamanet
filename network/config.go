package network

import (
	"github.com/softmaxer/ollamanet/server"
)

type NetworkConfig struct {
	Url   string                `json:"url"`
	Size  int                   `json:"size"`
	Nodes []server.OllamaConfig `json:"nodes"`
}
