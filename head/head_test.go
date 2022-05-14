package head_test

import (
	"net/http/httptest"
	"testing"

	"github.com/ofabricio/tie/head"
	"github.com/stretchr/testify/assert"
)

func TestWith(t *testing.T) {

	// Given.

	w := httptest.NewRecorder()

	// When.

	bodyFunc := head.With("A", "aaa", "B", "bbb")(w.Header())

	// Then.

	assert.Nil(t, bodyFunc)
	assert.Equal(t, "aaa", w.Header().Get("A"))
	assert.Equal(t, "bbb", w.Header().Get("B"))
}
