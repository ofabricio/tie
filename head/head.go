package head

import "github.com/ofabricio/tie"

func With(kv ...string) tie.WriteFunc {
	return func(c *tie.WriteConfig) {
		for i := 0; i < len(kv); i += 2 {
			k, v := kv[i+0], kv[i+1]
			c.Head.Set(k, v)
		}
	}
}
