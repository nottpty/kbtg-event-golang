package main

import "html/template"

var eventTemplate = template.Must(template.New("event").Parse(`<!doctype html>
<html>
  <head>
    <meta charset="UTF-8">
		<title>{{.Name}}</title>
		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
  </head>
  <body>
    <div><a href="/events/{{.ID}}/edit">Edit</a></div>
    <h1>{{.Name}}</h1>
		<p>
    Amount of attendee : {{.AmountAttendee}}
		</p>
		<form method="POST" action="/events/{{.ID}}/export">
			<input class="btn btn-default" type="submit" value="Export data"/>
		</form>
		<form method="POST" action="/events/{{.ID}}/join">
			<h3>User ID</h3>
			<input name="userid" />
			<h3>First name</h3>
			<input name="firstname" />
			<h3>Last name</h3>
			<input name="lastname" />
			<h3>Phone number</h3>
			<input name="phonenumber" />
    	<input class="btn btn-default" type="submit" value="Join"/>
  	</form>
	</body>
</html>
`))
