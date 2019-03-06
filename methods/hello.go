package methods

import (
	"net/http"
)

type HelloService struct{}

func (s *HelloService) Say(r *http.Request, args *string, reply *string) error {
	*reply = "Hello, " + *args + "!"
	return nil
}
