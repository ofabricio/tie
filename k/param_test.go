package k

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParamStr(t *testing.T) {

	// Given.

	tt := []struct {
		url string
		exp string
		err string
	}{
		{url: "/", exp: ""},
		{url: "/?a=", exp: ""},
		{url: "/?a=123", exp: "123"},
		{url: "/?a=mary", exp: "mary"},
		{url: "/?a=mary,john", exp: "mary,john"},
		{url: "/?a=2022-10-05T01:40:35Z", exp: "2022-10-05T01:40:35Z"},
	}

	for _, tc := range tt {

		r := httptest.NewRequest(http.MethodGet, tc.url, nil)

		// When.

		q := Param{Get: r.URL.Query().Get}

		// Then.

		assert.Equal(t, tc.exp, q.Name("a").Str(), tc)
	}
}

func TestParamInt(t *testing.T) {

	// Given.

	tt := []struct {
		url string
		exp int
		err string
	}{
		{url: "/", exp: 0},
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

		q := Param{Get: r.URL.Query().Get}

		// Then.

		assert.Equal(t, tc.exp, q.Name("a").Int(), tc.url)
		if tc.err != "" && assert.NotNil(t, q.Err, tc.url) {
			assert.Equal(t, tc.err, q.Err.Error(), tc.url)
		}
	}
}

func TestParamTime(t *testing.T) {

	// Given.

	tt := []struct {
		url string
		exp string
		err string
	}{
		{url: "/", exp: "12:00AM"},
		{url: "/?a=", exp: "12:00AM"},
		{url: "/?a=123", exp: "12:00AM", err: "key a is not a valid time layout"},
		{url: "/?a=mary", exp: "12:00AM", err: "key a is not a valid time layout"},
		{url: "/?a=5:30PM", exp: "5:30PM"},
	}

	for _, tc := range tt {

		r := httptest.NewRequest(http.MethodGet, tc.url, nil)

		// When.

		q := Param{Get: r.URL.Query().Get}

		// Then.

		assert.Equal(t, tc.exp, q.Name("a").Time(time.Kitchen).Format(time.Kitchen), tc.url)
		if tc.err != "" && assert.NotNil(t, q.Err, tc.url) {
			assert.Equal(t, tc.err, q.Err.Error(), tc.url)
		}
	}
}

func TestParamRequired(t *testing.T) {

	// Given.

	tt := []struct {
		url string
		exp string
		err string
	}{
		{url: "/", err: "key a is required"},
		{url: "/?a=", err: "key a is required"},
		{url: "/?a=hi", exp: "hi"},
	}

	for _, tc := range tt {

		r := httptest.NewRequest(http.MethodGet, tc.url, nil)

		// When.

		q := Param{Get: r.URL.Query().Get}

		// Then.

		assert.Equal(t, tc.exp, q.Name("a").Required().Str(), tc.url)
		if tc.err != "" && assert.NotNil(t, q.Err, tc.url) {
			assert.Equal(t, tc.err, q.Err.Error(), tc.url)
		}
	}
}

func TestParamSplit(t *testing.T) {

	// Given.

	tt := []struct {
		url string
		exp []string
	}{
		{url: "/", exp: nil},
		{url: "/?a=", exp: nil},
		{url: "/?a=a", exp: []string{"a"}},
		{url: "/?a=a,b", exp: []string{"a", "b"}},
	}

	for _, tc := range tt {

		r := httptest.NewRequest(http.MethodGet, tc.url, nil)

		// When.

		q := Param{Get: r.URL.Query().Get}

		// Then.

		assert.Equal(t, tc.exp, q.Name("a").Split(), tc.url)
	}
}

func TestParamDefault(t *testing.T) {

	// Given.

	tt := []struct {
		url string
		exp string
	}{
		{url: "/", exp: "10"},
		{url: "/?a=", exp: "10"},
		{url: "/?a=1", exp: "1"},
	}

	for _, tc := range tt {

		r := httptest.NewRequest(http.MethodGet, tc.url, nil)

		// When.

		q := Param{Get: r.URL.Query().Get}

		// Then.

		assert.Equal(t, tc.exp, q.Name("a").Default("10").Str(), tc.url)
	}
}

func ExampleParam() {

	r := httptest.NewRequest(http.MethodGet, "/?a=3&b=4,5,6", nil)

	// When.

	q := Param{Get: r.URL.Query().Get}

	a := q.Name("a").Str()
	b := q.Name("b").Split()
	c := q.Name("c").Int()

	fmt.Println(a, b, c, q.Err)

	// Output:
	// 3 [4 5 6] 0 <nil>
}
