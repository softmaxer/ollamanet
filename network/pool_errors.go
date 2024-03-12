package network

import (
	"fmt"
	"net/url"
)

type ServerNotAlive struct {
	url *url.URL
}

func (e *ServerNotAlive) Error() string {
	return fmt.Sprintf("Server %s not alive!", e.url)
}
