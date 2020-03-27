package main

import "fmt"

func main() {
	server := NewServer(":3000")
	fmt.Println("Server Running")
	server.Handle("/", HandleRoot)
	server.Handle("/api", HandleHome)
	server.Handle("/about", HandleAbout)
	server.Listen()
}
