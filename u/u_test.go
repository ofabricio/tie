package u_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ofabricio/tie/opt"
	"github.com/ofabricio/tie/u"
	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {

	// Given.

	w := httptest.NewRecorder()

	// When.

	u.Write(w, http.StatusCreated, func(h http.Header) opt.BodyFunc {
		h.Add("X-Version", "v1")
		return func(w io.Writer) error {
			w.Write([]byte("hello"))
			return nil
		}
	})

	// Then.

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "v1", w.Header().Get("X-Version"))
	assert.Equal(t, "hello", w.Body.String())
}

func TestHead(t *testing.T) {

	// Given.

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("X-Version", "123")

	// When.

	v := u.Head(r).Int("X-Version")

	// Then.

	assert.Equal(t, 123, v)
}

func TestPathSeg(t *testing.T) {

	// Given.

	tt := []struct {
		give string
		when int
		then string
	}{
		{give: "a", when: 0, then: "a"},
		{give: "a", when: 1, then: ""},
		{give: "a/", when: 0, then: "a"},
		{give: "a/", when: 1, then: ""},
		{give: "/a/", when: 0, then: "a"},
		{give: "/a/", when: 1, then: ""},
		{give: "/a", when: 0, then: "a"},
		{give: "/a", when: 1, then: ""},
		{give: "a/b", when: 0, then: "a"},
		{give: "a/b", when: 1, then: "b"},
		{give: "a/b", when: 2, then: ""},
		{give: "a/b/c", when: 2, then: "c"},
	}

	for _, tc := range tt {

		r := httptest.NewRequest(http.MethodGet, "/"+tc.give, nil)

		// When.

		seg := u.PathSeg(r, tc.when).Str()

		// Then.

		assert.Equal(t, tc.then, seg, tc.give)
	}
}

func TestRead(t *testing.T) {

	// Given.

	var buf bytes.Buffer
	r := httptest.NewRequest(http.MethodPut, "/", strings.NewReader("Hello"))

	// When.

	err := u.Read(r, func(r *http.Request) error {
		d, err := io.ReadAll(r.Body)
		buf.Write(d)
		return err
	})

	// Then.

	assert.Nil(t, err)
	assert.Equal(t, "Hello", buf.String())
}
