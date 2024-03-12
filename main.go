package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/softmaxer/ollamanet/network"
)

func main() {
	var configFile string
	var logFile string
	flag.StringVar(&configFile, "-config", "configuration.json", "A configuration file for the Ollama network")
	flag.StringVar(&logFile, "-log", "network_log", "File to use for logging Ollama network")
	jsonConfig, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal("Error reading configuration file: ", err.Error())
	}
	logFileDescriptor, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer logFileDescriptor.Close()
	log.SetOutput(logFileDescriptor)
	var netConfig network.NetworkConfig
	err = json.Unmarshal(jsonConfig, &netConfig)
	if err != nil {
		log.Fatal("Error loading ollama network config: ", err.Error())
	}
	net := network.Init(&netConfig)
	http.HandleFunc("/", net.Redirect)
	http.ListenAndServe(":8000", nil)
}
