package errors

import (
	"fmt"
	"github.com/gorilla/rpc/v2/json"
	"gopkg.in/go-playground/validator.v9"
)

type ErrorCode int

const (
	CodeValidationFailed   ErrorCode = 1000
	CodeBadLoginOrPassword ErrorCode = 1001
)

var (
	ErrBadLoginOrPassword = New(CodeBadLoginOrPassword, "bad login or password")
)

type PlainErr struct {
	Code    ErrorCode                `json:"code"`
	Message string                   `json:"message"`
	Fields  []map[string]interface{} `json:"fields"`
}

func (e PlainErr) Error() string {
	return e.Message
}

func New(code ErrorCode, message string) error {
	return &json.Error{
		Data: &PlainErr{
			Code:    code,
			Message: message,
		},
	}
}

//type ValidationErr struct {
//	Code    ErrorCode                `json:"code"`
//	Message string                   `json:"message"`
//	Errors  []map[string]interface{} `json:"errors"`
//}

//func (e ValidationErr) Error() string {
//	return json.Error{
//		Data: 1,
//	}
//
//}

func NewValidationErr(errors validator.ValidationErrors) error {
	var messages []map[string]interface{}
	for _, err := range errors {
		messages = append(messages, map[string]interface{}{
			"field":     err.Field(),
			"validator": err.Tag(),
			"message":   "",
		})
		fmt.Println("err.Namespace()", err.Namespace())
		fmt.Println("err.Field()", err.Field())
		fmt.Println("err.StructNamespace()", err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
		fmt.Println("err.StructField()", err.StructField())         // by passing alt name to ReportError like below
		fmt.Println("err.Tag()", err.Tag())
		fmt.Println("err.ActualTag()", err.ActualTag())
		fmt.Println("err.Kind()", err.Kind())
		fmt.Println("err.Type()", err.Type())
		fmt.Println("err.Value()", err.Value())
		fmt.Println("err.Param()", err.Param())
		fmt.Println("")

	}
	return &json.Error{
		Data: &PlainErr{
			Code:    CodeValidationFailed,
			Message: "validation failed",
			Fields:  messages,
		},
	}
}
