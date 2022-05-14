package body

import (
	"io"
	"net/http"

	"github.com/ofabricio/tie/opt"
)

func Copy(r io.Reader) opt.WriteFunc {
	return func(http.Header) opt.BodyFunc {
		return func(w io.Writer) error {
			_, err := io.Copy(w, r)
			return err
		}
	}
}

func CopyTo(w io.Writer) opt.ReadFunc {
	return func(r *http.Request) error {
		_, err := io.Copy(w, r.Body)
		return err
	}
}
