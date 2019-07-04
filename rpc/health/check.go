package health

import "net/http"

type CheckReply struct {
	Status string
}

func (s *Service) Check(r *http.Request, _, reply *CheckReply) error {
	reply.Status = "up"
	return nil
}
