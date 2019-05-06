package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "os/exec"
)

var allClients map[*Client]int

type Client struct {
    outgoing chan string
    reader   *bufio.Reader
    writer   *bufio.Writer
    conn     net.Conn
}

func (client *Client) Read() {
    /*for {
          var b []byte = make([]byte, 1)
          _, err := os.Stdin.Read(b)
          if err == nil {
              //fmt.Println("I got the byte", b, "("+string(b)+")")
              client.outgoing <- "Quiceno Pto\n"
          } else {
              break
          }
      }
      fmt.Println("Cliente desconectado")
      client.conn.Close()
      delete(allClients, client)
      client = nil*/
}

func (client *Client) Write() {
    for data := range client.outgoing {
        client.writer.WriteString(data)
        client.writer.Flush()
    }
}

func (client *Client) Listen() {
    go client.Read()
    go client.Write()
}

func NewClient(connection net.Conn) *Client {
    writer := bufio.NewWriter(connection)
    reader := bufio.NewReader(connection)

    client := &Client{
        // incoming: make(chan string),
        outgoing: make(chan string),
        conn:     connection,
        reader:   reader,
        writer:   writer,
    }
    client.Listen()

    return client
}
func ServerWrite() {
    for {
        var b []byte = make([]byte, 1)
        os.Stdin.Read(b)
        for client, i := range allClients {
            client.outgoing <- "Prueba1\n"
            fmt.Println(i)
        }
    }
}

func main() {
    exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    allClients = make(map[*Client]int)
    var i int = 1
    listener, _ := net.Listen("tcp", ":8080")
    go ServerWrite()
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println(err.Error())
        }
        //nReader := bufio.NewReader(conn)
        client := NewClient(conn)
        allClients[client] = i
        i++
        //go handleConection(nReader,allClients[client])
    }
}
