// Demo code for the DropDown primitive.
package main

import "github.com/cainy-a/gord/tview"

func main() {
	app := tview.NewApplication()
	dropdown := tview.NewDropDown().
		SetLabel("Select an option (hit Enter): ").
		SetOptions([]string{"First", "Second", "Third", "Fourth", "Fifth"}, nil)
	if err := app.SetRoot(dropdown, true).Run(); err != nil {
		panic(err)
	}
}
