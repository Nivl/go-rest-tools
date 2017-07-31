package ptrs_test

import (
	"fmt"
	"testing"

	"github.com/Nivl/go-rest-tools/primitives/ptrs"
	"github.com/stretchr/testify/assert"
)

func TestUnwrap(t *testing.T) {
	var nilInt *int
	var nilString *string

	testCases := []struct {
		description string
		value       interface{}
		expected    interface{}
	}{
		// integers
		{"int ptr 1", ptrs.NewInt(1), 1},
		{"int 1", 1, 1},
		{"int ptr nil", nilInt, 0},

		// string
		{`string ptr "str"`, ptrs.NewString("str"), "str"},
		{`string "str"`, "str", "str"},
		{`string ptr nil`, nilString, ""},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			ret := ptrs.Unwrap(tc.value)
			fmt.Println(ret)
			assert.Equal(t, tc.expected, ret, "Unwrap didn't return the expected value")
		})
	}
}
