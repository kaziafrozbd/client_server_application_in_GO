package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	
	nl, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Print("server started on :8888")

	for {
	conn, err := nl.Accept()
	if err != nil {
		fmt.Println(err.Error())
		// os.Exit(1)
	}
	// r := conn.RemoteAddr()
	// fmt.Println(r)

	bs := make([]byte, 1024)

	n, err := conn.Read(bs)
	if err != nil {
		fmt.Println(err.Error())
		// os.Exit(1)
	}
	//fmt.Println(n) //koto byte request ashse dekhar jonno
	req := string(bs[:n])
	fmt.Println(req)
	recTime := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf(`your massage is: %s, at the time of %s`,req ,recTime)
	conn.Write([]byte(msg))

// body := `<!DOCTYPE html><html><head><title>Page Title</title></head><body><h1>My First Heading</h1></body></html>`

// response := "HTTP/1.1 200 OK\r\n" + "Content-Length: %d\r\n" + "Content-Type: text/html\r\n" + "\r\n%s"

// msg := fmt.Sprintf(response, len(body), body)
// fmt.Println(msg)
// conn.Write([]byte(msg))
}

}
