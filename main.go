package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	"gopkg.in/natefinch/npipe.v2"
)

const NAMED_PIPE = "\\\\.\\pipe\\HelloEvent"

func handleConnection(conn net.Conn) {
	r := bufio.NewReader(conn)
	msg, err := r.ReadString('\n')
	if err != nil {
		fmt.Println("received msg failed:", err.Error())
		return
	}
	fmt.Println("received msg:", msg)
}

func main() {
	fmt.Println("hello Event")

	optNum := len(os.Args)
	if 3 > optNum {
		fmt.Println("Server")

		go func() {
			ln, err := npipe.Listen(NAMED_PIPE)
			if err != nil {
				fmt.Println("Listen named_pipe failed:", err.Error())
				os.Exit(-1)
			}

			for {
				conn, err := ln.Accept()
				if err != nil {
					fmt.Println("Accept named_pipe failed:", err.Error())
					continue
				}
				go handleConnection(conn)
			}
		}()
	} else {
		fmt.Println("Client")

		//option := os.Args[1]
		msg := os.Args[2]

		go func(msg string) {
			conn, err := npipe.Dial(NAMED_PIPE)
			if err != nil {
				fmt.Println("Dial named_pipe failed:", err.Error())
			}
			if _, err := fmt.Fprintln(conn, msg); err != nil {
				fmt.Println("write to named_pipe failed:", err.Error())
			}
			/*r := bufio.NewReader(conn)
			  msg, err := r.ReadString('\n')
			  if err != nil {
			      // handle eror
			  }
			  fmt.Println(msg)*/
		}(msg)
	}
	time.Sleep(100 * time.Second)
}
