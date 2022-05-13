package body

import (
	"io"

	"github.com/ofabricio/tie"
)

func Copy(r io.Reader) tie.WriteFunc {
	return func(c *tie.WriteConfig) {
		c.Body = func(w io.Writer) error {
			_, err := io.Copy(w, r)
			return err
		}
	}
}
