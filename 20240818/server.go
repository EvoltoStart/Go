package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type message struct {
	Createdate time.Time `json:"createdate"`
	SendUser   string    `json:"sendUser"`
	AcceptUser string    `json:"acceptUser"`
	Message    string    `json:"message"`
}

func mysql(messag *message) {
	dsn := "root:root@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True"

	// 打开数据库连接（初始化连接池）
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error on opening database connection: ", err)
	}
	defer db.Close()
	//rows, err := db.Query("select * from test1")
	res, err := db.Exec("insert into message(createdate,sendUser,acceptUser,message) values (?,?,?,?)", messag.Createdate, messag.SendUser, messag.AcceptUser, messag.Message)
	// 获取新插入记录的ID
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted record with ID: %d\n", lastInsertID)
}

func listmessage() []message {
	dsn := "root:root@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True"

	// 打开数据库连接（初始化连接池）
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error on opening database connection: ", err)
	}
	defer db.Close()
	rows, err := db.Query("select * from message")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	list := []message{}
	for rows.Next() {
		var row message
		if err := rows.Scan(&row.Createdate, &row.SendUser, &row.AcceptUser, &row.Message); err != nil {
			log.Fatal(err)
		}
		list = append(list, row)
	}
	return list
}

var user = make(map[string]net.Conn)

func MyTcpServer() {
	port := ":8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	checkServer(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkServer(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		UserExtis(conn)
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	sendMessage := message{}
	var buf [512]byte
	go func() {
		for {
			message, err := conn.Read(buf[0:])
			if err != nil {
				return
			}
			s := string(buf[0:message])
			fmt.Println("s:", s)
			if s == "search" {
				ss := make([]string, 10)
				for k, _ := range user {
					ss = append(ss, k+",")
				}
				conn.Write([]byte(strings.Join(ss, "")))
			} else if s == "fm" {
				list := listmessage()
				messagelist, _ := json.Marshal(list)
				fmt.Println(string(messagelist))
				conn.Write(messagelist)
			} else {
				ss := strings.Split(s, ":")
				fmt.Println("ss:", ss)
				n := ss[0]
				q := ss[1]

				fmt.Println("q:", q)
				qq := q[0:1]
				fmt.Println("qq", qq)
				if qq[0] == '@' {
					ns := q[1:]

					fmt.Println("ns:", ns)
					nss := strings.Split(ns, ",")
					fmt.Println("nss：", nss)
					sendMessage.Createdate = time.Now()
					sendMessage.SendUser = n
					sendMessage.AcceptUser = nss[0]
					sendMessage.Message = nss[1]
					user[nss[0]].Write([]byte(nss[0] + ":" + nss[1]))
					mysql(&sendMessage)
				} else {
					for k, v := range user {
						if k != n {
							//fmt.Print(string(buf[0:message]))
							v.Write(buf[0:message])
						}

					}
				}

			}

		}
	}()
}

func Write(conn net.Conn) {
	for {
		conn.Write([]byte("hello"))
	}

}
func UserExtis(conn net.Conn) {
	var buf [512]byte
	var is bool = false
	n, err := conn.Read(buf[0:])
	if err != nil {
		return
	}
	names := string(buf[0:n])
	for k, _ := range user {
		if k == names {
			is = true
			break
		}
	}
	if !is {
		user[names] = conn
	}
	conn.Write([]byte("欢迎您，" + names))
}
func checkServer(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
func main() {
	MyTcpServer()
}
