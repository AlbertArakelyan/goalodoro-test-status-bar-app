package main

import (
	"fmt"

	"github.com/caseymrm/menuet"
)

var (
	activeItem string
)

func defaultMenuState() *menuet.MenuState {
	return &menuet.MenuState{
		Image: "https://static-00.iconduck.com/assets.00/dart-icon-2043x2048-3dq18idp.png",
	}
}

func menuItems() []menuet.MenuItem {
	return []menuet.MenuItem{
		{
			Text: "Hello",
			Clicked: func() {
				switchClock("Hello")
			},
			State: activeItem == "Hello",
		},
		{
			Text: "World",
			Clicked: func() {
				switchClock("World")
			},
			State: activeItem == "World",
		},
		{
			Text: "Add...",
			Clicked: func() {
				input := menuet.App().Alert(menuet.Alert{
					MessageText: "Add new clock",
					Inputs: []string{
						"Test",
					},
				})

				// Test chanage with real one
				fmt.Println(input.Inputs[0])
				
				menuItems := menuet.App().Children()
				menuItems = append(menuItems, menuet.MenuItem{
					Text: input.Inputs[0],
					Clicked: func() {
						switchClock(input.Inputs[0])
					},
					State: false,
				})
				menuet.App().Children = func() []menuet.MenuItem {
					return menuItems
				}
			},
			State: false,
		},
		{
			Text: "Stop",
			Clicked: func() {
				menuet.App().SetMenuState(defaultMenuState())
				activeItem = ""
				close(stop)
			},
			State: false,
		},
	}
}

func main() {
	menuet.App().SetMenuState(defaultMenuState())

	menuet.App().Children = menuItems

	menuet.App().RunApplication()
}
