package body_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ofabricio/tie/body"
	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {

	// Given.

	payload := strings.NewReader(`{ "message": "hello" }`)

	w := httptest.NewRecorder()

	// When.

	err := body.Copy(payload)(w.Header())(w)

	// Then.

	assert.Nil(t, err)
	assert.Equal(t, `{ "message": "hello" }`, w.Body.String())
}

func TestCopyTo(t *testing.T) {

	// Given.

	var buf bytes.Buffer

	r := httptest.NewRequest(http.MethodPut, "/", strings.NewReader("Hello"))

	// When.

	err := body.CopyTo(&buf)(r)

	// Then.

	assert.Nil(t, err)
	assert.Equal(t, "Hello", buf.String())
}
