package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	errs "letschat/error"
	"net"
	"os"
	"time"
)

func main() {

	file, err := ioutil.ReadFile("conf/conf.txt")
	if err != nil {
		errs.PanicErr(err)
	}
	fmt.Println("连接主机:" + string(file))

	conn, err := net.Dial("tcp", string(file)+":8989")
	errs.PanicErr(err)

	go func() {
		read(conn)
	}()

	go func() {
		write(conn)
	}()

	time.Sleep(time.Duration(99999) * time.Hour)

}

func read(conn net.Conn) {
	for {
		readByte := make([]byte, 256)
		read, err := conn.Read(readByte)
		errs.PanicErr(err)
		if read <= 0 {
			continue
		}
		fmt.Println(string(readByte))
	}
}

func write(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		if text == "exit" {
			_, err := conn.Write([]byte("系统提示:对方已经下线"))
			errs.PanicErr(err)
			err = conn.Close()
			errs.PanicErr(err)
			return
		}

		_, err := conn.Write([]byte(text))
		errs.PanicErr(err)
	}
}
