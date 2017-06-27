package params_test

import (
	"net/url"
	"testing"

	"github.com/Nivl/go-rest-tools/router/params"
	"github.com/stretchr/testify/assert"
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

	if err := p.Parse(sources); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "1aa75114-6117-4908-b6ea-0d22ecdd4fc0", s.ID)
	assert.Equal(t, 24, s.Number)
	assert.True(t, s.RequiredBool)
	assert.Nil(t, s.PointerBool)
	assert.Equal(t, "pointer value", *s.PointerString)
	assert.Equal(t, 42, s.Default)
}

func TestinvalidStruct(t *testing.T) {
	type strct struct {
		ID            string  `from:"url" json:"id" params:"uuid,required"`
		Number        int     `from:"query" json:"number"`
		RequiredBool  bool    `from:"form" json:"required_bool" params:"required"`
		PointerBool   *bool   `from:"form" json:"pointer_bool"`
		PointerString *string `from:"form" json:"pointer_string" params:"trim"`
		Default       int     `from:"form" json:"default" default:"42"`
	}

	p := params.NewParams(&strct{})
	err := p.Parse(map[string]url.Values{})
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

	if err := p.Parse(sources); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "1aa75114-6117-4908-b6ea-0d22ecdd4fc0", s.ID)
	assert.Equal(t, 24, *s.Page)
	assert.Nil(t, s.PerPage)
}
