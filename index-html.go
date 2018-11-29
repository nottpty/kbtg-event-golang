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
			<div class="w-100 text-center">
				<form action="/events/">
					<input class="text-center" style="height: 28px; width: auto;" type="image" src="https://kasikornbank.com/SiteCollectionDocuments/about/img/logo/logo.png" alt="Submit">
				</form>
			</div>
		</nav>

		<div class="container">
			<div class="row">
				{{range .Events}}
					<div class="col-4" style="padding-top: 15px; padding-bottom: 15px;">
						<div class="card">
							<img class="card-img-top w-100" src="http://positioningmag.com/wp-content/uploads/2016/04/6_kbank.png" alt="Card image cap">
							<div class="card-body text-center">
								<h5 class="card-title">{{ .Name }}</h5>
								<p>Generation {{ .Generation }}</p>
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
