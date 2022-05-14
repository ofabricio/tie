package k_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ofabricio/tie/k"
	"github.com/stretchr/testify/assert"
)

func TestStr(t *testing.T) {

	// Given.

	tt := []struct {
		url string
		exp string
		err string
	}{
		{url: "/?a=", exp: ""},
		{url: "/?a=123", exp: "123"},
		{url: "/?a=mary", exp: "mary"},
		{url: "/?a=mary,john", exp: "mary,john"},
		{url: "/?a=2022-10-05T01:40:35Z", exp: "2022-10-05T01:40:35Z"},
	}

	for _, tc := range tt {

		r := httptest.NewRequest(http.MethodGet, tc.url, nil)

		// When.

		q := k.Param{Get: r.URL.Query().Get}

		// Then.

		assert.Equal(t, tc.exp, q.Str("a"), tc)
	}
}

func TestInt(t *testing.T) {

	// Given.

	tt := []struct {
		url string
		exp int
		err string
	}{
		{url: "/?a=", exp: 0},
		{url: "/?a=0", exp: 0},
		{url: "/?a=1", exp: 1},
		{url: "/?a=123", exp: 123},
		{url: "/?a=mary", exp: 0, err: "key a is not a valid value"},
		{url: "/?a=mary,john", exp: 0, err: "key a is not a valid value"},
		{url: "/?a=1,2", exp: 0, err: "key a is not a valid value"},
		{url: "/?a=2022-10-05T01:40:35Z", exp: 0, err: "key a is not a valid value"},
	}

	for _, tc := range tt {

		r := httptest.NewRequest(http.MethodGet, tc.url, nil)

		// When.

		q := k.Param{Get: r.URL.Query().Get}

		// Then.

		assert.Equal(t, tc.exp, q.Int("a"), tc.url)
		if tc.err != "" && assert.NotNil(t, q.Err, tc.url) {
			assert.Equal(t, tc.err, q.Err.Error(), tc.url)
		}
	}
}

func TestTime(t *testing.T) {

	// Given.

	tt := []struct {
		url string
		exp string
		err string
	}{
		{url: "/?a=", exp: "12:00AM"},
		{url: "/?a=123", exp: "12:00AM", err: "key a is not a valid time layout"},
		{url: "/?a=mary", exp: "12:00AM", err: "key a is not a valid time layout"},
		{url: "/?a=5:30PM", exp: "5:30PM"},
	}

	for _, tc := range tt {

		r := httptest.NewRequest(http.MethodGet, tc.url, nil)

		// When.

		q := k.Param{Get: r.URL.Query().Get}

		// Then.

		assert.Equal(t, tc.exp, q.Time("a", time.Kitchen).Format(time.Kitchen), tc.url)
		if tc.err != "" && assert.NotNil(t, q.Err, tc.url) {
			assert.Equal(t, tc.err, q.Err.Error(), tc.url)
		}
	}
}

func TestRequired(t *testing.T) {

	// Given.

	tt := []struct {
		url string
		exp string
		err string
	}{
		{url: "/?a=", err: "key a is required"},
		{url: "/?a=hi", exp: "hi"},
	}

	for _, tc := range tt {

		r := httptest.NewRequest(http.MethodGet, tc.url, nil)

		// When.

		q := k.Param{Get: r.URL.Query().Get}

		// Then.

		assert.Equal(t, tc.exp, q.Required("a").Str(), tc.url)
		if tc.err != "" && assert.NotNil(t, q.Err, tc.url) {
			assert.Equal(t, tc.err, q.Err.Error(), tc.url)
		}
	}
}

func TestSplit(t *testing.T) {

	// Given.

	tt := []struct {
		url string
		exp []string
	}{
		{url: "/?a=", exp: []string{""}},
		{url: "/?a=a", exp: []string{"a"}},
		{url: "/?a=a,b", exp: []string{"a", "b"}},
	}

	for _, tc := range tt {

		r := httptest.NewRequest(http.MethodGet, tc.url, nil)

		// When.

		q := k.Param{Get: r.URL.Query().Get}

		// Then.

		assert.Equal(t, tc.exp, q.Split("a").Values(), tc.url)
	}
}
