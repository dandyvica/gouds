package main

import (
	"fmt"
	"log"
	"net"
	"os"
)


func main() {
	// get cli args & flags
	options := GetArgs()

	// try to remove socket first to get rid of already in use errors
    if err := os.RemoveAll(options.domain); err != nil {
        log.Fatal(err)
    }

	// create socket
    sock, err := net.Listen("unix", options.domain)
    if err != nil {
        log.Fatal("listen error:", err)
    }
    defer sock.Close()

	// loop an client connection
    for {
        // Accept new connections
        fmt.Println("waiting for a client")
        conn, err := sock.Accept()
        if err != nil {
            log.Fatal("accept error:", err)
        }

		// call main function to manage receiving of data
        HandleConnection(conn, options.output)

        fmt.Println("end of client")
    }
}