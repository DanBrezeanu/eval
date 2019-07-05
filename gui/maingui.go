package gui

import (
	"github.com/andlabs/ui"
)

var MainWin *ui.Window

func makeEvalTab() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	grid := ui.NewGrid()
	grid.SetPadded(true)

	button := ui.NewButton("Open File")
	entry := ui.NewEntry()

	button.OnClicked(func(*ui.Button) {
		filename := ui.OpenFile(MainWin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry.SetText(filename)
	})

	grid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry,
		1, 0, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)
	vbox.Append(grid, false)

	cbox := ui.NewCombobox()
	cbox.Append("gcc")
	cbox.Append("java")
	cbox.SetSelected(0)
	vbox.Append(cbox, false)

	return vbox
}

func setupUI() {
	MainWin = ui.NewWindow("Eval", 640, 480, true)
	defer MainWin.Show()

	MainWin.OnClosing(
		func(*ui.Window) bool {
			ui.Quit()
			return true
		})

	ui.OnShouldQuit(
		func() bool {
			MainWin.Destroy()
			return true
		})

	tab := ui.NewTab()
	MainWin.SetChild(tab)
	MainWin.SetMargined(true)

	tab.Append("EvalTab", makeEvalTab())
	tab.SetMargined(0, true)
}

func MainGui() {
	ui.Main(setupUI)
}
