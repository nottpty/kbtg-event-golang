package main

import "html/template"

var newEventTemplate = template.Must(template.New("new-event").Parse(`<!doctype html>
<html>
  <head>
    <meta charset="UTF-8">
		<title>New Event</title>
		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
  </head>
  <body>
    <h2>New Event</h2>
    <form method="POST" action="/events/">
      <h3>Name event</h3>
      <input name="name" />
      <h3>Description</h3>
			<textarea name="description"></textarea></br>
			<h3>Location</h3>
			<textarea name="location"></textarea></br>
			<h3>Generation</h3>
			<input name="generation" />
			<h3>Speaker</h3>
			<input name="speaker" />
			<h3>Limit Attendee</h3>
			<input name="limit" />
			<h3>Start Datetime</h3>
			<input value="Jan 2, 2006 at 3:04 PM" name="start" />
			<h3>End Datetime</h3>
			<input value="Jan 2, 2006 at 3:04 PM" name="end" />
      <input type="submit" />
		</form>
  </body>
</html>
`))
