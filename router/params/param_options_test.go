package params_test

import (
	"reflect"
	"testing"

	"github.com/Nivl/go-rest-tools/router/params"
	"github.com/stretchr/testify/assert"
)

func TestUUIDOption(t *testing.T) {
	// sugar to avoid writing true/false
	shouldFail := true

	s := struct {
		ID string `from:"url" json:"id" params:"uuid"`
	}{}

	po := getParamOptions(s)
	assert.True(t, po.ValidateUUID, "UUId should be set for validation")
	assert.Equal(t, po.Name, "id")

	testCases := []struct {
		description string
		value       string
		shouldFail  bool
	}{
		{"Invalid uuid should NOT pass validation", "xxx", shouldFail},
		{"Empty uuid should pass validation", "", !shouldFail}, // cause not "required"
		{"Valid uuid should pass validation", "1db81012-ce7d-4445-aafe-7d8343636685", !shouldFail},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			err := po.Validate(tc.value)
			if tc.shouldFail {
				assert.NotNil(t, err, "the validation should have failed")
			} else {
				assert.Nil(t, err, "the validation should have succeed")
			}
		})
	}
}

func TestRequiredOption(t *testing.T) {
	// sugar to avoid writing true/false
	shouldFail := true

	s := struct {
		ID string `from:"url" json:"id" params:"required"`
	}{}

	po := getParamOptions(s)
	assert.True(t, po.Required, "ID should be checked as required")
	assert.Equal(t, po.Name, "id")

	testCases := []struct {
		description string
		value       string
		shouldFail  bool
	}{
		{"Missing ID should NOT pass the validation", "", shouldFail},
		{"Valid ID should pass the validation", "xxx", !shouldFail},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			err := po.Validate(tc.value)
			if tc.shouldFail {
				assert.NotNil(t, err, "the validation should have failed")
			} else {
				assert.Nil(t, err, "the validation should have succeed")
			}
		})
	}
}

func TestIgnoredOption(t *testing.T) {
	s := struct {
		ID string `from:"url" json:"-"`
	}{}

	po := getParamOptions(s)
	assert.True(t, po.Ignore, "ID should be marked as ignored")
	assert.Empty(t, po.Name, "ID should have no name")
}

func TestTrimOption(t *testing.T) {
	s := struct {
		ID string `from:"url" json:"id" params:"trim"`
	}{}

	po := getParamOptions(s)
	assert.True(t, po.Trim, "ID should be marked for trimming")
	assert.Equal(t, po.Name, "id")

	testCases := []struct {
		description    string
		input          string
		expectedOutput string
	}{
		{"Empty string", "", ""},
		{"Spaces only", "     ", ""},
		{"Word wrapped around spaces", "  word   ", "word"},
		{"Two words wrapped in spaces", "  two words  ", "two words"},
		{"No spaces at all", "a", "a"},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			output := po.ApplyTransformations(tc.input)
			assert.Equal(t, tc.expectedOutput, output)
		})
	}
}

// Helpers

func getParamOptions(s interface{}) *params.ParamOptions {
	st := reflect.TypeOf(s)
	tags := st.Field(0).Tag

	return params.NewParamOptions(&tags)
}