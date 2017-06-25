package params

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/Nivl/go-rest-tools/network/http/httperr"
	"github.com/Nivl/go-rest-tools/primitives/strngs"
)

// ParamOptions represent all the options for a field
type ParamOptions struct {
	// Ignore means the field should not been parsed
	// json:"-"
	Ignore bool

	// Name contains the name of the field in the payload
	// json:"my_field"
	Name string

	// Required means the request should fail with a Bad Request if the field is missing.
	// params:"required"
	Required bool

	// Trim means the field needs to be trimmed before being retrieved and checked
	// params:"trim"
	Trim bool

	// ValidateUUID means the field should contain a valid UUIDv4
	// params:"uuid"
	ValidateUUID bool

	// ValidateOptionalBool means the field should either be empty or contain a bool
	// params:"bool"
	ValidateOptionalBool bool
}

// Validate checks the given value passes the options set
func (opts *ParamOptions) Validate(value string) error {
	if value == "" && opts.Required {
		return httperr.NewBadRequest("parameter missing: %s", opts.Name)
	}

	if value != "" {
		if opts.ValidateUUID && !strngs.IsValidUUID(value) {
			return httperr.NewBadRequest("not a valid uuid: %s - %s", opts.Name, value)
		}

		if opts.ValidateOptionalBool {
			if _, err := strconv.ParseBool(value); err != nil {
				return httperr.NewBadRequest("not a valid bool: %s - %s", opts.Name, value)
			}
		}
	}

	return nil
}

// ApplyTransformations applies all the wanted transformations to the given value
func (opts *ParamOptions) ApplyTransformations(value string) string {
	if opts.Trim {
		value = strings.TrimSpace(value)
	}
	return value
}

// NewParamOptions returns a ParamOptions from a StructTag
func NewParamOptions(tags *reflect.StructTag) *ParamOptions {
	output := &ParamOptions{}

	// We use the json tag to get the field name
	jsonOpts := strings.Split(tags.Get("json"), ",")
	if len(jsonOpts) > 0 {
		if jsonOpts[0] == "-" {
			return &ParamOptions{Ignore: true}
		}

		output.Name = jsonOpts[0]
	}

	// We parse the params
	opts := strings.Split(tags.Get("params"), ",")
	nbOptions := len(opts)
	for i := 0; i < nbOptions; i++ {
		switch opts[i] {
		case "required":
			output.Required = true
		case "trim":
			output.Trim = true
		case "uuid":
			output.ValidateUUID = true
		case "bool":
			output.ValidateOptionalBool = true
		}
	}

	return output
}
