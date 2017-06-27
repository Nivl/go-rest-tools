package paginator_test

import (
	"testing"

	"github.com/Nivl/go-rest-tools/paginator"
	"github.com/Nivl/go-rest-tools/primitives/ptrs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlerParams(t *testing.T) {
	testCases := []struct {
		description         string
		params              *paginator.HandlerParams
		expectedCurrentPage int
		expectedPerPage     int
	}{
		{
			"Page 1, default PerPage",
			&paginator.HandlerParams{Page: ptrs.NewInt(1)},
			1,
			100,
		},
		{
			"Page 4, PerPage 50",
			&paginator.HandlerParams{Page: ptrs.NewInt(4), PerPage: ptrs.NewInt(50)},
			4,
			50,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()
			pginator := tc.params.Paginator(100)

			assert.Equal(t, tc.expectedCurrentPage, pginator.CurrentPage())
			assert.Equal(t, tc.expectedPerPage, pginator.PerPage())
		})
	}
}

func TestPaginator(t *testing.T) {
	testCases := []struct {
		description    string
		p              *paginator.Paginator
		shouldBeValid  bool
		expectedOffset int
		expectedLimit  int
	}{
		{
			"Page 1, PerPage 100",
			paginator.New(1, 100),
			true,
			0,
			100,
		},
		{
			"Page 6, PerPage 600",
			paginator.New(6, 600),
			false,
			0, 0,
		},
		{
			"Page 50, PerPage 10",
			paginator.New(50, 10),
			true,
			490,
			10,
		},
		{
			"Page 0, PerPage 10",
			paginator.New(0, 10),
			false, 0, 0,
		},
		{
			"Page 10, PerPage 0",
			paginator.New(10, 0),
			false, 0, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tc.shouldBeValid, tc.p.IsValid())
			if tc.shouldBeValid {
				assert.Equal(t, tc.expectedLimit, tc.p.Limit())
				assert.Equal(t, tc.expectedOffset, tc.p.Offset())
			}
		})
	}
}
