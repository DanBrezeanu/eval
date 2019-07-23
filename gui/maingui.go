package gui

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/DanBrezeanu/eval/evaluators"
	"github.com/andlabs/ui"
)

var MainWin *ui.Window
var compiler evaluators.Compiler

func makeOnChangedForEntries(entry *ui.Entry, parent *ui.Box) {
	if len(entry.Text()) > 0 && len(strings.TrimSpace(entry.Text())) > 0 {
		if entry.Text()[len(entry.Text())-1] == ' ' {
			parent.Append(func() *ui.Button {
				button := ui.NewButton(strings.TrimSpace(entry.Text()))
				button.OnClicked(func(*ui.Button) {
					button.Hide()
				})

				return button
			}(), false)
			entry.SetText("")
		}
	}
}

func createCompiler(filename string) {
	//TODO switch compiler

	compiler = evaluators.NewGccCompiler()

	compiler.AddSources(filename)
}

func evaluateSources(flags, links, args string) bool {
	defer compiler.EraseErrorHandler()

	compiler.AddLinks(strings.Split(links, " ")...)
	compiler.AddFlags(strings.Split(flags, " ")...)
	compiler.AddArgs(strings.Split(args, " ")...)

	compiler.CompileSources()

	if compiler.RaisedError() {
		switch compiler.GetErrorHandler().WhatType() {
		case evaluators.NoCompilerFound:
			ui.MsgBoxError(MainWin,
				compiler.GetErrorHandler().What(),
				compiler.GetName()+" could not be found. Try installing it or adding it in PATH.")
			return false

		case evaluators.CompileError:
			ui.MsgBoxError(MainWin,
				compiler.GetErrorHandler().What(),
				compiler.GetErrorHandler().Error())
			return false
		}
	} else {
		fmt.Println(compiler.RunExec())
		if compiler.RaisedError() {
			switch compiler.GetErrorHandler().WhatType() {
			case evaluators.RunTimeError:
				ui.MsgBoxError(MainWin,
					compiler.GetErrorHandler().What(),
					compiler.GetErrorHandler().Error())
				return false
			}
		}
	}
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

	flagsHorizontalBox := ui.NewHorizontalBox()
	flagsHorizontalBox.SetPadded(true)

	linksHorizontalBox := ui.NewHorizontalBox()
	linksHorizontalBox.SetPadded(true)

	argsHorizontalBox := ui.NewHorizontalBox()
	argsHorizontalBox.SetPadded(true)

	flagsEntry := ui.NewEntry()
	flagsHorizontalBox.Append(flagsEntry, true)

	flagsEntry.OnChanged(func(*ui.Entry) {
		makeOnChangedForEntries(flagsEntry, flagsHorizontalBox)
	})

	linksEntry := ui.NewEntry()
	linksHorizontalBox.Append(linksEntry, true)

	linksEntry.OnChanged(func(*ui.Entry) {
		makeOnChangedForEntries(linksEntry, linksHorizontalBox)
	})

	argsEntry := ui.NewEntry()
	argsHorizontalBox.Append(argsEntry, true)

	argsEntry.OnChanged(func(*ui.Entry) {
		makeOnChangedForEntries(argsEntry, argsHorizontalBox)
	})

	multiSourceEntry := ui.NewNonWrappingMultilineEntry()

	form.Append("Flags", flagsHorizontalBox, false)
	form.Append("Links", linksHorizontalBox, false)
	form.Append("Args", argsHorizontalBox, false)

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
					sourceEntry.Hide()
				} else { /* multiple sources */
					multiSourceEntry.SetText(multiSourceEntry.Text() + "\n" + filename)
				}
			}
		}
	})

	cbox := ui.NewCombobox()
	cbox.Append("gcc")
	cbox.Append("java")
	cbox.Append("python3")
	cbox.SetSelected(0)
	vbox.Append(cbox, false)

	compileButton := ui.NewButton("Evaluate")

	for i := 0; i < 6; i++ {
		vbox.Append(ui.NewLabel(""), false)
	}
	vbox.Append(compileButton, false)

	ip := ui.NewProgressBar()
	ip.SetValue(-1)
	ip.Hide()
	vbox.Append(ip, false)

	F := ui.NewVerticalBox()
	compileButton.OnClicked(func(*ui.Button) {
		ip.Show()
		defer ip.Hide()

		for idxC, C := range F.GetChildren( {
			fmt.Println(reflect.TypeOf(C))
		}
		evaluateSources(flagsEntry.Text(), linksEntry.Text(), argsEntry.Text())
		// errors

	})

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
