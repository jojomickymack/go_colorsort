package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

func time_left(next_time uint32) uint32 {
	var now uint32 = sdl.GetTicks()

	if next_time <= now {
		return 0
	} else {
		return next_time - now
	}
}

const (
	winWidth     = 940
	winHeight    = 720
	tickInterval = 30
)

type Button struct {
	left    int32
	right   int32
	top     int32
	bottom  int32
	pressed bool
}

var (
	next_time uint32

	err      error
	window   *sdl.Window
	renderer *sdl.Renderer

	quit      bool
	event     sdl.Event
	locationX = 0.0
	locationY = 0.0
	step      = 0.5
	deviation = 2.0

	sorted bool = false

	rb = Button{100, 150, winHeight - 10, winHeight - 40, false}
	gb = Button{170, 220, winHeight - 10, winHeight - 40, false}
	bb = Button{240, 290, winHeight - 10, winHeight - 40, false}
	ab = Button{310, 360, winHeight - 10, winHeight - 40, false}
)

func run() int {
	// var boxX int32 = 100
	// var boxY int32 = 100

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize SDL: %s\n", err)
		return 1
	}
	defer sdl.Quit()

	if window, err = sdl.CreateWindow("colorsort", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 2
	}
	defer window.Destroy()

	if renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 3 // don't use os.Exit(3); otherwise, previous deferred calls will never run
	}
	//renderer.SetDrawColor(0.0, 0.0, 0.0, 0.0)
	renderer.Clear()
	defer renderer.Destroy()

	cl := NewColorList(winWidth / 2)
	myList := cl.list
	lw := winWidth / len(myList)

	quit = false
	for !quit {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				quit = true
			case *sdl.MouseButtonEvent:
				if t.Button == 1 && t.State == 1 {
					if t.X > rb.left && t.X < rb.right && t.Y > rb.bottom && t.Y < rb.top {
						rb.pressed = true
					} else {
						rb.pressed = false
					}
					if t.X > gb.left && t.X < gb.right && t.Y > gb.bottom && t.Y < gb.top {
						gb.pressed = true
					} else {
						gb.pressed = false
					}
					if t.X > bb.left && t.X < bb.right && t.Y > bb.bottom && t.Y < bb.top {
						bb.pressed = true
					} else {
						bb.pressed = false
					}
					if t.X > ab.left && t.X < ab.right && t.Y > ab.bottom && t.Y < ab.top {
						ab.pressed = true
					} else {
						ab.pressed = false
					}
				} else if t.Button == 1 && t.State == 0 {
					rb.pressed = false
					gb.pressed = false
					bb.pressed = false
					ab.pressed = false
				}
			}
		}

		renderer.Clear()
		gfx.BoxRGBA(renderer, 0, 0, winWidth, winHeight, 0, 0, 0, 255)
		//gfx.BoxColor(renderer, boxX, boxY, boxX + 100, boxY + 100, sdl.Color{0, 255, 255, 255})

		for i := 0; i < len(myList); i++ {
			gfx.BoxRGBA(renderer, int32(lw*i), int32(10), int32(lw*i+lw-2), int32(winHeight-50), myList[i].r, myList[i].g, myList[i].b, myList[i].a)
		}

		if rb.pressed {
			gfx.BoxRGBA(renderer, rb.left, rb.bottom, rb.right, rb.top, 255, 0, 0, 255)
			myList = cl.sortRed(*cl)
		} else {
			gfx.RectangleRGBA(renderer, rb.left, rb.bottom, rb.right, rb.top, 255, 0, 0, 255)
		}

		if gb.pressed {
			gfx.BoxRGBA(renderer, gb.left, gb.bottom, gb.right, gb.top, 0, 255, 0, 255)
			myList = cl.sortGreen(*cl)
		} else {
			gfx.RectangleRGBA(renderer, gb.left, gb.bottom, gb.right, gb.top, 0, 255, 0, 255)
		}

		if bb.pressed {
			gfx.BoxRGBA(renderer, bb.left, bb.bottom, bb.right, bb.top, 0, 0, 255, 255)
			myList = cl.sortBlue(*cl)
		} else {
			gfx.RectangleRGBA(renderer, bb.left, bb.bottom, bb.right, bb.top, 0, 0, 255, 255)
		}

		if ab.pressed {
			gfx.BoxRGBA(renderer, ab.left, ab.bottom, ab.right, ab.top, 255, 255, 255, 255)
			myList = cl.sortAlpha(*cl)
		} else {
			gfx.RectangleRGBA(renderer, ab.left, ab.bottom, ab.right, ab.top, 255, 255, 255, 255)
		}

		renderer.Present()
		sdl.Delay(time_left(next_time))
		next_time += tickInterval
	}
	return 1
}

func main() {
	os.Exit(run())
}
