# mustache

Mustache template renderer for [Forgery2](https://github.com/goforgery/forgery2).

## Install

    go get github.com/goforgery/mustache

## Use

By default template files are found in a folder named `./views` in the directory where __Forgery2__ is started.

```javascript
package main

import(
    "github.com/goforgery/forgery2"
    "github.com/goforgery/mustache"
)

func main() {
    app := f.CreateApp()
    app.Engine(".html", mustache.Create())
    app.Get("/", func(req *f.Request, res *f.Response, next func()) {
        res.Render("index.html", map[string]string{"title": "Mustache"})
    })
    app.Listen(3000)
}
```

The above code will look for the file `./views/index.html`.

## Test

    go test
