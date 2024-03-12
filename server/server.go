package server

import "net/http"

type Server interface {
	IsAlive() bool
	SetAlive(alive bool)
	Serve(http.ResponseWriter, *http.Request)
}
