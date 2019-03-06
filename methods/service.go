package methods

import "github.com/gorilla/rpc/v2/json"

type Service struct {
}

func (s *Service) err(code int, message string) error {
	return &json.Error{
		Data: err{
			Code:    code,
			Message: message,
		},
	}
}
