package tie_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ofabricio/tie"
	"github.com/stretchr/testify/assert"
)

func TestJson(t *testing.T) {

	// Given.

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	u := tie.New(w, r)

	// When.

	err := u.Json(http.StatusCreated, map[string]string{"msg": "hi"})

	// Then.

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"msg":"hi"}`, w.Body.String())
}

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

func TestQueryParam(t *testing.T) {

	// Given.

	r := httptest.NewRequest(http.MethodGet, "/?name=mary", nil)

	u := tie.New(nil, r)

	// When.

	name := u.QueryParam("name")

	// Then.

	assert.Equal(t, "mary", name)
}
