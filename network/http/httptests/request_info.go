package httptests

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/url"
	"strings"

	"github.com/Nivl/go-params"
	"github.com/Nivl/go-params/formfile"
	"github.com/Nivl/go-rest-tools/router"
	"github.com/gorilla/mux"
)

// RequestInfo represents the params accepted by NewRequest
type RequestInfo struct {
	Endpoint *router.Endpoint

	Params interface{}  // Optional
	Auth   *RequestAuth // Optional
	// Router is used to parse Mux Variables. Default on the api router
	Router *mux.Router

	urlParams   map[string]string
	bodyParams  map[string]interface{} // can be string or []string
	queryParams map[string]interface{} // can be string or []string
	fileParams  map[string]*formfile.FormFile
}

// ParseParams parses the params and copy them in the right list:
// urlParams, bodyParams, and queryParams
func (ri *RequestInfo) ParseParams() {
	ri.urlParams = map[string]string{}
	ri.bodyParams = map[string]interface{}{}
	ri.queryParams = map[string]interface{}{}

	if ri.Params == nil {
		return
	}

	var sources map[string]url.Values
	p := params.New(ri.Params)
	sources, ri.fileParams = p.Extract()

	for k, v := range sources["url"] {
		ri.urlParams[k] = v[0]
	}
	for k, v := range sources["form"] {
		if len(v) == 1 {
			ri.bodyParams[k] = v[0]
		} else {
			ri.bodyParams[k] = v
		}
	}
	for k, v := range sources["query"] {
		if len(v) == 1 {
			ri.queryParams[k] = v[0]
		} else {
			ri.queryParams[k] = v
		}
	}
}

// URL returns the full URL
func (ri *RequestInfo) URL() string {
	url := ri.Endpoint.Path
	for param, value := range ri.urlParams {
		url = strings.Replace(url, "{"+param+"}", value, -1)
	}
	return url
}

// PopulateQuery populate the query string of a request
func (ri *RequestInfo) PopulateQuery(qs url.Values) {
	for key, value := range ri.queryParams {
		if slice, isSlice := value.([]string); isSlice {
			qs[key] = slice
		} else {
			qs.Add(key, value.(string))
		}
	}
}

// Body returns the full Body of the request
func (ri *RequestInfo) Body() (mime string, body io.Reader, err error) {
	if ri.fileParams != nil && len(ri.fileParams) > 0 {
		return ri.BodyMultipart()
	}
	return ri.BodyJSON()
}

// BodyJSON returns the body of the request encoded in JSON
// FIXME(melvin): because the data are from a map of string, all the
// 		JSON data will also be string. There's no way to use the output
// 		to recreate a new JSON object containing non string value
func (ri *RequestInfo) BodyJSON() (mime string, body io.Reader, err error) {
	mime = "application/json; charset=utf-8"
	body = bytes.NewBufferString("")

	// Parse the body as a JSON object
	if len(ri.bodyParams) > 0 {
		var jsonDump []byte
		jsonDump, err = json.Marshal(ri.bodyParams)
		if err != nil {
			return
		}
		body = bytes.NewBuffer(jsonDump)
	}
	return
}

// BodyMultipart returns the body of the request as multipart data
func (ri *RequestInfo) BodyMultipart() (mime string, body io.Reader, err error) {
	output := &bytes.Buffer{}
	writer := multipart.NewWriter(output)

	// We attach the files
	for name, f := range ri.fileParams {
		part, err := writer.CreateFormFile(name, f.Header.Filename)
		if err != nil {
			return "", nil, err
		}

		_, err = io.Copy(part, f.File)
		if err != nil {
			return "", nil, err
		}
	}

	// We attach any other form params
	for name, value := range ri.bodyParams {
		if slice, isSlice := value.([]string); isSlice {
			for _, elem := range slice {
				if err = writer.WriteField(name, elem); err != nil {
					return "", nil, err
				}
			}
		} else {
			if err = writer.WriteField(name, value.(string)); err != nil {
				return "", nil, err
			}
		}
	}

	if err := writer.Close(); err != nil {
		return "", nil, err
	}

	return writer.FormDataContentType(), output, nil
}
