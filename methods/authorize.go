package methods

import (
	"net/http"
	"github.com/labstack/gommon/random"
)

type AuthorizeArgs struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthorizeReply struct {
	Token string `json:"token"`
}

type AuthorizeService struct {
	Service
}

func (h *AuthorizeService) Do(r *http.Request, args *AuthorizeArgs, reply *AuthorizeReply) error {
	if args.Login == "test" && args.Password == "test" {
		reply.Token = random.String(32)
		return nil
	}
	return errBadLoginOrPassword
}
