package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"strings"
)
const (
	ancho = 1080
	alto = 680
)

func matriz(cad string, m int, n int) ([][]string){
	mtr := make([][]string,m)
	cadena := strings.Split(cad,"*")
	for i:=0; i < m; i++{
		cad2 := strings.Split(cadena[i],",")
		mtr[i] = make([]string,n)
		for j:=0; j < n; j++{
			mtr[i][j]=cad2[j]
		}
	}
	return mtr
}

var muros []muro
var infin []muro
func main(){
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Iniciando SDL:", err)
		return
	}
	var m,n int
	m = 17
	n = 27
	var str string
	//str="1,0,1,1,0,0,0,1,0,1,0,1,1,1,0,1,1,0,1*1,0,1,1,0,0,0,0,0,1,0,1,0,1,0,1,1,0,1*1,0,1,1,0,0,0,1,0,1,0,1,1,0,0,1,1,1,1*1,0,1,1,0,0,0,1,0,1,0,1,0,1,0,1,1,0,1*1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1*1,0,1,1,0,0,0,0,0,1,0,1,0,0,1,1,1,0,1*1,0,1,1,0,1,1,0,0,1,0,1,0,1,0,1,1,0,1*1,1,1,1,0,0,0,1,0,1,0,1,0,1,0,1,1,0,1*1,0,1,1,0,0,0,0,0,1,0,1,0,0,0,1,1,0,1*1,0,1,1,0,1,0,1,0,1,0,1,0,0,0,1,1,0,1*1,0,1,1,0,0,0,0,0,1,0,1,0,1,0,1,1,0,1*1,0,1,1,0,0,1,0,0,1,0,1,0,1,0,1,1,0,1*1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1"
	str ="1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1*1,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,1,0,1,0,0,0,0,0,1*1,0,1,0,1,0,1,1,1,1,1,1,1,0,1,1,1,0,1,0,1,0,1,0,1,0,1*1,0,1,0,1,0,1,0,0,0,0,0,1,0,1,0,1,0,1,0,0,0,1,0,1,0,1*1,0,1,0,1,1,1,0,1,1,1,1,1,0,1,0,1,1,1,0,1,1,1,1,1,0,1*1,0,1,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,1,0,0,0,1,0,0,0,1*1,1,1,0,1,1,1,0,1,1,1,0,1,1,1,0,1,1,1,0,1,1,1,1,1,0,1*1,0,1,0,1,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,1,0,1*1,0,1,0,1,1,1,0,1,1,1,0,1,1,1,0,1,1,1,0,1,1,1,1,1,0,1*1,0,0,0,0,0,1,0,0,0,1,0,0,0,1,0,1,0,0,0,0,0,0,0,1,0,1*1,1,1,1,1,0,1,0,1,0,1,1,1,1,1,1,1,1,1,0,1,0,1,0,1,1,1*1,0,0,0,0,0,1,0,1,0,0,0,0,0,1,0,0,0,1,0,1,0,1,0,1,0,1*1,0,1,0,1,1,1,1,1,1,1,0,1,1,1,0,1,0,1,0,1,0,1,0,1,0,1*1,0,1,0,0,0,1,0,0,0,0,0,0,0,0,0,1,0,1,0,1,0,1,0,0,0,1*1,0,1,0,1,1,1,1,1,0,1,1,1,0,1,0,1,1,1,1,1,1,1,1,1,0,1*1,0,1,0,1,0,0,0,0,0,0,0,1,0,1,0,0,0,0,0,0,0,0,0,1,0,1*1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1*"
	tablero := matriz(str,m,n)

	window, err := sdl.CreateWindow(
		"Laby",
		sdl.WINDOWPOS_UNDEFINED,sdl.WINDOWPOS_UNDEFINED,
		ancho,alto,
		sdl.WINDOW_OPENGL)
	if err != nil{
		fmt.Println("Iniciando SDL:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil{
		fmt.Println("Iniciando SDL:", err)
		return
	}
	defer renderer.Destroy()

	var incx,incy float64

	jug,err := nuevoJugador(renderer,45,45)
	if err != nil{
		fmt.Println("Creando jugador:", err)
		return
	}

	for i:=0; i < m; i++{
		incx=0
		for j:=0; j < n; j++{
			if i == 1 && j == 1{
				x := incx
				y := incy
				ob, err := nuevoMuro(renderer,x,y,"sprites/inicio.bmp")
				if err != nil{
					fmt.Println("Creando muro: ", err)
				}
				infin = append(infin,ob)
			} else if tablero[i][j] == "1"{
					x := incx
					y := incy
					ob,err := nuevoMuro(renderer, x, y, "sprites/muro.bmp")
					if err != nil{
						fmt.Println("Creando muro: ", err)
						return
					}
				muros = append(muros,ob)
			} else if i==m-2 && j==n-2{
				x := incx
				y := incy
				ob, err := nuevoMuro(renderer,x,y,"sprites/final.bmp")
				if err != nil{
					fmt.Println("Creando muro: ", err)
				}
				infin = append(infin, ob)
			}
			incx += TAMMURO
		}
		incy += TAMMURO
	}

	for{
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent(){
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(129, 228, 198, 1)
		renderer.Clear()

		for _,mr := range muros{
			mr.dibujar(renderer)
		}

		for _,inf := range infin{
			inf.dibujar(renderer)
		}

		jug.actualizar()
		jug.dibujar(renderer)

		renderer.Present()
	}
}