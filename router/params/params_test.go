package params_test

import (
	"net/url"
	"os"
	"testing"

	"strconv"

	"github.com/Nivl/go-rest-tools/primitives/ptrs"
	"github.com/Nivl/go-rest-tools/router/formfile"
	"github.com/Nivl/go-rest-tools/router/formfile/testformfile"
	"github.com/Nivl/go-rest-tools/router/params"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidStruct(t *testing.T) {
	type strct struct {
		ID            string  `from:"url" json:"id" params:"uuid,required"`
		Number        int     `from:"query" json:"number"`
		RequiredBool  bool    `from:"form" json:"required_bool" params:"required"`
		PointerBool   *bool   `from:"form" json:"pointer_bool"`
		PointerString *string `from:"form" json:"pointer_string" params:"trim"`
		Default       int     `from:"form" json:"default" default:"42"`
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
	assert.NotNil(t, err)
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
	}{
		String:        "String value",
		Number:        42,
		Bool:          true,
		PointerBool:   ptrs.NewBool(false),
		PointerString: ptrs.NewString("string pointer"),
		PointerNumber: ptrs.NewInt(24),
		Nil:           nil,
		File:          testformfile.NewFormFile(t, cwd, "black_pixel.png"),
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
}
