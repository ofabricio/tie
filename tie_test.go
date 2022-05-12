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

func TestWrite(t *testing.T) {

	// Given.

	tt := []struct {
		name string
		when tie.WriteFunc
		then string
	}{
		{
			name: "Json",
			when: tie.Json(map[string]any{"message": "hi"}),
			then: `{ "message": "hi" }`,
		},
		{
			name: "JsonKV",
			when: tie.JsonKV("field_a", "value_a", "field_b", "value_b"),
			then: `{ "field_a": "value_a", "field_b": "value_b" }`,
		},
		{
			name: "Msg",
			when: tie.Msg("hello %s", "world"),
			then: `{ "message": "hello world" }`,
		},
	}

	for _, tc := range tt {

		w := httptest.NewRecorder()

		u := tie.New(w, nil)

		// When.

		err := u.Write(http.StatusCreated, tc.when)

		// Then.

		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
		assert.JSONEq(t, tc.then, w.Body.String())
	}
}

func TestNDJson(t *testing.T) {

	// Given.

	type Payload struct {
		Name string `json:"name"`
	}

	w := httptest.NewRecorder()

	u := tie.New(w, nil)

	ch := make(chan *Payload)
	go func() {
		ch <- &Payload{Name: "Mary"}
		ch <- &Payload{Name: "John"}
		close(ch)
	}()

	// When.

	err := u.Write(http.StatusCreated, tie.NDJson(ch))

	// Then.

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "application/x-ndjson; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, "{\"name\":\"Mary\"}\n{\"name\":\"John\"}\n", w.Body.String())
}

func TestCopy(t *testing.T) {

	// Given.

	payload := strings.NewReader(`{ "message": "hello" }`)

	w := httptest.NewRecorder()

	u := tie.New(w, nil)

	// When.

	err := u.Write(http.StatusCreated, tie.Copy(payload))

	// Then.

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, `{ "message": "hello" }`, w.Body.String())
}

func TestHeader(t *testing.T) {

	// Given.

	w := httptest.NewRecorder()

	u := tie.New(w, nil)

	// When.

	err := u.Write(http.StatusCreated, tie.Header("A", "aaa", "B", "bbb"))

	// Then.

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "aaa", w.Header().Get("A"))
	assert.Equal(t, "bbb", w.Header().Get("B"))
}
