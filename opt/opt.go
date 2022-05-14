package opt

import (
	"io"
	"net/http"
)

type WriteFunc func(http.Header) BodyFunc

type BodyFunc func(io.Writer) error

type ReadFunc func(*http.Request) error
