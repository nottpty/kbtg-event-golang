package main

import "html/template"

var indexTemplate = template.Must(template.New("index").Parse(`<!doctype html>
<html>
  <head>
    <meta charset="UTF-8">
		<title>KBTG Event</title>
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css">
  </head>
	<body>
		<nav class="navbar navbar-dark bg-success">
			<a class="navbar-brand" href="/events/">HOME</a>
		</nav>
		<div class="text-center" style="padding-top: 15px; padding-bottom: 15px;">
			<h1>Events</h1>
		</div>
		<div class="container">
			<div class="row">
				{{range .Events}}
					<div class="col-4" style="padding-bottom: 15px;">
							<div class="card border border-success">
								<h5 class="card-header text-center" style="background-color: #84d57f;">{{ .Name }}</h5>
								<div class="card-body">
									<h5 class="card-title">By {{ .Speaker }}</h5>
									<p class="card-text">{{ .Description }}</p>
									<a href="/events/{{.ID}}" class="btn btn-info" style="background-color: #1e7325;">Go to detail</a>
								</div>
							</div>
					</div>
				{{end}}
			</div>
		</div>
  </body>
</html>
`))
