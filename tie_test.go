package tie_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ofabricio/tie"
	"github.com/stretchr/testify/assert"
)

func TestBind(t *testing.T) {

	// Given.

	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"name":"mary"}`))

	u := tie.New(nil, r)

	var payload struct {
		Name string `json:"name"`
	}

	// When.

	err := u.Bind(&payload)

	// Then.

	assert.Nil(t, err)
	assert.Equal(t, "mary", payload.Name)
}

func TestQuery(t *testing.T) {

	// Given.

	r := httptest.NewRequest(http.MethodGet, "/?name=mary", nil)

	u := tie.New(nil, r)

	// When.

	name := u.Query("name")

	// Then.

	assert.Equal(t, "mary", name)
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

		u := tie.New(nil, r)

		// When.

		seg := u.PathSeg(tc.when)

		// Then.

		assert.Equal(t, tc.then, seg, tc.give)
	}
}
