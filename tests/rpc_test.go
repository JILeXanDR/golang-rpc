package tests

import (
	"bytes"
	"encoding/json"
	"github.com/JILeXanDR/golang-rpc/rpc"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRPCServer(t *testing.T) {

	handler := rpc.NewRPCServer()

	t.Run("should return text error when request method is not POST", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/rpc", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		assert.Equal(t, 405, rr.Code)
		assert.Equal(t, "rpc: POST method required, received GET", rr.Body.String())
	})

	t.Run("should return when request method is POST", func(t *testing.T) {
		rr, _ := rpcCall("Health.Check", map[string]interface{}{}, 1)
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, `{"result":{"Status":"up"},"error":null,"id":1}`, rr.Body.String())
	})

	t.Run("should return when request method is POST", func(t *testing.T) {
		rr, _ := rpcCall("Authorize.Login", map[string]interface{}{"login": "test", "password": "test"}, 1)
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, `{"result":{"Token":"up"},"error":null,"id":1}`, rr.Body.String())
	})
}

func rpcCall(method string, params interface{}, id interface{}) (*httptest.ResponseRecorder, error) {
	b, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  []interface{}{params},
		"id":      id,
	})
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(b)

	req, err := http.NewRequest("POST", "/rpc", reader)
	if err != nil {
		return nil, err
	}

	rr := httptest.NewRecorder()

	rpc.NewRPCServer().ServeHTTP(rr, req)

	return rr, nil
}
