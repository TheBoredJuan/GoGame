package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	TAMJUGADOR = 30
	VEL        = 2.5
)

type cuadrado struct {
	x, y, w, h float64
}

type jugador struct {
	tex  *sdl.Texture
	x, y float64
}

func nuevoJugador(renderer *sdl.Renderer, x, y float64) (p jugador, err error) {
	img, err := sdl.LoadBMP("sprites/lel.bmp")
	if err != nil {
		return jugador{}, fmt.Errorf("Cargando bmp: %v", err)
	}
	defer img.Free()

	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return jugador{}, fmt.Errorf("Cargando textura: %v", err)
	}
	p.x = x
	p.y = y

	return p, nil
}

func (p *jugador) dibujar(renderer *sdl.Renderer) {
	x := p.x
	y := p.y

	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 300, H: 320},
		&sdl.Rect{X: int32(x), Y: int32(y), W: TAMJUGADOR, H: TAMJUGADOR},
	)
}

func (p *jugador) actualizar() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 && colisionIzq(p) != true {
		if p.x > 0 {
			p.x -= VEL
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 && colisionDer(p) != true {
		if p.x+TAMJUGADOR < ancho {
			p.x += VEL
		}
	} else if keys[sdl.SCANCODE_UP] == 1 && colisionArr(p) != true {
		if p.y > 0 {
			p.y -= VEL
		}
	} else if keys[sdl.SCANCODE_DOWN] == 1 && colisionAbj(p) != true {
		if p.y+TAMJUGADOR < alto {
			p.y += VEL
		}
	}
}

func colisionIzq(p *jugador) bool {

	r1 := cuadrado{x: p.x, y: p.y, w: TAMJUGADOR, h: TAMJUGADOR}

	for _, m := range muros {
		r2 := cuadrado{x: m.x, y: m.y, w: TAMMURO, h: TAMMURO}
		colizq := (r1.x-VEL < r2.x+r2.w && r1.x-VEL+r1.w > r2.x && r1.y < r2.y+r2.h && r1.w+r1.y > r2.y)
		if colizq == true {
			return true
		}
	}
	return false

	// return (r1.x < r2.x + r2.w && r1.x + r1.w > r2.x && r1.y < r2.y + r2.h && r1.w + r1.y > r2.y)
}

func colisionDer(p *jugador) bool {

	r1 := cuadrado{x: p.x, y: p.y, w: TAMJUGADOR, h: TAMJUGADOR}

	for _, m := range muros {
		r2 := cuadrado{x: m.x, y: m.y, w: TAMMURO, h: TAMMURO}
		colizq := (r1.x+VEL < r2.x+r2.w && r1.x+VEL+r1.w > r2.x && r1.y < r2.y+r2.h && r1.w+r1.y > r2.y)
		if colizq == true {
			return true
		}
	}
	return false
}

func colisionArr(p *jugador) bool {

	r1 := cuadrado{x: p.x, y: p.y, w: TAMJUGADOR, h: TAMJUGADOR}

	for _, m := range muros {
		r2 := cuadrado{x: m.x, y: m.y, w: TAMMURO, h: TAMMURO}
		colizq := (r1.x < r2.x+r2.w && r1.x+r1.w > r2.x && r1.y-VEL < r2.y+r2.h && r1.w+r1.y-VEL > r2.y)
		if colizq == true {
			return true
		}
	}
	return false
}

func colisionAbj(p *jugador) bool {

	r1 := cuadrado{x: p.x, y: p.y, w: TAMJUGADOR, h: TAMJUGADOR}

	for _, m := range muros {
		r2 := cuadrado{x: m.x, y: m.y, w: TAMMURO, h: TAMMURO}
		colizq := (r1.x < r2.x+r2.w && r1.x+r1.w > r2.x && r1.y+VEL < r2.y+r2.h && r1.w+r1.y+VEL > r2.y)
		if colizq == true {
			return true
		}
	}
	return false
}
