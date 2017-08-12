package main
import (
//	"fmt"
//	"github.com/gizak/termui"
	"github.com/marcusolsson/tui-go"
	"github.com/wargarblgarbl/libgosubs/ass"
	

)




func main() {
	//submem := &ass.Ass{}
//	load := ass.ParyseAss("./test.ass")
	//	fmt.Print(load)

	content := tui.NewTable(0,0)
	content.SetColumnStretch(0, 1)
	content.SetColumnStretch(1, 1)
	content.SetColumnStretch(2, 4)
	var (
	//	start = tui.NewLabel("")
	//	end = tui.NewLabel("")
		text = tui.NewLabel("")
	)
	header := tui.NewTable(0,0)
	header.SetColumnStretch(0, 1)
	header.SetColumnStretch(1, 1)
	header.SetColumnStretch(2, 4)
	header.AppendRow(
		tui.NewLabel("Start"),
		tui.NewLabel("End"),
		tui.NewLabel("Text"),
	
	)

	
//	content.AppendRow(
//		tui.NewLabel("Start"),
//		tui.NewLabel("End"),
//		tui.NewLabel("Text"),
		
//	)
	
	load := ass.ParseAss("./test.ass")
	for _, i := range load.Events.Body {
	//	label := setChannelLabel(slackChan, false)
	/*	content.Append(tui.NewHBox(
			tui.NewLabel(i.Start+"|"+i.End),
			tui.NewPadder(1, 0, tui.NewLabel("|")),
			tui.NewLabel(i.Text),
			tui.NewSpacer(),
		))
*/
	//	fmt.Println(i)
		content.AppendRow(
			
			tui.NewLabel(i.Start),
			tui.NewLabel(i.End),
			tui.NewLabel(i.Text),
	
		)
			//tui.NewHBox(
//			tui.NewLabel(i.Start),
//			tui.NewLabel(i.End),
	
		
		
	}


////////

	
	input := tui.NewEntry()
	input.SetFocused(true)

	input.SetSizePolicy(tui.Maximum, tui.Maximum)

//	inputBox := tui.NewHBox(input)
//	inputBox.SetBorder(true)
//	inputBox.SetFocused(true)
//	inputBox.SetSizePolicy(0, tui.Maximum)
	
	/////
	
	textbx := tui.NewGrid(0,0)
	textbx.SetSizePolicy(tui.Maximum, tui.Maximum)

	textbx.AppendRow(tui.NewLabel("Current Selection:"),text)
	textbx.SetBorder(true)
	textbx.SetColumnStretch(0, 1)
	textbx.SetColumnStretch(1, 1)


	
	content.OnItemActivated(func(t *tui.Table) {
	//.OnSelectionChanged(func(t *tui.Table) {
		m := load.Events.Body[t.Selected()]
	//	input.SetSelected(m.Text)
	//	input.SetFocused(true)
		text.SetText(m.Text)
	//	textbx.SetText(m.Text)
	})

		



	content.Select(0)
	editor := tui.NewVBox(header, content, textbx, input)
	
	editor.SetSizePolicy(tui.Expanding, tui.Expanding)
	root := tui.NewHBox(editor)


//	tui.DefaultFocusChain.Set(content, input)

	ui := tui.New(root)
	
	ui.SetKeybinding("Esc", func() { ui.Quit() })
	if err := ui.Run(); err != nil {
		panic(err)
	}
	
}
