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
		result, err := conn.Read(readByte)
		errs.PanicErr(err)

		decrypt, err := secret.Decrypt(readByte[:result])
		if err != nil {
			fmt.Println("解密失败", err)
		}
		fmt.Println(string(decrypt))
	}
}

func write(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		encrypt, err := secret.Encrypt([]byte(text))
		if err != nil {
			fmt.Println("加密失败", err)
		}
		_, err = conn.Write(encrypt)
		errs.PanicErr(err)
	}
}
