package main

import (
	"bufio"
	"fmt"
	//"io/ioutil"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	nReader := bufio.NewReader(conn)
	nWriter := bufio.NewWriter(conn)
	for {
		line, err := nReader.ReadString('\n')
		if err == nil {
			fmt.Print(string(line))
		}
		nWriter.Write()
		nWriter.Flush()
	}
}
