package params

import (
	"net/url"
	"reflect"
	"strings"

	"github.com/Nivl/go-rest-tools/network/http/httperr"
)

type Params struct {
	data interface{}
}

func NewParams(data interface{}) *Params {
	return &Params{
		data: data,
	}
}

func (p *Params) Parse(sources map[string]url.Values) error {
	paramList := reflect.ValueOf(p.data)
	if paramList.Kind() == reflect.Ptr {
		paramList = paramList.Elem()
	}

	return p.parseRecursive(paramList, sources)
}

func (p *Params) parseRecursive(paramList reflect.Value, sources map[string]url.Values) error {
	nbParams := paramList.NumField()
	for i := 0; i < nbParams; i++ {
		value := paramList.Field(i)
		info := paramList.Type().Field(i)
		tags := info.Tag

		// We make sure we can update the value of field
		if !value.CanSet() {
			return httperr.NewServerError("field [%s] could not be set", info.Name)
		}

		// Handle embedded struct
		if value.Kind() == reflect.Struct && info.Anonymous {
			p.parseRecursive(value, sources)
			continue
		}

		// We control the source of the param. If nothing is provided, we take from the URL
		paramLocation := strings.ToLower(tags.Get("from"))
		if paramLocation == "" {
			paramLocation = "url"
		}

		source, found := sources[paramLocation]
		if !found {
			return httperr.NewServerError("source [%s] for field [%s] does not exists", paramLocation, info.Name)
		}

		param := &Param{
			value: &value,
			info:  &info,
			tags:  &tags,
		}
		if err := param.setValue(&source); err != nil {
			return err
		}
	}

	return nil
}
