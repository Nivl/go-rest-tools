package router

import (
	"net/url"
	"reflect"

	"github.com/Nivl/go-rest-tools/router/params"
)

type Guard struct {
	// ParamStruct is an instance of a struct that describes the http params
	// accepted by an endpoint
	ParamStruct interface{}
}

// ParseParams parses and returns the list of params needed
// Returns an error if a required param is missing, or if a type is wrong
func (g *Guard) ParseParams(sources map[string]url.Values) (interface{}, error) {
	// We give p the same type as g.ParamStruct
	p := reflect.New(reflect.TypeOf(g.ParamStruct).Elem()).Interface()
	err := params.NewParams(p).Parse(sources)
	return p, err
}

// func NewGuard(paramsType interface{}, req HTTPRequest, sources map[string]url.Values) *Guard {
// 	return &Guard{}
// }
