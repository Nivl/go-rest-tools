package params_test

// this file makes sure the options are set correctly and the
// validation works as intended.

import (
	"reflect"
	"testing"

	"github.com/Nivl/go-rest-tools/router/params"
	"github.com/Nivl/go-types/ptrs"
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

			err := po.Validate(tc.value, true)
			if tc.shouldFail {
				assert.Error(t, err, "the validation should have failed")
			} else {
				assert.NoError(t, err, "the validation should have succeed")
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

			err := po.Validate(tc.value, true)
			if tc.shouldFail {
				assert.Error(t, err, "the validation should have failed")
			} else {
				assert.NoError(t, err, "the validation should have succeed")
			}
		})
	}
}

func TestNoEmptyOptionOnString(t *testing.T) {
	// sugar to avoid writing true/false
	shouldFail := true

	s := struct {
		ID string `from:"url" json:"id" params:"noempty"`
	}{}

	po := getParamOptions(s)
	assert.True(t, po.NoEmpty, "ID should be checked as noempty")
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

			err := po.Validate(tc.value, true)
			if tc.shouldFail {
				assert.Error(t, err, "the validation should have failed")
			} else {
				assert.NoError(t, err, "the validation should have succeed")
			}
		})
	}
}

func TestNoEmptyOptionOnPointer(t *testing.T) {
	// sugar to avoid writing true/false
	shouldFail := true

	s := struct {
		ID *string `from:"url" json:"id" params:"noempty"`
	}{}

	po := getParamOptions(s)
	assert.True(t, po.NoEmpty, "ID should be checked as noempty")
	assert.Equal(t, po.Name, "id")

	testCases := []struct {
		description string
		value       *string
		shouldFail  bool
	}{
		{"Nil pointer are accepted", nil, !shouldFail},
		{"Empty value should fail", ptrs.NewString(""), shouldFail},
		{"Any value should work", ptrs.NewString("value"), !shouldFail},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			value := ""
			provided := false
			if tc.value != nil {
				provided = true
				value = *tc.value
			}

			err := po.Validate(value, provided)
			if tc.shouldFail {
				assert.Error(t, err, "the validation should have failed")
			} else {
				assert.NoError(t, err, "the validation should have succeed")
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

func TestMaxlenOption(t *testing.T) {
	// sugar to avoid writing true/false
	shouldFail := true

	s := struct {
		ID string `from:"url" json:"id" maxlen:"10"`
	}{}

	po := getParamOptions(s)
	assert.Equal(t, 10, po.MaxLen, "MaxLen should be set to 10")
	assert.Equal(t, po.Name, "id")

	testCases := []struct {
		description string
		value       string
		shouldFail  bool
	}{
		{"empty value should work", "", !shouldFail},
		{"5 chars should work", "12345", !shouldFail},
		{"10 chars should work", "1234567890", !shouldFail},
		{"11 chars should fail", "1234567890a", shouldFail},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			err := po.Validate(tc.value, true)
			if tc.shouldFail {
				assert.Error(t, err, "the validation should have failed")
			} else {
				assert.NoError(t, err, "the validation should have succeed")
			}
		})
	}
}

func TestEnumOption(t *testing.T) {
	// sugar to avoid writing true/false
	shouldFail := true

	s := struct {
		ID string `from:"url" json:"id" enum:"or,and"`
	}{}

	po := getParamOptions(s)
	assert.Equal(t, []string{"or", "and"}, po.AuthorizedValues, "AuthorizedValues doesn not contains the right value(s)")
	assert.Equal(t, po.Name, "id")

	testCases := []struct {
		description string
		value       string
		shouldFail  bool
	}{
		{"empty values are allowed", "", !shouldFail},
		{"'something' is not allowed", "something", shouldFail},
		{"'and' is allowed", "and", !shouldFail},
		{"'or' is allowed", "or", !shouldFail},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			err := po.Validate(tc.value, true)
			if tc.shouldFail {
				assert.Error(t, err, "the validation should have failed")
			} else {
				assert.NoError(t, err, "the validation should have succeed")
			}
		})
	}
}

// Helpers

func getParamOptions(s interface{}) *params.ParamOptions {
	st := reflect.TypeOf(s)
	tags := st.Field(0).Tag

	return params.NewParamOptions(&tags)
}
