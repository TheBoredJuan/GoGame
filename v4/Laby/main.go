package main

import (
	"bufio"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"strings"
	//"io/ioutil"
	"net"
	"os"
	"os/exec"
)

const (
	ancho = 1080
	alto  = 680
)

var str string = ""
var Read chan string = make(chan string)

func ClientRead(nReader *bufio.Reader) {
	for {
		line, err := nReader.ReadString('\n')
		if err == nil {
			fmt.Print(string(line))
			Read <- line
		}
	}
}
func Decode() {
	for i := range Read {
		//fmt.Println("i:", i)
		chain := strings.Split(i, ":")
		if chain[0] == "MATRIX" {
			//fmt.Println("Prueba1", chain)
			str = chain[1]
		}
	}
}
func ClientWrite(nWriter *bufio.Writer) {
	for {
		var b []byte = make([]byte, 1)
		os.Stdin.Read(b)
		nWriter.WriteString(string(b))
		nWriter.Flush()
	}
}
func matriz(cad string, m int, n int) [][]string {
	mtr := make([][]string, m)
	cadena := strings.Split(cad, "*")
	for i := 0; i < m; i++ {
		cad2 := strings.Split(cadena[i], ",")
		mtr[i] = make([]string, n)
		for j := 0; j < n; j++ {
			mtr[i][j] = cad2[j]
		}
	}
	return mtr
}

var muros []muro
var infin []muro

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
	go ClientRead(nReader)
	go Decode()
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Iniciando SDL:", err)
		return
	}
	var m, n int
	m = 17
	n = 27

	//str="1,0,1,1,0,0,0,1,0,1,0,1,1,1,0,1,1,0,1*1,0,1,1,0,0,0,0,0,1,0,1,0,1,0,1,1,0,1*1,0,1,1,0,0,0,1,0,1,0,1,1,0,0,1,1,1,1*1,0,1,1,0,0,0,1,0,1,0,1,0,1,0,1,1,0,1*1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1*1,0,1,1,0,0,0,0,0,1,0,1,0,0,1,1,1,0,1*1,0,1,1,0,1,1,0,0,1,0,1,0,1,0,1,1,0,1*1,1,1,1,0,0,0,1,0,1,0,1,0,1,0,1,1,0,1*1,0,1,1,0,0,0,0,0,1,0,1,0,0,0,1,1,0,1*1,0,1,1,0,1,0,1,0,1,0,1,0,0,0,1,1,0,1*1,0,1,1,0,0,0,0,0,1,0,1,0,1,0,1,1,0,1*1,0,1,1,0,0,1,0,0,1,0,1,0,1,0,1,1,0,1*1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1"
	//str = "1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1*1,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,1,0,1,0,0,0,0,0,1*1,0,1,0,1,0,1,1,1,1,1,1,1,0,1,1,1,0,1,0,1,0,1,0,1,0,1*1,0,1,0,1,0,1,0,0,0,0,0,1,0,1,0,1,0,1,0,0,0,1,0,1,0,1*1,0,1,0,1,1,1,0,1,1,1,1,1,0,1,0,1,1,1,0,1,1,1,1,1,0,1*1,0,1,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,1,0,0,0,1*1,1,1,0,1,1,1,0,1,1,1,0,1,1,1,0,1,1,1,0,1,1,1,1,1,0,1*1,0,1,0,1,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,1,0,1*1,0,1,0,1,1,1,0,1,1,1,0,1,1,1,0,1,1,1,0,1,1,1,1,1,0,1*1,0,0,0,0,0,1,0,0,0,1,0,0,0,1,0,1,0,0,0,0,0,0,0,1,0,1*1,1,1,1,1,0,1,0,1,0,1,1,1,1,1,1,1,1,1,0,1,0,1,0,1,1,1*1,0,0,0,0,0,1,0,1,0,0,0,0,0,1,0,0,0,1,0,1,0,1,0,1,0,1*1,0,1,0,1,1,1,1,1,1,1,0,1,1,1,0,1,0,1,0,1,0,1,0,1,0,1*1,0,1,0,0,0,1,0,0,0,0,0,0,0,0,0,1,0,1,0,1,0,1,0,0,0,1*1,0,1,0,1,1,1,1,1,0,1,1,1,0,1,0,1,1,1,1,1,1,1,1,1,0,1*1,0,1,0,1,0,0,0,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,1,0,1*1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1*"
	for str == "" {
		//fmt.Println("str=", str)
	}
	tablero := matriz(str, m, n)
	window, err := sdl.CreateWindow(
		"Laby",
		0, 0,
		ancho, alto,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Iniciando SDL:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Iniciando SDL:", err)
		return
	}
	defer renderer.Destroy()

	var incx, incy float64

	jug, err := nuevoJugador(renderer, 45, 45)
	if err != nil {
		fmt.Println("Creando jugador:", err)
		return
	}

	for i := 0; i < m; i++ {
		incx = 0
		for j := 0; j < n; j++ {
			if i == 1 && j == 1 {
				x := incx
				y := incy
				ob, err := nuevoMuro(renderer, x, y, "sprites/inicio.bmp")
				if err != nil {
					fmt.Println("Creando muro: ", err)
				}
				infin = append(infin, ob)
			} else if tablero[i][j] == "1" {
				x := incx
				y := incy
				ob, err := nuevoMuro(renderer, x, y, "sprites/muro.bmp")
				if err != nil {
					fmt.Println("Creando muro: ", err)
					return
				}
				muros = append(muros, ob)
			} else if i == m-2 && j == n-2 {
				x := incx
				y := incy
				ob, err := nuevoMuro(renderer, x, y, "sprites/final.bmp")
				if err != nil {
					fmt.Println("Creando muro: ", err)
				}
				infin = append(infin, ob)
			}
			incx += TAMMURO
		}
		incy += TAMMURO
	}

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(129, 228, 198, 1)
		renderer.Clear()

		for _, mr := range muros {
			mr.dibujar(renderer)
		}

		for _, inf := range infin {
			inf.dibujar(renderer)
		}

		jug.actualizar()
		jug.dibujar(renderer)

		renderer.Present()
	}
}
