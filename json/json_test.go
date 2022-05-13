package json_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ofabricio/tie"
	"github.com/ofabricio/tie/json"
	"github.com/stretchr/testify/assert"
)

func TestBody_With_Msg(t *testing.T) {

	// Given.

	tt := []struct {
		name string
		when tie.WriteFunc
		then string
	}{
		{
			name: "Body",
			when: json.Body(map[string]any{"message": "hi"}),
			then: `{ "message": "hi" }`,
		},
		{
			name: "JsonKV",
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

func TestND(t *testing.T) {

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

	err := u.Write(http.StatusCreated, json.ND(ch))

	// Then.

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "application/x-ndjson; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, "{\"name\":\"Mary\"}\n{\"name\":\"John\"}\n", w.Body.String())
}
