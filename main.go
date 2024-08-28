package main

import (
	s "forum/backend/server"
)

func main() {
	server := s.New() //create new server
	server.Start()    // start the server
}
