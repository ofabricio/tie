# Tie

This is enough: a simple http handler utility.

[![Go](https://github.com/ofabricio/tie/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/ofabricio/tie/actions/workflows/go.yml)

## Example

```go
import "github.com/ofabricio/tie"
import "github.com/ofabricio/tie/json"

func HandlerExample(w http.ResponseWriter, r *http.Request) {

    var payload struct {
        Name string `json:"name"`
    }

    u := tie.New(w, r)

    u.Bind(&payload)

    u.Write(http.StatusOK, json.Body(&payload))
}
```

## API

Even though the main module `tie` is the simplest way to handle requests,
keep in mind that the `u` (utility) module has a few extra features.

U

- [u.Read](#uRead)
- [u.Write](#uWrite)
- [u.Query](#uQuery)
- [u.Head](#uHead)
- [u.PathSeg](#uPathSeg)

Json

- [json.Body](#jsonBody)
- [json.With](#jsonWith)
- [json.Msg](#jsonMsg)
- [json.Bind](#jsonBind)
- [json.BindND](#jsonBindND)
- [json.ND](#jsonND)

Header

- [head.With](#headWith)

Body

- [body.Copy](#bodyCopy)
- [body.CopyTo](#bodyCopyTo)

### u.Read

Read reads the request content.

```go
import "github.com/ofabricio/tie/u"
import "github.com/ofabricio/tie/json"

func Example(w http.ResponseWriter, r *http.Request) {

    var payload struct {
        Name string `json:"name"`
    }

    u.Read(r, json.Bind(&payload))
}
```

### u.Write

Write writes the response content.

```go
import "github.com/ofabricio/tie/u"
import "github.com/ofabricio/tie/json"
import "github.com/ofabricio/tie/head"

func Example(w http.ResponseWriter, r *http.Request) {

    u.Write(w, http.StatusOK, json.Msg("Hello"), head.With("X-Version", "1"))
}
```

### u.Query

Query reads the URL query string content.

```go
import "github.com/ofabricio/tie/u"
import "github.com/ofabricio/tie/json"

func Example(w http.ResponseWriter, r *http.Request) {

    q := u.Query(r)

    p := Payload {
        ID:   q.Required("id").Int(),
        Name: q.Str("name"),
        Port: q.Int("port"),
        List: q.Split("list").Values(),
        Date: q.Time("date", time.Kitchen),
    }

    if q.Err != nil {
        u.Write(w, http.StatusBadRequest, json.Msg(q.Err.Error()))
    }
}
```

### u.Head

Head reads the request headers.

```go
import "github.com/ofabricio/tie/u"
import "github.com/ofabricio/tie/json"

func Example(w http.ResponseWriter, r *http.Request) {

    h := u.Head(r)

    p := Payload {
        Cntt: h.Required("Content-Type").Str(),
        Data: h.Int("X-Data"),
        Date: h.Time("X-Date", time.Kitchen),
    }

    if h.Err != nil {
        u.Write(w, http.StatusBadRequest, json.Msg(h.Err.Error()))
    }
}
```

### u.PathSeg

PathSeg reads the request URL path segments.

```go
import "github.com/ofabricio/tie/u"
import "github.com/ofabricio/tie/json"

func Example(w http.ResponseWriter, r *http.Request) {

    // Suppose Path = "/one/two/3"

    a := u.PathSeg(r, 0).Str()      // one
    b := u.PathSeg(r, 1).Str()      // two
    c, err := u.PathSeg(r, 2).Int() // 3

    if err != nil {
        u.Write(w, http.StatusBadRequest, json.Msg(err.Error()))
    }
}
```

### json.Body

Body writes a response in Json format given a struct.

```go
import "github.com/ofabricio/tie/u"
import "github.com/ofabricio/tie/json"

func Example(w http.ResponseWriter, r *http.Request) {

    d := Data{Name: "Mary"}

    u.Write(w, http.StatusOK, json.Body(&d))
}
```

### json.With

With writes a response in Json format given a key value map.

```go
import "github.com/ofabricio/tie/u"
import "github.com/ofabricio/tie/json"

func Example(w http.ResponseWriter, r *http.Request) {

    u.Write(w, http.StatusOK, json.With("name", "Mary"))
}
```

### json.Msg

Msg writes a response in Json format given a key value map with a `message` field: `{ "message": "hello" }`

```go
import "github.com/ofabricio/tie/u"
import "github.com/ofabricio/tie/json"

func Example(w http.ResponseWriter, r *http.Request) {

    u.Write(w, http.StatusOK, json.Msg("hello"))
}
```

### json.Bind

Bind reads a Json from the request body and binds it to a struct.

```go
import "github.com/ofabricio/tie/u"
import "github.com/ofabricio/tie/json"

func Example(w http.ResponseWriter, r *http.Request) {

    var payload struct {
        Name string `json:"name"`
    }

    u.Read(r, json.Bind(&payload))
}
```

### json.BindND

BindND reads a NDJson from the request body and binds it to a struct.

```go
import "github.com/ofabricio/tie/u"
import "github.com/ofabricio/tie/json"

func Example(w http.ResponseWriter, r *http.Request) {

    ch := make(chan Data)

    go u.Read(req, json.BindND(ch))

    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

### json.ND

ND writes a response in NDJson format.

```go
import "github.com/ofabricio/tie/u"
import "github.com/ofabricio/tie/json"

func Example(w http.ResponseWriter, r *http.Request) {

    ch := make(chan Data)

    go func() {
        ch <- Data{Name: "Mary"}
        ch <- Data{Name: "John"}
        close(ch)
    }()

    u.Write(w, http.StatusOK, json.ND(ch))
}
```

### head.With

With writes response headers.

```go
import "github.com/ofabricio/tie/u"
import "github.com/ofabricio/tie/head"

func Example(w http.ResponseWriter, r *http.Request) {

    u.Write(w, http.StatusOK, head.With("Content-Type", "text/plain"))
}
```

### body.Copy

Copy copies a Reader to the response body.

```go
import "github.com/ofabricio/tie/u"
import "github.com/ofabricio/tie/body"

func Example(w http.ResponseWriter, r *http.Request) {

    s := strings.NewReader("Hello")

    u.Write(w, http.StatusOK, body.Copy(s))
}
```

### body.CopyTo

CopyTo copies the request body to a Writer.

```go
import "github.com/ofabricio/tie/u"
import "github.com/ofabricio/tie/body"

func Example(w http.ResponseWriter, r *http.Request) {

    var buf bytes.Buffer

    u.Read(w, body.CopyTo(&buf))
}
```
