package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// this will keep JSON data
type JsonData struct {
	Args   []string
	Global map[string]string
	Vars   map[string]string
}

func HandleConnection(conn net.Conn, output string) {
	var jsonData JsonData
	defer conn.Close()

	// open file for writing JSON data
	file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error on open file: %s %s", err, output)
		return
	}
	writer := bufio.NewWriter(file)
	defer file.Close()

	// loop to receive JSON data: payload size first, then JSON
	for {
		// first receive JSON lenght in network order (big endian)
		buf := make([]byte, 2)
		_, err := conn.Read(buf)

		// connection closed
		if err == io.EOF {
			return
		} else if err != nil {
			log.Fatalf("Error on read: %s", err)
			return
		}

		// convert size to a little endian
		jsonSize := binary.BigEndian.Uint16(buf)

		// now receive whole JSON
		buf = make([]byte, jsonSize)
		_, err = conn.Read(buf)

		if err != nil {
			if err != io.EOF {
				log.Fatalf("Error on read: %s", err)
			}
			return
		} else {
			json.Unmarshal(buf, &jsonData)
			fmt.Fprintf(writer, "JSON: %#v\n", jsonData)
			writer.Flush()
		}
	}
}

