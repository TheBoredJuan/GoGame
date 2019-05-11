package main

import (
	"bufio"
	"fmt"
	//"io/ioutil"
	"net"
	"os"
	"os/exec"
)

func ClientWrite(nWriter *bufio.Writer) {
	for {
		var b []byte = make([]byte, 1)
		os.Stdin.Read(b)
		nWriter.WriteString(string(b))
		nWriter.Flush()
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	nReader := bufio.NewReader(conn)
	nWriter := bufio.NewWriter(conn)
	go ClientWrite(nWriter)
	for {
		line, err := nReader.ReadString('\n')
		if err == nil {
			fmt.Print(string(line))
		}
	}
}
