package main

import "html/template"

var editTemplate = template.Must(template.New("edit").Parse(`<!doctype html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>Edit Event</title>
  </head>
  <h2>Edit Event</h2>
  <form method="POST" action="/events/{{.ID}}/update">
    <h3>Name</h3>
		<input name="name" value="{{.Name}}"/>
		<h3>Location</h3>
		<input name="location" value="{{.Location}}"/>
		<h3>Generation</h3>
    <input name="generation" value="{{.Generation}}"/>
    <h3>Description</h3>
		<textarea name="description">{{.Description}}</textarea></br>
		<h3>Speaker</h3>
		<input name="speaker" value="{{.Speaker}}"/>
		<h3>Limit Attendee</h3>
		<input name="limit" value="{{.LimitAttendee}}"/>
		<h3>Start Datetime</h3>
		<input name="start" value="{{.StartDatetime}}"/>
		<h3>End Datetime</h3>
    <input name="end" value="{{.EndDatetime}}"/>
    <input type="submit" />
  </form>
</html>
`))
