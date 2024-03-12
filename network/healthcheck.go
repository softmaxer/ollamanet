package network

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/softmaxer/ollamanet/server"
)

func HealthCheck(interval *int, network *OllamaNet) {
	scheduler := gocron.NewScheduler(time.Local)
	for _, node := range network.Servers {
		_, err := scheduler.Every(*interval).Minutes().Do(func(n server.Server) {
			healthy := n.CheckHealth()
			if !healthy {
				log.Printf("%s is not healthy", n.Address())
			}
		}, node)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	scheduler.StartAsync()
}
