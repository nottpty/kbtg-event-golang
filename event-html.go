package main

import "html/template"

var eventTemplate = template.Must(template.New("event").Parse(`<!doctype html>
<html>
  <head>
    <meta charset="UTF-8">
		<title>{{.Name}}</title>
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css">
  </head>
  <body>
		<nav class="navbar navbar-dark bg-success">
			<a class="navbar-brand" href="/events/">HOME</a>
		</nav>
		<div class="text-center" style="padding-top: 15px; padding-bottom: 15px;">
			<h1>{{.Name}}</h1>
		</div>
		<div class="text-center">
			<p>Amount of attendees : {{.AmountAttendee}}</p>
		</div>

		{{ if lt .AmountAttendee .LimitAttendee }}
		<div class="container" style="padding-top: 25px; padding-bottom: 15px;">
			<div class="row">
				<div class="col-3"></div>
				<div class="col-6">
					<div class="container">
						<form method="POST" action="/events/{{.ID}}/join">
							<div class="form-group">
								<label for="userid">User ID</label>
								<input type="text" class="form-control" name="userid" id="userid" placeholder="999888">
							</div>
							<div class="form-group">
								<label for="firstname">First name</label>
								<input type="text" class="form-control" name="firstname" id="firstname" placeholder="Korakade">
							</div>
							<div class="form-group">
								<label for="lastname">Last name</label>
								<input type="text" class="form-control" name="lastname" id="lastname" placeholder="Desaeek">
							</div>
							<div class="form-group">
								<label for="phonenumber">Phone number</label>
								<input type="text" class="form-control" name="phonenumber" id="phonenumber" placeholder="089-247-1567">
							</div>
							
							<div class="text-center"><input class="btn btn-outline-success" type="submit" value="Join event"></div>
							
						</form>
					</div>
				</div>
				<div class="col-3"></div>
			</div>
		</div>
		{{ end }}
	</body>
</html>
`))
