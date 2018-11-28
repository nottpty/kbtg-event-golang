package main

import "html/template"

var newEventTemplate = template.Must(template.New("new-event").Parse(`<!doctype html>
<html>
  <head>
    <meta charset="UTF-8">
		<title>New Event</title>
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css">
  </head>
	<body>
		<nav class="navbar navbar-dark bg-success">
			<a class="navbar-brand" href="/events/">HOME</a>
		</nav>

		<div class="text-center" style="padding-top: 15px; padding-bottom: 15px;">
			<h1>Create Event</h1>
		</div>
		<div class="container" style="padding-top: 15px; padding-bottom: 15px;">
			<div class="row">
				<div class="col-3"></div>
				<div class="col-6">
					<div class="container">
						<form method="POST" action="/events/">
							<div class="form-group">
								<label for="name">Name</label>
								<input class="form-control" name="name" id="name" placeholder="Golang intensive course">
							</div>
							<div class="form-group">
								<label for="location">Location</label>
								<textarea class="form-control" name="location" id="location" placeholder="KBTG"></textarea>
							</div>
							<div class="form-group">
								<label for="generation">Generation</label>
								<input class="form-control" name="generation" id="generation" placeholder="1">
							</div>
							<div class="form-group">
								<label for="description">Description</label>
								<textarea class="form-control" name="description" id="description" placeholder="For beginner to expert"></textarea>
							</div>
							<div class="form-group">
								<label for="speaker">Speaker</label>
								<input class="form-control" name="speaker" id="speaker" placeholder="Mr.MasterX">
							</div>
							<div class="form-group">
								<label for="limit">Limit Attendee</label>
								<input class="form-control" name="limit" id="limit" placeholder="30">
							</div>
							<div class="form-group">
								<label for="start">Start Datetime</label>
								<input class="form-control" name="start" id="start" placeholder="Jan 2, 2006 at 3:04 PM">
							</div>
							<div class="form-group">
								<label for="end">End Datetime</label>
								<input class="form-control" name="end" id="end" placeholder="Jan 2, 2006 at 3:04 PM">
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
