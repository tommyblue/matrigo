package matrigo

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Input is the interface that must be implemented by clients to receive input callbacks.
// The implementation must then be passed to ManageInput that will call the callbacks when the
// corresponding event will happen
type Input interface {
	// x and y passed to the mouse click functions are the position, in pixels, based on the Window
	// where (0, 0) is the top-left (NW) corner. The receiver should then convert the position to
	// the clicked tile
	MouseLeftClickDown(x, y int32)
	MouseRightClickDown(x, y int32)
	// TODO: keyboard event
	Quit()
}

// ManageInput listens for events from the user and calls the callbacks
// provided by the Input interface
func ManageInput(inputInterface Input) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch ev := event.(type) {
		case *sdl.MouseButtonEvent:
			switch event.(*sdl.MouseButtonEvent).Type {
			case sdl.MOUSEBUTTONDOWN:
				switch ev.Button {
				case sdl.BUTTON_LEFT:
					if ui.conf.Debug {
						fmt.Printf("Mouse left click at (%d, %d)\n", ev.X, ev.Y)
					}
					inputInterface.MouseLeftClickDown(ev.X, ev.Y)
					break
				case sdl.BUTTON_RIGHT:
					if ui.conf.Debug {
						fmt.Printf("Mouse right click at (%d, %d)\n", ev.X, ev.Y)
					}
					inputInterface.MouseRightClickDown(ev.X, ev.Y)
					break
				}
				break
			}
			break
		case *sdl.QuitEvent:
			if ui.conf.Debug {
				println("Quitting...")
			}
			inputInterface.Quit()
			break
		}
	}
}
