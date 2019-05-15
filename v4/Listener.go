package main

import (
    "bufio"
    "fmt"
    "math"
    "math/rand"
    "net"
    "os"
    "os/exec"
    "time"
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
    visited    int
    node_up    *Node
    node_down  *Node
    node_left  *Node
    node_right *Node
}

type NodexDir struct {
    node *Node
    dir  string
}

func NewNode() *Node {
    node := &Node{
        visited:    0,
        node_up:    nil,
        node_down:  nil,
        node_left:  nil,
        node_right: nil,
    }
    return node
}
func (nodexdir *NodexDir) getNode() *Node {
    return nodexdir.node
}
func (nodexdir *NodexDir) getDir() string {
    return nodexdir.dir
}
func NewNodexDir(node *Node, dir string) *NodexDir {
    nodexdir := &NodexDir{
        node: node,
        dir:  dir,
    }
    return nodexdir
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

func (node *Node) visit() {
    node.visited = 1
}

func (node *Node) isVisited() int {
    return node.visited
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

func (node *Node) getNeighborsNV() []*NodexDir {
    returnList := []*NodexDir{}
    if node.node_up != nil {
        if node.node_up.isVisited() == 0 {
            returnList = append(returnList, NewNodexDir(node, "up"))
        }
    }
    if node.node_down != nil {
        if node.node_down.isVisited() == 0 {
            returnList = append(returnList, NewNodexDir(node, "down"))
        }
    }
    if node.node_left != nil {
        if node.node_left.isVisited() == 0 {
            returnList = append(returnList, NewNodexDir(node, "left"))
        }
    }
    if node.node_right != nil {
        if node.node_right.isVisited() == 0 {
            returnList = append(returnList, NewNodexDir(node, "right"))
        }
    }
    return returnList

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
    var actualNode *Node = labyrinth[0][0]
    actualNode.visit()
    var listNeighbors []*NodexDir = actualNode.getNeighborsNV()
    fmt.Println(listNeighbors)
    for len(listNeighbors) > 0 {
        rand.Seed(time.Now().UnixNano())
        i := rand.Intn(len(listNeighbors))
        actualNodexDir := listNeighbors[i]
        actualNode := actualNodexDir.getNode()
        actualNode.visit()
        //fmt.Println("Dir :", actualNodexDir.getDir())
        //fmt.Println(actualNode.node_up, actualNode.node_down, actualNode.node_left, actualNode.node_right, "\n")
        listNeighbors = append(listNeighbors[:i], listNeighbors[i+1:]...)
        if actualNodexDir.getDir() == "up" {
            if actualNode.node_up.isVisited() == 0 {
                actualNode.node_up.visit()
                actualNode.node_up.node_down = nil
                listNeighbors = append(listNeighbors, actualNode.node_up.getNeighborsNV()...)
                actualNode.node_up = nil
            }
        } else if actualNodexDir.getDir() == "down" {
            if actualNode.node_down.isVisited() == 0 {
                actualNode.node_down.visit()
                actualNode.node_down.node_up = nil
                listNeighbors = append(listNeighbors, actualNode.node_down.getNeighborsNV()...)
                actualNode.node_down = nil
            }
        } else if actualNodexDir.getDir() == "left" {
            if actualNode.node_left.isVisited() == 0 {
                actualNode.node_left.visit()
                actualNode.node_left.node_right = nil
                listNeighbors = append(listNeighbors, actualNode.node_left.getNeighborsNV()...)
                actualNode.node_left = nil
            }
        } else if actualNodexDir.getDir() == "right" {
            if actualNode.node_right.isVisited() == 0 {
                actualNode.node_right.visit()
                actualNode.node_right.node_left = nil
                listNeighbors = append(listNeighbors, actualNode.node_right.getNeighborsNV()...)
                actualNode.node_right = nil
            }
        }
    }
    return labyrinth
}

func createLabyRaw() [N][M]*Node {
    return createMaze(createConectionsOfLaby(createMatrixOfLaby()))
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
                finalString = finalString + "1*"
            } else if math.Mod(float64(j), 2) == 1 && math.Mod(float64(i), 2) == 1 {
                finalString = finalString + "1,"
            } else if math.Mod(float64(j), 2) == 0 && math.Mod(float64(i), 2) == 0 {
                finalString = finalString + "0,"
            } else {
                if math.Mod(float64(j), 2) == 1 && math.Mod(float64(i), 2) == 0 {
                    if labyrinth[(i/2)-1][((j-1)/2)-1].getNodeRight() != nil {
                        finalString = finalString + "1,"
                    } else {
                        finalString = finalString + "0,"
                    }
                } else if math.Mod(float64(j), 2) == 0 && math.Mod(float64(i), 2) == 1 {
                    if labyrinth[((i-1)/2)-1][(j/2)-1].getNodeDown() != nil {
                        finalString = finalString + "1,"
                    } else {
                        finalString = finalString + "0,"
                    }
                }

            }
        }
    }
    return finalString
}

func (client *Client) Write() {
    for data := range client.outgoing {
        _, err := client.writer.WriteString(data)
        client.writer.Flush()
        if err != nil {
            fmt.Println("The Client: ", allClients[client], " has been logout")
            delete(allClients, client)
        }
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
func ServerWrite(s string) {
    for {
        var b []byte = make([]byte, 1)
        os.Stdin.Read(b)
        for client, i := range allClients {
            client.outgoing <- s + ":"
            fmt.Println(i)
        }
    }
}
func ServerWriteAMsg(s string) {
    for client, i := range allClients {
        client.outgoing <- s + ":"
        fmt.Println("Menssage Sent to Client:", i)
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
    var Laby [N][M]*Node = createLabyRaw()
    var s string = "MATRIX:" + translationOfLaby(Laby) + ":\n"
    fmt.Println(s)
    go ServerWriteAMsg(s)
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println(err.Error())
        }
        go ServerWriteAMsg(s)
        nReader := bufio.NewReader(conn)
        go ServerRead(nReader)
        client := NewClient(conn)
        allClients[client] = i
        i++
        //go handleConection(nReader,allClients[client])
    }
}
