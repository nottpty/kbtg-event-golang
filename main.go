package main

import (
	"log"

	// "web-event/pq/post"
	"web-event/pq/event"
)

func main() {
	event.ConnectDB()
	log.Fatal(startServer())
}
