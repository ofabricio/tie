package json_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ofabricio/tie/json"
	"github.com/ofabricio/tie/opt"
	"github.com/stretchr/testify/assert"
)

func TestBody_With_Msg(t *testing.T) {

	// Given.

	tt := []struct {
		name string
		when opt.WriteFunc
		then string
	}{
		{
			name: "Body",
			when: json.Body(map[string]any{"message": "hi"}),
			then: `{ "message": "hi" }`,
		},
		{
			name: "With",
			when: json.With("field_a", "value_a", "field_b", "value_b"),
			then: `{ "field_a": "value_a", "field_b": "value_b" }`,
		},
		{
			name: "Msg",
			when: json.Msg("hello %s", "world"),
			then: `{ "message": "hello world" }`,
		},
	}

	for _, tc := range tt {

		w := httptest.NewRecorder()

		// When.

		err := tc.when(w.Header())(w)

		// Then.

		assert.Nil(t, err)
		assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
		assert.JSONEq(t, tc.then, w.Body.String())
	}
}

func TestND(t *testing.T) {

	// Given.

	type Payload struct {
		Name string `json:"name"`
	}

	w := httptest.NewRecorder()

	ch := make(chan *Payload)
	go func() {
		ch <- &Payload{Name: "Mary"}
		ch <- &Payload{Name: "John"}
		close(ch)
	}()

	// When.

	err := json.ND(ch)(w.Header())(w)

	// Then.

	assert.Nil(t, err)
	assert.Equal(t, "application/x-ndjson; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, "{\"name\":\"Mary\"}\n{\"name\":\"John\"}\n", w.Body.String())
}

func TestBind(t *testing.T) {

	// Given.

	var payload struct {
		Name string `json:"name"`
	}

	r := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(`{ "name": "Mary" }`))

	// When.

	err := json.Bind(&payload)(r)

	// Then.

	assert.Nil(t, err)
	assert.Equal(t, "Mary", payload.Name)
}

func TestBindND(t *testing.T) {

	// Given.

	type payload struct {
		Name string `json:"name"`
	}

	r := httptest.NewRequest(http.MethodPut, "/",
		strings.NewReader("{\"name\":\"Mary\"}\n{\"name\":\"John\"}"))

	ch := make(chan payload)

	// When.

	go json.BindND(ch)(r)

	a := <-ch
	b := <-ch

	// Then.

	assert.Equal(t, "Mary", a.Name)
	assert.Equal(t, "John", b.Name)
}
