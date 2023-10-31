package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", ":8888")
	if err!=nil {
		fmt.Println(err.Error())
	}
	//os.Exit(1)
	conn.Write([]byte("hello"))

	bs := make([]byte, 1024)
	n, err := conn.Read(bs)
	if err!=nil {
		fmt.Println(err.Error())
	}
	r := string(bs[:n])
	fmt.Println(r)


	conn.Close()



}