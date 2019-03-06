package methods

import "net/http"

type PingService struct{}

func (s *PingService) Ping(r *http.Request, _, reply *string) error {
	*reply = "pong"
	return nil
}
