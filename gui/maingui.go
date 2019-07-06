package gui

import (
	"fmt"

	"github.com/DanBrezeanu/eval/evaluators"
	"github.com/andlabs/ui"
)

var MainWin *ui.Window
var compiler evaluators.Compiler

func createCompiler(filename string) {
	compiler = evaluators.NewGccCompiler()

	compiler.AddSources(filename)
	compiler.CompileSources()
	fmt.Println(compiler.RunExec())
}

func makeEvalTab() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	grid := ui.NewGrid()
	grid.SetPadded(true)

	button := ui.NewButton("Open File")
	sourceEntry := ui.NewEntry()
	// flagsEntry := ui.NewEntry()
	// linksEntry := ui.NewEntry()

	button.OnClicked(func(*ui.Button) {
		filename := ui.OpenFile(MainWin)
		if filename == "" {
			filename = "(cancelled)"
		}
		sourceEntry.SetText(filename)
		if filename != "(cancelled)" {
			createCompiler(filename)
		}
	})

	grid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(sourceEntry,
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
