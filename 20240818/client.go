package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

var input string
var name string
var wg = sync.WaitGroup{}

func MyTcpClient() {
	var buf [512]byte
	wg.Add(2)
	port := "localhost:8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	checkErrCli(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkErrCli(err)
	login(conn)
	defer conn.Close()
	go func() {
		defer wg.Done()
		for {
			message, err := conn.Read(buf[0:])
			if err != nil {
				return
			}
			fmt.Printf("\r%s\n%s:", string(buf[0:message]), name)
		}

	}()

	go func() {
		defer wg.Done()
		for input != "exit" {
			//time.Sleep(time.Second)
			fmt.Print(name + ":")
			fmt.Scanln(&input)
			if input == "search" {
				conn.Write([]byte(input))
			} else if input == "fm" {
				conn.Write([]byte(input))
			} else {
				s := name + ":" + input
				//fmt.Println("tesst:", s)
				_, err = conn.Write([]byte(s))
				//_, err := conn.Write([]byte("aa:@b fdsff"))
				if err != nil {
					return
				}
			}

		}
	}()
	wg.Wait()
}

func login(conn net.Conn) {

	var buf [512]byte
	fmt.Print("请输入用户名：")
	fmt.Scanln(&input)
	name = input
	conn.Write([]byte(name))
	message, err := conn.Read(buf[0:])
	if err != nil {
		return
	}
	fmt.Println(string(buf[0:message]))
}

func checkErrCli(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	MyTcpClient()
}
