package head_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ofabricio/tie"
	"github.com/ofabricio/tie/head"
	"github.com/stretchr/testify/assert"
)

func TestWith(t *testing.T) {

	// Given.

	w := httptest.NewRecorder()

	u := tie.New(w, nil)

	// When.

	err := u.Write(http.StatusCreated, head.With("A", "aaa", "B", "bbb"))

	// Then.

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "aaa", w.Header().Get("A"))
	assert.Equal(t, "bbb", w.Header().Get("B"))
}
