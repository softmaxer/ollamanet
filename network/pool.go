package network

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/softmaxer/ollamanet/server"
)

type OllamaNet struct {
	Url        *url.URL
	LastServed int
	Servers    []server.Server
	Turn       int
}

func (network *OllamaNet) GetNextAvailableServer() (server.Server, error) {
	nextTurn := (network.Turn + 1) % len(network.Servers)
	targetServer := network.Servers[nextTurn]
	if !targetServer.IsAlive() {
		return nil, &ServerNotAlive{}
	}
	network.LastServed = nextTurn
	network.Turn++
	return targetServer, nil
}

func (network *OllamaNet) Redirect(rw http.ResponseWriter, req *http.Request) {
	server, err := network.GetNextAvailableServer()
	if err != nil {
		log.Fatal("Couldn't get any available server: ", err.Error())
	}

	server.Serve(rw, req)
}

func Init(config *NetworkConfig) OllamaNet {
	var serverList []server.Server
	for _, node := range config.Nodes {
		var backend server.OllamaBackend
		parsedUrl, err := url.Parse(node.BaseUrl)
		if err != nil {
			log.Fatal("Error setting url: ", err.Error())
		}
		backend.Url = parsedUrl
		backend.Config = &node
		backend.Proxy = httputil.NewSingleHostReverseProxy(backend.Url)
		backend.SetAlive(true)
		initStatus := backend.Init()
		if initStatus != 200 {
			continue
		}

		serverList = append(serverList, &backend)
	}

	networkUrl, err := url.Parse(config.Url)
	if err != nil {
		log.Fatal("Error parsing URL")
	}
	newNetwork := OllamaNet{
		Url:        networkUrl,
		LastServed: 0,
		Servers:    serverList,
		Turn:       0,
	}

	return newNetwork
}
