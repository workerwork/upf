package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/workerwork/upf/conf"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	conf.Setup()
	serverAddr := conf.ServerConf.IP + ":" + strconv.Itoa(conf.ServerConf.Port)
	fmt.Println(serverAddr)
	conn, err := net.Dial("udp", serverAddr)
	checkError(err)

	defer conn.Close()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()

		lineLen := len(line)

		n := 0
		for written := 0; written < lineLen; written += n {
			var toWrite string
			if lineLen-written > conf.ServerConf.RecvLen {
				toWrite = line[written : written+conf.ServerConf.RecvLen]
			} else {
				toWrite = line[written:]
			}

			n, err = conn.Write([]byte(toWrite))
			checkError(err)

			fmt.Println("Write:", toWrite)

			msg := make([]byte, conf.ServerConf.RecvLen)
			n, err = conn.Read(msg)
			checkError(err)

			fmt.Println("Response:", string(msg))
		}
	}
}
