package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

var boxX int = 10
var boxY int = 10
var cl termbox.Attribute
var sz int = 5
var bottom_text string = ""

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func redraw_all() {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)

	tbprint(0, 0, termbox.ColorGreen|termbox.AttrBold, termbox.ColorDefault,
		"Try dragging the rectange, using mouse scroll, and clicking everywhere")
	tbprint(0, 1, termbox.ColorWhite, termbox.ColorDefault,
		"Press ESC to exit demo")

	_, h := termbox.Size()
	tbprint(0, h-1, termbox.ColorDefault, termbox.ColorDefault, bottom_text)

	for x := boxX; x < boxX+sz; x++ {
		for y := boxY; y < boxY+sz; y++ {
			termbox.SetCell(x, y, '*', cl, termbox.ColorBlack)
		}
	}

	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	mousePressed := 0
	diffX := 0
	diffY := 0

	redraw_all()
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			}
		case termbox.EventMousePress:
			bottom_text = ""
			if ev.Key == termbox.MouseLeft {
				if ev.MouseX >= boxX && ev.MouseX < boxX+sz && ev.MouseY >= boxY && ev.MouseY < boxY+sz {
					mousePressed = 1
					diffX = ev.MouseX - boxX
					diffY = ev.MouseY - boxY
					cl = termbox.ColorYellow
				}
			}
		case termbox.EventMouseRelease:
			if ev.Key == termbox.MouseLeft {
				mousePressed = 0
				cl = termbox.ColorWhite
			}
		case termbox.EventMouseMove:
			if mousePressed == 1 {
				boxX = ev.MouseX - diffX
				boxY = ev.MouseY - diffY
			}
		case termbox.EventMouseScroll:
			if ev.MouseY < 0 && sz > 2 {
				sz--
			}
			if ev.MouseY > 0 && sz < 15 {
				sz++
			}
		case termbox.EventMouseClick:
			button := "Left"
			if ev.Key == termbox.MouseRight {
				button = "Right"
			} else if ev.Key == termbox.MouseMiddle {
				button = "Middle"
			}
			bottom_text = fmt.Sprintf("Mouse button %s clicked at %d : %d", button, ev.MouseX, ev.MouseY)
			if boxX <= ev.MouseX && boxX+sz > ev.MouseX && boxY <= ev.MouseY && boxY+sz > ev.MouseY {
				bottom_text = bottom_text + " - You clicked the box"
			}
		case termbox.EventError:
			panic(ev.Err)
		}
		redraw_all()
	}
}
