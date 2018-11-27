package main

import "html/template"

var indexTemplate = template.Must(template.New("index").Parse(`<!doctype html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>Blog</title>
  </head>
  <body>
  <h1>Blog</h1>
  <div><a href="/events/new">New Post</a></div>
  <h2>Events</h2>
  {{range .Events}}<div><a href="/events/{{.ID}}">{{ .Name }}</a></div>{{end}}
  </body>
</html>
`))
