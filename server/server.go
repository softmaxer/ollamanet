package server

import "net/http"

type Server interface {
	Address() string
	IsAlive() bool
	SetAlive(alive bool)
	Serve(http.ResponseWriter, *http.Request)
	CheckHealth() bool
}
