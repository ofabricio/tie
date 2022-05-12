# Tie

This is enough: a simple http handler utility.

[![Go](https://github.com/ofabricio/tie/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/ofabricio/tie/actions/workflows/go.yml)

## Example

```go
func HandlerExample(w http.ResponseWriter, r *http.Request) {

    var payload struct {
        Name string `json:"name"`
    }

    u := tie.New(w, r)

    u.Bind(&payload)

    u.Write(http.StatusOk, tie.Json(&payload))
}
```

See more [examples](/tie_test.go).
