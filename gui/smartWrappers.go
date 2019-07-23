package gui

import (
	"strings"

	"github.com/andlabs/ui"
)

type SmartEntry struct {
	Entry   *ui.Entry
	Hbox    *ui.Box
	Buttons []*ui.Button
}

func NewSmartEntry() *SmartEntry {
	var smartEntry SmartEntry = SmartEntry{}

	smartEntry.Entry = ui.NewEntry()
	smartEntry.Hbox = ui.NewHorizontalBox()
	smartEntry.Hbox.Append(smartEntry.Entry, true)
	smartEntry.Hbox.SetPadded(true)

	smartEntry.Entry.OnChanged(func(*ui.Entry) {
		text := smartEntry.Entry.Text()

		if len(text) > 0 && len(strings.TrimSpace(text)) > 0 {
			if text[len(text)-1] == ' ' {
				newButton := func() *ui.Button {
					button := ui.NewButton(strings.TrimSpace(text))
					button.OnClicked(func(*ui.Button) {
						button.Hide()

						for i, smartButton := range smartEntry.Buttons {
							if smartButton == button {
								smartEntry.Buttons = append(smartEntry.Buttons[:i], smartEntry.Buttons[i+1:]...)
								break
							}
						}
					})

					return button
				}()
				smartEntry.Hbox.Append(newButton, false)
				smartEntry.Entry.SetText("")
				smartEntry.Buttons = append(smartEntry.Buttons, newButton)
			}
		}
	})

	return &smartEntry
}

func (smart *SmartEntry) GetButtonTexts() string {
	resultString := ""

	for _, button := range smart.Buttons {
		resultString = resultString + " " + button.Text()
	}

	return resultString
}

type SmartFileChooser struct {
	Button         *ui.Button
	Entry          *ui.Entry
	MultilineEntry *ui.MultilineEntry
	Grid           *ui.Grid
	MultipleLines  bool
}

func NewSmartFileChooser() *SmartFileChooser {
	var smart SmartFileChooser = SmartFileChooser{}

	smart.Button = ui.NewButton("Open File")
	smart.Entry = ui.NewEntry()
	smart.Grid = ui.NewGrid()

	smart.Grid.Append(smart.Button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	smart.Grid.Append(smart.Entry,
		1, 0, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	smart.MultipleLines = false

	smart.Button.OnClicked(func(*ui.Button) {
		filename := ui.OpenFile(MainWin)
		if filename != "" { /* added succesfully */
			if compiler == nil { /* first source */
				smart.Entry.SetText(filename)
				createCompiler(filename)
			} else { /* multiple sources */
				compiler.AddSources(filename)
				if !smart.MultipleLines { /* two sources */
					smart.MultilineEntry = ui.NewNonWrappingMultilineEntry()
					smart.MultilineEntry.SetText(smart.Entry.Text() + "\n" + filename)
					smart.Grid.Append(smart.MultilineEntry,
						1, 0, 1, 1,
						true, ui.AlignFill, true, ui.AlignFill)
					smart.Entry.Hide()
					smart.MultipleLines = true
				} else { /* multiple sources */
					smart.MultilineEntry.SetText(smart.MultilineEntry.Text() + "\n" + filename)
				}
			}
		}
	})

	return &smart
}
