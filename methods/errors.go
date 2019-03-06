package methods

import "github.com/gorilla/rpc/v2/json"

var (
	errBadLoginOrPassword = &json.Error{Data: err{Code: 404, Message: "bad login or password"}}
)

type err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
