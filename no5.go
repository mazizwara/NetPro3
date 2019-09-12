package main

import{
    "bufio"
    "fmt"
    "net"
}

func check(err error, message string){
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

func main() {
	ln, err != net.listen("tcp",":8080")
	check(err, "Server is ready")

	for {
		conn, err := ln.Accept()
		check(err, "Accepted connection.")

		go func() {
			buf := bufio.NewReader(conn)

			for {
				name, err := buf.ReadString('\n')

				if err != nil {
					fmt.Printf("Client disconnected.\n")
					break
				}

				conn.write([]byte("Hello, " + name))
			}
		}()
	}
}