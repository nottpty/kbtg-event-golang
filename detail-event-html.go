package main

import "html/template"

var detailEventTemplate = template.Must(template.New("detailEvent").Parse(`<!doctype html>
<html>
  <head>
    <meta charset="UTF-8">
		<title>{{.Name}}</title>
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css">
  </head>
  <body>
		<nav class="navbar navbar-dark bg-success">
			<div class="w-100 text-center">
				<form action="/events/">
					<input class="text-center" style="height: 28px; width: auto;" type="image" src="https://kasikornbank.com/SiteCollectionDocuments/about/img/logo/logo.png" alt="Submit">
				</form>
			</div>
		</nav>
		<div class="text-center" style="padding-top: 15px; padding-bottom: 15px;">
			<h1>{{.Name}}</h1>
		</div>
		
		<div class="container">
			<div class="row">
				{{range .Events}}
					<div class="col-4" style="padding-bottom: 15px;">
							<div class="card border border-success">
								<h5 class="card-header text-center" style="background-color: #84d57f;">{{ .Generation }}</h5>
								<div class="card-body">
									<h5 class="card-title">By {{ .Speaker }}</h5>
									<p class="card-text">{{ .Description }}</p>
									<a href="/events/{{.ID}}" class="btn btn-info" style="background-color: #1e7325;">Register</a>
								</div>
							</div>
					</div>
				{{end}}
			</div>
		</div>
	</body>
</html>
`))
