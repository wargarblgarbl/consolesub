package main
import (
//	"fmt"
	"github.com/gizak/termui"

	"github.com/wargarblgarbl/libgosubs/ass"

)




func main() {
	//submem := &ass.Ass{}
//	load := ass.ParyseAss("./test.ass")
//	fmt.Print(load)
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()
	subtitles := termui.NewTable()

	var rows [][]string
	var text []string
	fat := 0
	tfat := 0
	load := ass.ParseAss("./test.ass")
	for _, i := range load.Events.Body {
	//	label := setChannelLabel(slackChan, false)
		var row []string
		
		row = append(row, i.Start)
		row = append(row, i.End)
		text = append(text, i.Text)
		rows = append(rows, row)
		text = append(text, "")
		fatter := len(i.Start) + len(i.End) +9
		if fat < fatter {
			fat = fatter
		}
		tfatter := len(text)
		if tfat < tfatter {
			tfat = tfatter
		}
		
	}
	subtitles.Rows= rows
	subtitles.Y = 0
	subtitles.X = 0
	subtitles.Width = fat
	
	subtitles.Height = len(rows) *2

	contents := termui.NewList()
	contents.X = fat
	contents.Items = text
	contents.Height = subtitles.Height
	contents.Width = tfat

	


	termui.Render(subtitles, contents)

	
	
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})
	termui.Loop()
}
