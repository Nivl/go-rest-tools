package params_test

// This file mainly tests that a struct of params get all its fields filled with
// the right data.

import (
	"errors"
	"net/url"
	"os"
	"testing"

	"strconv"

	"github.com/Nivl/go-rest-tools/router/formfile"
	"github.com/Nivl/go-rest-tools/router/formfile/testformfile"
	"github.com/Nivl/go-rest-tools/router/params"
	"github.com/Nivl/go-rest-tools/storage/db"
	"github.com/Nivl/go-rest-tools/types/ptrs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type StructWithValidator struct {
	String     string `from:"query" json:"string" default:"default value"`
	TrueToFail bool   `from:"query" json:"true_to_fail" default:"false"`
}

func (p *StructWithValidator) IsValid() (isValid bool, fieldName string, err error) {
	if !p.TrueToFail {
		return true, "", nil
	}
	return false, "true_to_fail", errors.New("cannot be set to true")
}

func TestValidStruct(t *testing.T) {
	type strct struct {
		ID            string  `from:"url" json:"id" params:"uuid,required"`
		Number        int     `from:"query" json:"number"`
		RequiredBool  bool    `from:"form" json:"required_bool" params:"required"`
		PointerBool   *bool   `from:"form" json:"pointer_bool"`
		PointerString *string `from:"form" json:"pointer_string" params:"trim"`
		Default       int     `from:"form" json:"default" default:"42"`
		Emum          int     `from:"form" json:"enum" enum:"21,42"`
	}

	s := &strct{}
	p := params.NewParams(s)

	urlSource := url.Values{}
	urlSource.Set("id", "1aa75114-6117-4908-b6ea-0d22ecdd4fc0")

	querySource := url.Values{}
	querySource.Set("number", "24")

	formSource := url.Values{}
	formSource.Set("required_bool", "true")
	formSource.Set("pointer_string", "     pointer value      ")
	formSource.Set("enum", "42")

	sources := map[string]url.Values{
		"url":   urlSource,
		"form":  formSource,
		"query": querySource,
	}

	if err := p.Parse(sources, nil); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "1aa75114-6117-4908-b6ea-0d22ecdd4fc0", s.ID)
	assert.Equal(t, 24, s.Number)
	assert.True(t, s.RequiredBool)
	assert.Nil(t, s.PointerBool)
	assert.Equal(t, "pointer value", *s.PointerString)
	assert.Equal(t, 42, s.Default)
}

func TestInvalidStruct(t *testing.T) {
	type strct struct {
		ID            string  `from:"url" json:"id" params:"uuid,required"`
		Number        int     `from:"query" json:"number"`
		RequiredBool  bool    `from:"form" json:"required_bool" params:"required"`
		PointerBool   *bool   `from:"form" json:"pointer_bool"`
		PointerString *string `from:"form" json:"pointer_string" params:"trim"`
		Default       int     `from:"form" json:"default" default:"42"`
	}

	p := params.NewParams(&strct{})
	err := p.Parse(map[string]url.Values{}, nil)
	assert.Error(t, err)
}

func TestEmbeddedStruct(t *testing.T) {
	type Paginator struct {
		Page    *int `from:"query" json:"page" default:"1"`
		PerPage *int `from:"query" json:"per_page"`
	}

	type strct struct {
		Paginator

		ID string `from:"url" json:"id" params:"uuid,required"`
	}

	s := &strct{}
	p := params.NewParams(s)

	urlSource := url.Values{}
	urlSource.Set("id", "1aa75114-6117-4908-b6ea-0d22ecdd4fc0")

	querySource := url.Values{}
	querySource.Set("page", "24")

	sources := map[string]url.Values{
		"url":   urlSource,
		"query": querySource,
	}

	if err := p.Parse(sources, nil); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "1aa75114-6117-4908-b6ea-0d22ecdd4fc0", s.ID)
	assert.Equal(t, 24, *s.Page)
	assert.Nil(t, s.PerPage)
}

func TestEmbeddedStructWithCustomValidation(t *testing.T) {
	// sugar
	shouldFail := true

	type strct struct {
		StructWithValidator
	}

	testCases := []struct {
		description string
		params      url.Values
		shouldFail  bool
	}{
		{
			"Trigger a failure",
			url.Values{
				"string":       []string{"value"},
				"true_to_fail": []string{"true"},
			},
			shouldFail,
		},
		{
			"Valid params should work",
			url.Values{
				"string": []string{"value"},
			},
			!shouldFail,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			s := &strct{}
			p := params.NewParams(s)
			sources := map[string]url.Values{
				"query": tc.params,
			}

			err := p.Parse(sources, nil)
			if tc.shouldFail {
				assert.Error(t, err, "Parse() should have failed")
			} else {
				assert.NoError(t, err, "Parse() should have succeed")
			}
		})
	}
}

func TestCustomValidation(t *testing.T) {
	// sugar
	shouldFail := true

	testCases := []struct {
		description string
		params      url.Values
		shouldFail  bool
	}{
		{
			"Trigger a failure",
			url.Values{
				"string":       []string{"value"},
				"true_to_fail": []string{"true"},
			},
			shouldFail,
		},
		{
			"Valid params should work",
			url.Values{
				"string": []string{"value"},
			},
			!shouldFail,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			s := &StructWithValidator{}
			p := params.NewParams(s)
			sources := map[string]url.Values{
				"query": tc.params,
			}

			err := p.Parse(sources, nil)
			if tc.shouldFail {
				assert.Error(t, err, "Parse() should have failed")
			} else {
				assert.NoError(t, err, "Parse() should have succeed")
			}
		})
	}
}

func TestExtraction(t *testing.T) {
	cwd, _ := os.Getwd()

	s := struct {
		String        string             `from:"url" json:"string"`
		Number        int                `from:"query" json:"number"`
		Bool          bool               `from:"form" json:"bool"`
		PointerBool   *bool              `from:"form" json:"pointer_bool"`
		PointerString *string            `from:"form" json:"pointer_string"`
		PointerNumber *int               `from:"form" json:"pointer_number"`
		Nil           *int               `from:"form" json:"nil"`
		File          *formfile.FormFile `from:"file" json:"file"`
		Stringer      *db.Date           `from:"form" json:"stringer"`
	}{
		String:        "String value",
		Number:        42,
		Bool:          true,
		PointerBool:   ptrs.NewBool(false),
		PointerString: ptrs.NewString("string pointer"),
		PointerNumber: ptrs.NewInt(24),
		Nil:           nil,
		File:          testformfile.NewFormFile(t, cwd, "black_pixel.png"),
		Stringer:      db.Today(),
	}

	p := params.NewParams(&s)
	sources, files := p.Extract()

	// Check file data
	fileData, found := files["file"]
	require.True(t, found, "file should be present")
	assert.NotNil(t, fileData, "fileData should not be nil")
	assert.NotNil(t, fileData.File, "fileData.File should not be nil")
	assert.NotNil(t, fileData.Header, "fileData.header should not be nil")
	assert.Equal(t, "image/png", fileData.Mime)

	// Check url data
	urlValue, found := sources["url"]
	require.True(t, found, "url data should be present")
	assert.Equal(t, urlValue.Get("string"), s.String)

	// Check query data
	queryValue, found := sources["query"]
	require.True(t, found, "query data should be present")
	assert.Equal(t, queryValue.Get("number"), strconv.Itoa(s.Number))

	// Check form data
	formValue, found := sources["form"]
	require.True(t, found, "for data should be present")
	assert.Equal(t, formValue.Get("bool"), strconv.FormatBool(s.Bool))
	assert.Equal(t, formValue.Get("pointer_bool"), strconv.FormatBool(*s.PointerBool))
	assert.Equal(t, formValue.Get("pointer_string"), *s.PointerString)
	assert.Equal(t, formValue.Get("pointer_number"), strconv.Itoa(*s.PointerNumber))
	assert.Empty(t, formValue.Get("nil"))
	d, err := db.NewDate(formValue.Get("stringer"))
	assert.NoError(t, err, "db.NewDate() should have succeed")
	assert.True(t, s.Stringer.Equal(d), "The date changed from %s to %s", s.Stringer, d)
}
