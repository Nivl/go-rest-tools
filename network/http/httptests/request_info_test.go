package httptests_test

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/url"
	"os"
	"testing"

	"strconv"

	"mime/multipart"

	"github.com/Nivl/go-params/formfile"
	"github.com/Nivl/go-params/formfile/testformfile"
	"github.com/Nivl/go-rest-tools/network/http/httptests"
	"github.com/Nivl/go-rest-tools/router"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInfoURLParams(t *testing.T) {
	p := struct {
		ID     string `from:"url" json:"id"`
		ItemID string `from:"url" json:"item_id"`
	}{
		ID:     "a3f00e1d-7603-46bb-aaf2-82f8ff6d99fd",
		ItemID: "a244d106-1cb2-4f52-9cb5-0f758f8d5a8d",
	}

	ri := &httptests.RequestInfo{
		Endpoint: &router.Endpoint{
			Verb: "GET",
			Path: "/item/{item_id}/subitem/{id}",
		},
		Params: p,
	}

	ri.ParseParams()

	expectedURL := fmt.Sprintf("/item/%s/subitem/%s", p.ItemID, p.ID)
	assert.Equal(t, expectedURL, ri.URL())
}

func TestInfoQueryParams(t *testing.T) {
	p := struct {
		Page    int   `from:"query" json:"page"`
		PerPage int   `from:"query" json:"per_page"`
		Slice   []int `from:"query" json:"slice"`
	}{
		Page:    1,
		PerPage: 30,
		Slice:   []int{1, 2, 3},
	}

	ri := &httptests.RequestInfo{
		Endpoint: &router.Endpoint{
			Verb: "GET",
			Path: "/items",
		},
		Params: p,
	}

	ri.ParseParams()

	qs := url.Values{}
	ri.PopulateQuery(qs)

	assert.Equal(t, strconv.Itoa(p.Page), qs.Get("page"))
	assert.Equal(t, strconv.Itoa(p.PerPage), qs.Get("per_page"))
	require.Len(t, qs["slice"], len(p.Slice), "wrong number of slice parsed")
}

func TestInfoJSONBody(t *testing.T) {
	type structTest struct {
		Name  string   `from:"form" json:"name"`
		Email string   `from:"form" json:"email"`
		Slice []string `from:"form" json:"slice"`
	}

	p := &structTest{
		Name:  "User Name",
		Email: "email.domain.tld",
		Slice: []string{"1", "2", "3"},
	}

	ri := &httptests.RequestInfo{
		Endpoint: &router.Endpoint{
			Verb: "POST",
			Path: "/items",
		},
		Params: p,
	}

	ri.ParseParams()
	_, body, err := ri.Body()
	assert.NoError(t, err)
	assert.NotNil(t, body)

	var pld *structTest
	err = json.NewDecoder(body).Decode(&pld)
	require.NoError(t, err)

	assert.Equal(t, p.Name, pld.Name)
	assert.Equal(t, p.Email, pld.Email)
	assert.Equal(t, p.Slice, pld.Slice)
}

func TestInfoMultipartBody(t *testing.T) {
	cwd, _ := os.Getwd()
	type structTest struct {
		Name  string             `from:"form" json:"name"`
		Email string             `from:"form" json:"email"`
		Img   *formfile.FormFile `from:"file" json:"img"`
	}

	p := &structTest{
		Name:  "User Name",
		Email: "email.domain.tld",
		Img:   testformfile.NewFormFile(t, cwd, "black_pixel.png"),
	}
	defer p.Img.File.Close()

	ri := &httptests.RequestInfo{
		Endpoint: &router.Endpoint{
			Verb: "POST",
			Path: "/items",
		},
		Params: p,
	}

	ri.ParseParams()
	contentType, body, err := ri.Body()

	require.NoError(t, err)
	assert.NotNil(t, body)

	// We parse the multipart body
	_, params, err := mime.ParseMediaType(contentType)
	require.NoError(t, err)
	r := multipart.NewReader(body, params["boundary"])
	form, err := r.ReadForm(20 << 32)
	require.NoError(t, err)

	// Validate the "name" field
	vals, found := form.Value["name"]
	require.True(t, found, "the param \"name\" should have been found")
	require.Equal(t, 1, len(vals), "the param \"name\" should have 1 data")
	require.Equal(t, p.Name, vals[0], "name does not match its original value")

	// Validate the "email" field
	vals, found = form.Value["email"]
	require.True(t, found, "the param \"email\" should have been found")
	require.Equal(t, 1, len(vals), "the param \"email\" should have 1 data")
	require.Equal(t, p.Email, vals[0], "email does not match its original value")

	// Validate the "img" field
	headers, found := form.File["img"]
	require.True(t, found, "the param \"img\" should have been found")
	require.Equal(t, 1, len(headers), "the param \"img\" should have 1 data")
	header := headers[0]
	assert.Equal(t, "black_pixel.png", header.Filename, "the filename of img should have not changed")
	file, err := header.Open()
	require.NoError(t, err)
	file.Close()
}

func TestInfoNoParams(t *testing.T) {
	expectedURL := "/resource"

	ri := &httptests.RequestInfo{
		Endpoint: &router.Endpoint{
			Verb: "GET",
			Path: expectedURL,
		},
	}

	ri.ParseParams()
	assert.Equal(t, expectedURL, ri.URL(), "Wrong url returned")
}
