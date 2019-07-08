package gui

import (
	"fmt"
	"strings"

	"github.com/DanBrezeanu/eval/evaluators"
	"github.com/andlabs/ui"
)

var MainWin *ui.Window
var compiler evaluators.Compiler

func createCompiler(filename string) {
	//TODO switch compiler

	compiler = evaluators.NewGccCompiler()

	compiler.AddSources(filename)
}

func evaluateSources(flags, links, args string) bool {
	compiler.AddLinks(strings.Split(links, " ")...)
	compiler.AddFlags(strings.Split(flags, " ")...)
	compiler.AddArgs(strings.Split(args, " ")...)

	compiler.CompileSources()
	fmt.Println(compiler.RunExec())

	return true
}

func makeEvalTab() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	grid := ui.NewGrid()
	grid.SetPadded(true)

	button := ui.NewButton("Open File")
	sourceEntry := ui.NewEntry()

	grid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	grid.Append(sourceEntry,
		1, 0, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	vbox.Append(grid, false)

	form := ui.NewForm()
	vbox.Append(form, false)
	form.SetPadded(true)

	flagsEntry := ui.NewEntry()
	linksEntry := ui.NewEntry()
	argsEntry := ui.NewEntry()
	multiSourceEntry := ui.NewNonWrappingMultilineEntry()

	form.Append("Flags", flagsEntry, false)
	form.Append("Links", linksEntry, false)
	form.Append("Args", argsEntry, false)

	button.OnClicked(func(*ui.Button) {
		filename := ui.OpenFile(MainWin)
		if filename != "" { /* added succesfully */
			if compiler == nil { /* first source */
				sourceEntry.SetText(filename)
				createCompiler(filename)
			} else { /* multiple sources */
				compiler.AddSources(filename)
				if multiSourceEntry.Text() == "" { /* two sources */
					multiSourceEntry.SetText(sourceEntry.Text() + "\n" + filename)
					grid.Append(multiSourceEntry,
						1, 0, 1, 1,
						true, ui.AlignFill, true, ui.AlignFill)
				} else { /* multiple sources */
					multiSourceEntry.SetText(multiSourceEntry.Text() + "\n" + filename)
				}
			}
		}
	})

	cbox := ui.NewCombobox()
	cbox.Append("gcc")
	cbox.Append("java")
	cbox.SetSelected(0)
	vbox.Append(cbox, false)

	compileButton := ui.NewButton("Evaluate")
	compileButton.OnClicked(func(*ui.Button) {
		evaluateSources(flagsEntry.Text(), linksEntry.Text(), argsEntry.Text())
	})

	vbox.Append(ui.NewLabel(""), false)
	vbox.Append(ui.NewLabel(""), false)
	vbox.Append(compileButton, false)
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

	hbox := ui.NewHorizontalBox()
	MainWin.SetChild(hbox)
	MainWin.SetMargined(true)

	hbox.Append(makeEvalTab(), true)
}

func MainGui() {
	ui.Main(setupUI)
}
