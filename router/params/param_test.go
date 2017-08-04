package params_test

import (
	"reflect"
	"testing"

	"net/http"
	"net/url"

	"strconv"

	"os"

	"github.com/Nivl/go-rest-tools/types/ptrs"
	"github.com/Nivl/go-rest-tools/router/formfile"
	"github.com/Nivl/go-rest-tools/router/formfile/mockformfile"
	"github.com/Nivl/go-rest-tools/router/formfile/testformfile"
	"github.com/Nivl/go-rest-tools/router/params"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		Url      string `from:"url" json:"url" params:"url"`
		Email    string `from:"url" json:"email" params:"email"`
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
		{
			"Valid URL should work",
			strct{}, 5, "url",
			"http://google.com", "http://google.com",
			!shouldFail,
		},
		{
			"invalid URL should fail",
			strct{}, 5, "url",
			"ftp://google.com", "",
			shouldFail,
		},
		{
			"Valid email should work",
			strct{}, 6, "email",
			"email@domain.tld", "email@domain.tld",
			!shouldFail,
		},
		{
			"invalid email should fail",
			strct{}, 6, "email",
			"not-an-email", "",
			shouldFail,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			paramList := reflect.ValueOf(&tc.s).Elem()
			p := params.NewParamFromStructValue(&paramList, tc.fieldIndex)
			sources := url.Values{
				tc.fieldName: []string{tc.value},
			}

			err := p.SetValue(sources)
			newValue := paramList.Field(tc.fieldIndex).String()

			if tc.shouldFail {
				assert.Error(t, err, "Expected SetValue to be failing with an error")
				assert.Empty(t, newValue, "Expected no value to be set")
			} else {
				assert.NoError(t, err, "Expected SetValue not to return an error")
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

			source := url.Values{}
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

			sources := url.Values{}
			if tc.value != nil {
				sources.Set("pointer", strconv.Itoa(*tc.value))
			}

			err := p.SetValue(sources)
			assert.Nil(t, err, "Expected SetValue not to return an error")

			if tc.value == nil {
				assert.Nil(t, tc.s.Pointer, "Expected Pointer to be nil")
			} else {
				require.NotNil(t, tc.s.Pointer, "the pointer should not be nil")
				assert.Equal(t, *tc.value, *tc.s.Pointer, "the value should not have changed")
			}
		})
	}
}

func TestStringPointerParam(t *testing.T) {
	type strct struct {
		Pointer *string `from:"url" json:"pointer"`
	}

	testCases := []struct {
		description string
		s           strct
		value       *string
	}{
		{"Nil pointer should work", strct{}, nil},
		{"A value should work", strct{}, ptrs.NewString("something")},
		{"Empty value should work", strct{}, ptrs.NewString("")},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			paramList := reflect.ValueOf(&tc.s).Elem()
			p := params.NewParamFromStructValue(&paramList, 0)

			sources := url.Values{}
			if tc.value != nil {
				sources.Set("pointer", *tc.value)
			}

			err := p.SetValue(sources)
			assert.Nil(t, err, "Expected SetValue not to return an error")

			if tc.value == nil {
				assert.Nil(t, tc.s.Pointer, "Expected Pointer to be nil")
			} else {
				require.NotNil(t, tc.s.Pointer, "the pointer should not be nil")
				assert.Equal(t, *tc.value, *tc.s.Pointer, "the value should not have changed")
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

			source := url.Values{}
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

func TestFileParamValid(t *testing.T) {
	type strct struct {
		File *formfile.FormFile `from:"file" json:"file"`
	}

	testCases := []struct {
		description string
		s           strct
		sendNil     bool
	}{
		{"Nil pointer should work", strct{}, true},
		{"Pointers should work", strct{}, false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			// create the multipart data
			cwd, _ := os.Getwd()
			licenseHeader, licenseFile := testformfile.NewMultipartData(t, cwd, "LICENSE")
			defer licenseFile.Close()

			// Expectation
			fileHolder := new(mockformfile.FileHolder)
			onFormFile := fileHolder.On("FormFile", "file")

			if tc.sendNil {
				onFormFile.Return(nil, nil, http.ErrMissingFile)
			} else {
				onFormFile.Return(licenseFile, licenseHeader, nil)
			}

			paramList := reflect.ValueOf(&tc.s).Elem()
			p := params.NewParamFromStructValue(&paramList, 0)

			err := p.SetFile(fileHolder)
			fileHolder.AssertExpectations(t)
			assert.Nil(t, err, "Expected SetFile not to return an error")

			if tc.sendNil {
				assert.Nil(t, tc.s.File, "Expected File to be nil")
			} else {
				if assert.NotNil(t, tc.s.File, "Expected File NOT to be nil") {
					assert.NotNil(t, tc.s.File.File, "Expected File.File NOT to be nil")
					assert.NotNil(t, tc.s.File.Header, "Expected File.Header NOT to be nil")
					assert.Equal(t, licenseHeader.Filename, tc.s.File.Header.Filename)
				}
			}
		})
	}
}

func TestFileParamRequired(t *testing.T) {
	type strct struct {
		File *formfile.FormFile `from:"file" json:"file" params:"required"`
	}

	testCases := []struct {
		description string
		s           strct
		sendNil     bool
		shouldFail  bool
	}{
		{"Nil pointer should work", strct{}, true, true},
		{"Pointers should work", strct{}, false, false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			// create the multipart data
			cwd, _ := os.Getwd()
			licenseHeader, licenseFile := testformfile.NewMultipartData(t, cwd, "LICENSE")
			defer licenseFile.Close()

			// Expectations
			fileHolder := new(mockformfile.FileHolder)
			onFormFile := fileHolder.On("FormFile", "file")

			if tc.sendNil {
				onFormFile.Return(nil, nil, http.ErrMissingFile)
			} else {
				onFormFile.Return(licenseFile, licenseHeader, nil)
			}

			paramList := reflect.ValueOf(&tc.s).Elem()
			p := params.NewParamFromStructValue(&paramList, 0)

			err := p.SetFile(fileHolder)
			fileHolder.AssertExpectations(t)

			if tc.sendNil {
				assert.Error(t, err, "Expected SetFile to return an error")

			} else {
				assert.NoError(t, err, "Expected SetFile not to return an error")

				if assert.NotNil(t, tc.s.File, "Expected File NOT to be nil") {
					assert.NotNil(t, tc.s.File.File, "Expected File.File NOT to be nil")
					assert.NotNil(t, tc.s.File.Header, "Expected File.Header NOT to be nil")
					assert.Equal(t, licenseHeader.Filename, tc.s.File.Header.Filename)
				}
			}
		})
	}
}

func TestFileParamValidImage(t *testing.T) {
	type strct struct {
		File *formfile.FormFile `from:"file" json:"file" params:"image"`
	}

	testCases := []struct {
		description  string
		s            strct
		filename     string
		expectedMime string
	}{
		{"Valid PNG", strct{}, "black_pixel.png", "image/png"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			// create the multipart data
			cwd, _ := os.Getwd()
			imageHeader, imageFile := testformfile.NewMultipartData(t, cwd, tc.filename)
			defer imageFile.Close()

			// Set the expectations
			fileHolder := new(mockformfile.FileHolder)
			fileHolder.On("FormFile", "file").Return(imageFile, imageHeader, nil)

			// Call the function to test
			paramList := reflect.ValueOf(&tc.s).Elem()
			p := params.NewParamFromStructValue(&paramList, 0)
			err := p.SetFile(fileHolder)

			// assert
			fileHolder.AssertExpectations(t)
			assert.NoError(t, err, "Expected SetFile to return an error")
			assert.Equal(t, tc.expectedMime, tc.s.File.Mime, "Wrong mime type")
		})
	}
}

func TestFileParamInvalidImage(t *testing.T) {
	type strct struct {
		File *formfile.FormFile `from:"file" json:"file" params:"image"`
	}

	testCases := []struct {
		description string
		s           strct
		filename    string
	}{
		{"Invalid magic", strct{}, "invalid_magic.png"},
		{"Invalid content", strct{}, "invalid_content.png"},
		{"Not an image", strct{}, "LICENSE"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			// create the multipart data
			cwd, _ := os.Getwd()
			imageHeader, imageFile := testformfile.NewMultipartData(t, cwd, tc.filename)
			defer imageFile.Close()

			// Set the expectations
			fileHolder := new(mockformfile.FileHolder)
			fileHolder.On("FormFile", "file").Return(imageFile, imageHeader, nil)

			// Call the function to test
			paramList := reflect.ValueOf(&tc.s).Elem()
			p := params.NewParamFromStructValue(&paramList, 0)
			err := p.SetFile(fileHolder)

			// Assert
			fileHolder.AssertExpectations(t)
			assert.Error(t, err, "Expected SetFile to return an error")
			assert.Equal(t, "not a valid image", err.Error())
		})
	}
}
