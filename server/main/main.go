package main

import (
	"fmt"
	"io"
	errs "letschat/error"
	"net"
)

var conns = make([]*net.Conn, 0)
var nicks = make(map[*net.Conn]string)

func main() {

	listen, err := net.Listen("tcp", ":8989")
	errs.PanicErr(err)

	for {
		// 等待接收连接
		accept, err := listen.Accept()
		for _, conn := range conns {
			_, _ = (*conn).Write([]byte(fmt.Sprintf("欢迎%s进入聊天室", accept.RemoteAddr().String())))
		}

		conns = append(conns, &accept)
		errs.PanicErr(err)
		go func() {
			dealClient(&accept)
		}()

	}

}

func dealClient(conn *net.Conn) {
	addr := (*conn).RemoteAddr()
	fmt.Println(fmt.Sprintf("%s连接进来了", addr))
	_, err := (*conn).Write([]byte("欢迎来到本聊天室,第一个输入的文本将作为你的昵称"))
	if err != nil {
		removeConn(conn)
		return
	}

	// 有客户端连接
	for {
		// max size
		tmp := make([]byte, 1024)
		n, err := (*conn).Read(tmp)
		if err != nil && err != io.EOF {
			fmt.Println("read error:", err)
			removeConn(conn)
			break
		} else if err == io.EOF {
			fmt.Println(fmt.Sprintf("%s退出了", addr))
			removeConn(conn)
			break
		}

		waitCloseConn := make([]*net.Conn, 0)

		if string(tmp[:n]) == "\n" {
			continue
		}

		if _, ok := nicks[conn]; ok == false {
			nicks[conn] = string(tmp[:n-1])
		}

		fmt.Print(nicks[conn] + "说:" + string(tmp[:n]))
		for i, conn2 := range conns {
			someBody := nicks[conn]
			if (*conn2).RemoteAddr().String() == addr.String() {
				someBody = "你"
			}
			_, err := (*conn2).Write([]byte(someBody + "说: " + string(tmp[:n])))
			if err != nil {
				fmt.Println(addr.String() + "出现了异常, 关闭此连接")
				waitCloseConn = append(waitCloseConn, conns[i])
				continue
			}
		}

		for _, temp := range waitCloseConn {
			removeConn(temp)
		}

	}
}

func removeConn(conn *net.Conn) {
	newConns := make([]*net.Conn, 0)
	for index := range conns {
		if conns[index] != conn {
			newConns = append(newConns, conns[index])
		}
	}

	conns = newConns
}
