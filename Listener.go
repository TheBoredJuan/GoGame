package main

import (
    "bufio"
    "fmt"
    "math"
    "net"
    "os"
    "os/exec"
)

//https://en.wikipedia.org/wiki/Maze_generation_algorithm
const N = 8
const M = 13

var allClients map[*Client]int
var labyrinth [N][M]*Node

type Client struct {
    outgoing chan string
    // reader   *bufio.Reader
    writer *bufio.Writer
    conn   net.Conn
}

type Node struct {
    node_up    *Node
    node_down  *Node
    node_left  *Node
    node_right *Node
}

func NewNode() *Node {
    node := &Node{
        node_up:    nil,
        node_down:  nil,
        node_left:  nil,
        node_right: nil,
    }
    return node
}

func (node *Node) setNodeUp(conecNode *Node) {
    node.node_up = conecNode
}
func (node *Node) setNodeDown(conecNode *Node) {
    node.node_down = conecNode
}
func (node *Node) setNodeLeft(conecNode *Node) {
    node.node_left = conecNode
}
func (node *Node) setNodeRight(conecNode *Node) {
    node.node_right = conecNode
}

func (node *Node) getNodeUp() *Node {
    return node.node_up
}
func (node *Node) getNodeDown() *Node {
    return node.node_down
}
func (node *Node) getNodeLeft() *Node {
    return node.node_left
}
func (node *Node) getNodeRight() *Node {
    return node.node_right
}

func (node *Node) getNeighbors() [4]*Node {
    var neighborsList [4]*Node
    if node.node_up != nil {
        neighborsList[0] = (node.node_up)
    }
    if node.node_down != nil {
        neighborsList[1] = (node.node_down)
    }
    if node.node_left != nil {
        neighborsList[2] = (node.node_left)
    }
    if node.node_right != nil {
        neighborsList[3] = (node.node_right)
    }
    return neighborsList

}

func createMatrixOfLaby() [N][M]*Node {
    fmt.Println(N)
    for i := 1; i <= N; i++ {
        for j := 1; j <= M; j++ {
            labyrinth[i-1][j-1] = NewNode()
            //fmt.Print(1)
        }
        fmt.Print("\n")
    }
    return labyrinth
}

func createConectionsOfLaby(labyrinth [N][M]*Node) [N][M]*Node {
    for i := 1; i <= N; i++ {
        for j := 1; j <= M; j++ {
            if i > 1 {
                labyrinth[i-1][j-1].setNodeUp(labyrinth[i-2][j-1])
            }
            if i < N {
                labyrinth[i-1][j-1].setNodeDown(labyrinth[i][j-1])
            }
            if j > 1 {
                labyrinth[i-1][j-1].setNodeLeft(labyrinth[i-1][j-2])
            }
            if j < M {
                labyrinth[i-1][j-1].setNodeRight(labyrinth[i-1][j])
            }

        }
    }
    return labyrinth
}

func createMaze(labyrinth [N][M]*Node) [N][M]*Node {
    wallArray = labyrinth[0][0].getNeighbors()

}

func createLabyRaw() [N][M]*Node {
    return createConectionsOfLaby(createMatrixOfLaby())
}

func translationOfLaby(labyrinth [N][M]*Node) string {
    var tamañoI int = 3 + (2 * (N - 1))
    var tamañoJ int = 3 + (2 * (M - 1))
    fmt.Println(tamañoI)
    fmt.Println(tamañoJ)
    var finalString string = ""
    for i := 1; i <= tamañoI; i++ {
        for j := 1; j <= tamañoJ; j++ {
            if (i == 1 || i == tamañoI) && j < tamañoJ {
                finalString = finalString + "1,"
            } else if j == 1 {
                finalString = finalString + "1,"
            } else if j == tamañoJ {
                5
                finalString = finalString + "1*"
            } else if math.Mod(float64(j), 2) == 1 && math.Mod(float64(i), 2) == 1 {
                finalString = finalString + "1,"
            } else if math.Mod(float64(j), 2) == 0 && math.Mod(float64(i), 2) == 0 {
                finalString = finalString + "0,"
            } else {
                if math.Mod(float64(j), 2) == 1 && math.Mod(float64(i), 2) == 0 {
                    if labyrinth[(i/2)-1][((j-1)/2)-1].getNodeRight() != nil {
                        finalString = finalString + "0,"
                    } else {
                        finalString = finalString + "1,"
                    }
                } else if math.Mod(float64(j), 2) == 0 && math.Mod(float64(i), 2) == 1 {
                    if labyrinth[((i-1)/2)-1][(j/2)-1].getNodeDown() != nil {
                        finalString = finalString + "0,"
                    } else {
                        finalString = finalString + "1,"
                    }
                }

            }
        }
    }
    return finalString
}

func (client *Client) Write() {
    for data := range client.outgoing {
        client.writer.WriteString(data)
        client.writer.Flush()
    }
}

func (client *Client) Listen() {
    go client.Write()
}

func NewClient(connection net.Conn) *Client {
    writer := bufio.NewWriter(connection)
    //reader := bufio.NewReader(connection)

    client := &Client{
        // incoming: make(chan string),
        outgoing: make(chan string),
        conn:     connection,
        //reader:   reader,
        writer: writer,
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

func ServerRead(nReader *bufio.Reader) {
    for {
        line, err := nReader.ReadString('\n')
        if err == nil {
            //fmt.Print("Jelo")
            fmt.Print(string(line))
        }
    }
}

func main() {
    exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    allClients = make(map[*Client]int)
    var i int = 1
    listener, _ := net.Listen("tcp", ":8080")
    go ServerWrite()
    var Laby [N][M]*Node = createLabyRaw()
    var s string = translationOfLaby(Laby)
    fmt.Println(s)
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println(err.Error())
        }
        nReader := bufio.NewReader(conn)
        go ServerRead(nReader)
        client := NewClient(conn)
        allClients[client] = i
        i++
        //go handleConection(nReader,allClients[client])
    }
}
