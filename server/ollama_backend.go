package server

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Download struct {
	Name string `json:"name"`
}

type OllamaBackend struct {
	Url     *url.URL
	Config  *OllamaConfig
	Proxy   *httputil.ReverseProxy
	isAlive bool
}

type OllamaConfig struct {
	BaseUrl     string  `json:"base_url"`
	System      string  `json:"system"`
	Model       string  `json:"model"`
	Temperature float32 `json:"temperature"`
}

func (ob *OllamaBackend) Init() int {
	body, err := json.Marshal(Download{Name: ob.Config.Model})
	if err != nil {
		log.Fatal("Error marshalling json: ", err.Error())
	}
	response, err := http.Post(ob.Url.String()+"/api/pull", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("Couldn't sent request")
	}
	defer response.Body.Close()
	log.Printf("Ollama backend %s init status: %d\n", ob.Url.String(), response.StatusCode)
	return response.StatusCode
}

func (ob *OllamaBackend) IsAlive() bool {
	return ob.isAlive
}

func (ob *OllamaBackend) SetAlive(alive bool) {
	ob.isAlive = alive
}

func (ob *OllamaBackend) Serve(rw http.ResponseWriter, req *http.Request) {
	ob.Proxy.ServeHTTP(rw, req)
}

func (ob *OllamaBackend) Address() string {
	return ob.Url.String()
}

func (ob *OllamaBackend) CheckHealth() bool {
	response, err := http.Head(ob.Url.String())
	if err != nil {
		ob.SetAlive(false)
		return false
	}
	if response.StatusCode != http.StatusOK {
		ob.SetAlive(false)
		return false
	}
	return true
}
