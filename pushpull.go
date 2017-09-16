package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/wargarblgarbl/libgosubs/ass"
)


var subcount int


func printf_tb(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	print_tb(x, y, fg, bg, s)
}
func print_tb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func updatewin(sub int, event string) {

}

func subup (ass *ass.Ass) {
  if subcount != 0 {
    subcount = subcount - 1
  }
  event := ass.Events.Body[subcount].Text
  printf_tb(1, 1, termbox.ColorBlack, termbox.ColorWhite, event )
  termbox.Flush()
}


func subdown (ass *ass.Ass) {
  if len(ass.Events.Body)-1 > subcount {
    subcount = subcount + 1
  }
  event := ass.Events.Body[subcount].Text
  printf_tb(1, 1, termbox.ColorBlack, termbox.ColorWhite, event )
  termbox.Flush()
}



func main() {
  subcount = 0
	load, err := ass.ParseAss("./test.ass")
	if err != nil {
		panic(err)
	}

	errb := termbox.Init()
	if errb != nil {
		panic(errb)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlS {
				termbox.Sync()
			}
			if ev.Key == termbox.KeyCtrlQ {
				fmt.Println("FART")
			}
			if ev.Key == termbox.KeyArrowUp {
        subup(load)

			}
			if ev.Key == termbox.KeyArrowDown {
        subdown(load)
			}
      if ev.Key == termbox.KeyEsc {
      panic("QUIT")}

		}
}
}
