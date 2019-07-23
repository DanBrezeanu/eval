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

func evaluateSources(flagsSmart, linksSmart, argsSmart *SmartEntry) bool {
	defer compiler.EraseErrorHandler()

	links := linksSmart.GetButtonTexts()
	flags := flagsSmart.GetButtonTexts()
	args := argsSmart.GetButtonTexts()

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

	fileChooser := NewSmartFileChooser()

	vbox.Append(fileChooser.Grid, false)

	form := ui.NewForm()
	vbox.Append(form, false)
	form.SetPadded(true)

	flags := NewSmartEntry()
	links := NewSmartEntry()
	args := NewSmartEntry()

	form.Append("Flags", flags.Hbox, false)
	form.Append("Links", links.Hbox, false)
	form.Append("Args", args.Hbox, false)

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

	compileButton.OnClicked(func(*ui.Button) {
		ip.Show()
		defer ip.Hide()

		evaluateSources(flags, links, args)
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
