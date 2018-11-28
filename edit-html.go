package main

import "html/template"

var editTemplate = template.Must(template.New("edit").Parse(`<!doctype html>
<html>
  <head>
    <meta charset="UTF-8">
		<title>Edit Event</title>
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css">
	</head>
	<body>
		<nav class="navbar navbar-dark bg-success">
			<a class="navbar-brand" href="/events/">HOME</a>
		</nav>
		<div class="text-center" style="padding-top: 15px; padding-bottom: 15px;">
			<h1>Edit Event</h1>
		</div>
		<form class="text-center" method="POST" action="/events/{{.ID}}/export">
			<input class="btn btn-outline-info" type="submit" value="Export data to Excel">
		</form>
		<div class="container" style="padding-top: 30px; padding-bottom: 15px;">
			<div class="row">
				<div class="col-3"></div>
				<div class="col-6">
					<div class="container">
						<form method="POST" action="/events/{{.ID}}/update">
							<div class="form-group">
								<label for="name">Name</label>
								<input class="form-control" name="name" id="name" value="{{.Name}}">
							</div>
							<div class="form-group">
								<label for="location">Location</label>
								<textarea class="form-control" name="location" id="location">{{.Location}}</textarea>
							</div>
							<div class="form-group">
								<label for="generation">Generation</label>
								<input class="form-control" name="generation" id="generation" value="{{.Generation}}">
							</div>
							<div class="form-group">
								<label for="description">Description</label>
								<textarea class="form-control" name="description" id="description">{{.Description}}</textarea>
							</div>
							<div class="form-group">
								<label for="speaker">Speaker</label>
								<input class="form-control" name="speaker" id="speaker" value="{{.Speaker}}">
							</div>
							<div class="form-group">
								<label for="limit">Limit Attendee</label>
								<input class="form-control" name="limit" id="limit" value="{{.LimitAttendee}}">
							</div>
							<div class="form-group">
								<label for="start">Start Datetime</label>
								<input class="form-control" name="start" id="start" value="{{.StartDatetime}}">
							</div>
							<div class="form-group">
								<label for="end">End Datetime</label>
								<input class="form-control" name="end" id="end" value="{{.EndDatetime}}">
							</div>
							<div class="text-center"><input class="btn btn-outline-success" type="submit"></div>
						</form>
					</div>
				</div>
				<div class="col-3"></div>
			</div>
		</div>
	</body>
</html>
`))
