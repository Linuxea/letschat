package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	errs "letschat/error"
	"letschat/secret"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	file, err := ioutil.ReadFile("conf/conf.txt")
	if err != nil {
		errs.PanicErr(err)
	}
	host := strings.TrimRight(string(file), "\n")
	conn, err := net.Dial("tcp", host+":8989")
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
		_, err := conn.Read(readByte)
		errs.PanicErr(err)
		fmt.Println(secret.Decrypt(string(readByte)))
	}
}

func write(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		_, err := conn.Write([]byte(secret.Encrypt(text)))
		errs.PanicErr(err)
	}
}
