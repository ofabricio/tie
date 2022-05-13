package body_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ofabricio/tie"
	"github.com/ofabricio/tie/body"
	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {

	// Given.

	payload := strings.NewReader(`{ "message": "hello" }`)

	w := httptest.NewRecorder()

	u := tie.New(w, nil)

	// When.

	err := u.Write(http.StatusCreated, body.Copy(payload))

	// Then.

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, `{ "message": "hello" }`, w.Body.String())
}
