package dump_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/ofabricio/tie/dump"
	"github.com/stretchr/testify/assert"
)

func TestRes(t *testing.T) {

	res := &http.Response{
		StatusCode: 201,
		Header:     http.Header{"A": []string{"B"}},
		Body:       io.NopCloser(strings.NewReader("Hello")),
	}

	d := dump.Res(res)

	assert.Equal(t, "HTTP/0.0 201 Created\r\nA: B\r\n\r\nHello", d)
}
