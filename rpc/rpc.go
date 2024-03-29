package rpc

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/rpc/v2"
	jsonCodec "github.com/gorilla/rpc/v2/json"
	"log"
	"os"
	"reflect"
)

func NewRPCServer(services map[string]interface{}) *rpc.Server {

	handler := rpc.NewServer()

	handler.RegisterCodec(jsonCodec.NewCodec(), "application/json")

	for name, service := range services {
		handler.RegisterService(service, name)
	}

	debugMethods(handler)

	return handler
}

func debugMethods(rpc *rpc.Server) {
	services := reflect.Indirect(reflect.ValueOf(*rpc).FieldByName("services")).FieldByName("services")
	var data []string
	for _, serviceName := range services.MapKeys() {
		service := reflect.Indirect(services.MapIndex(serviceName))
		methodsMap := service.FieldByName("methods")
		for _, methodName := range methodsMap.MapKeys() {
			data = append(data, fmt.Sprintf("%v.%v", serviceName, methodName))
		}
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
	fmt.Fprintf(os.Stdout, "\ndebug RPC methods:\n%s\n", jsonData)
}
