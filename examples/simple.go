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
