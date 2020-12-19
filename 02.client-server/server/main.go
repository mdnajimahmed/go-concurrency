package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	// infinitely waiting for connection
	fmt.Println("Waiting for client ...")
	for {
		// received a connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error conencting client", err)
		} else {
			go streamMovie(conn)
		}
	}
}

func streamMovie(conn net.Conn) {
	fmt.Println("Found a client")
	movieName := make([]byte, 1024)
	conn.Read(movieName)
	//broadcasting movie from my server
	for {
		n, e := io.WriteString(conn,
			fmt.Sprintf("bit stream 010101 of movie %s\n", movieName))
		if e != nil {
			fmt.Println("Error sending message", e)
		} else {
			fmt.Printf("Written %d bytes\n", n)
		}
		time.Sleep(1 * time.Second)
	}
}
