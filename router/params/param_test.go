package params_test

import (
	"reflect"
	"testing"

	"net/url"

	"strconv"

	"github.com/Nivl/go-rest-tools/primitives/ptrs"
	"github.com/Nivl/go-rest-tools/router/params"
	"github.com/stretchr/testify/assert"
)

func TestBasicParam(t *testing.T) {
	// sugar to avoid using true and false
	shouldFail := true

	type strct struct {
		String   string `from:"url" json:"string"`
		ID       string `from:"url" json:"id" params:"uuid"`
		Required string `from:"url" json:"required" params:"required"`
		Trim     string `from:"url" json:"trim" params:"trim"`
		Default  string `from:"url" json:"default" default:"default value"`
	}

	testCases := []struct {
		description string
		s           strct
		fieldIndex  int
		fieldName   string
		value       string
		expected    string
		shouldFail  bool
	}{
		{
			"Regular valid string should work",
			strct{}, 0, "string",
			"value", "value",
			!shouldFail,
		},
		{
			"Valid uuid should work",
			strct{}, 1, "id",
			"a2bfcbfa-5944-40b0-8930-3e5661ec4f09", "a2bfcbfa-5944-40b0-8930-3e5661ec4f09",
			!shouldFail,
		},
		{
			"Invalid uuid should fail",
			strct{}, 1, "id",
			"xxx", "",
			shouldFail,
		},
		{
			"Valid required should work",
			strct{}, 2, "required",
			"data", "data",
			!shouldFail,
		},
		{
			"Missing required should fail",
			strct{}, 2, "required",
			"", "",
			shouldFail,
		},
		{
			"Trimmed data should work",
			strct{}, 3, "trim",
			"   q   e   ", "q   e",
			!shouldFail,
		},
		{
			"Default value should be used",
			strct{}, 4, "default",
			"", "default value",
			!shouldFail,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			paramList := reflect.ValueOf(&tc.s).Elem()
			p := params.NewParamFromStructValue(&paramList, tc.fieldIndex)
			sources := &url.Values{
				tc.fieldName: []string{tc.value},
			}

			err := p.SetValue(sources)
			newValue := paramList.Field(tc.fieldIndex).String()

			if tc.shouldFail {
				assert.NotNil(t, err, "Expected SetValue to be failing with an error")
				assert.Empty(t, newValue, "Expected no value to be set")
			} else {
				assert.Nil(t, err, "Expected SetValue not to return an error")
				assert.Equal(t, tc.expected, newValue)
			}
		})
	}
}

func TestIntParam(t *testing.T) {
	// sugar to avoid using true and false
	shouldFail := true

	type strct struct {
		Number int `from:"url" json:"number"`
	}

	testCases := []struct {
		description   string
		s             strct
		value         string
		expectedValue int
		shouldFail    bool
	}{
		{"42 as int should work", strct{}, "42", 42, !shouldFail},
		{"-5 as int should work", strct{}, "-5", -5, !shouldFail},
		{"NaN as int should fail", strct{}, "NaN", 0, shouldFail},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			paramList := reflect.ValueOf(&tc.s).Elem()
			p := params.NewParamFromStructValue(&paramList, 0)

			source := &url.Values{}
			source.Set("number", tc.value)

			err := p.SetValue(source)

			if tc.shouldFail {
				assert.NotNil(t, err, "Expected SetValue to be failing with an error")
			} else {
				assert.Nil(t, err, "Expected SetValue not to return an error")
				assert.Equal(t, tc.expectedValue, tc.s.Number)
			}
		})
	}
}

func TestIntPointerParam(t *testing.T) {
	type strct struct {
		Pointer *int `from:"url" json:"pointer"`
	}

	testCases := []struct {
		description string
		s           strct
		value       *int
	}{
		{"Nil pointer should work", strct{}, nil},
		{"Pointers should work", strct{}, ptrs.NewInt(42)},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			paramList := reflect.ValueOf(&tc.s).Elem()
			p := params.NewParamFromStructValue(&paramList, 0)

			sources := &url.Values{}
			if tc.value != nil {
				sources.Set("pointer", strconv.Itoa(*tc.value))
			}

			err := p.SetValue(sources)
			assert.Nil(t, err, "Expected SetValue not to return an error")

			if tc.value == nil {
				assert.Nil(t, tc.s.Pointer, "Expected Pointer to be nil")
			} else {
				assert.Equal(t, *tc.value, *tc.s.Pointer)
			}
		})
	}
}

func TestBooleanParam(t *testing.T) {
	// sugar to avoid using true and false
	shouldFail := true

	type strct struct {
		Boolean bool `from:"url" json:"boolean"`
	}

	testCases := []struct {
		description   string
		s             strct
		value         string
		expectedValue bool
		shouldFail    bool
	}{
		{"0 as boolean should work", strct{}, "0", false, !shouldFail},
		{"false as boolean should work", strct{}, "false", false, !shouldFail},
		{"1 as boolean should work", strct{}, "1", true, !shouldFail},
		{"true as boolean should work", strct{}, "true", true, !shouldFail},
		{"xxx as boolean should fail", strct{}, "xxx", true, shouldFail},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			paramList := reflect.ValueOf(&tc.s).Elem()
			p := params.NewParamFromStructValue(&paramList, 0)

			source := &url.Values{}
			source.Set("boolean", tc.value)

			err := p.SetValue(source)

			if tc.shouldFail {
				assert.NotNil(t, err, "Expected SetValue to be failing with an error")
			} else {
				assert.Nil(t, err, "Expected SetValue not to return an error")
				assert.Equal(t, tc.expectedValue, tc.s.Boolean)
			}
		})
	}
}
