package matrigo

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func (ui *sdlWrapper) syncFPS() {
	// Reset timers
	ui.previousTimer = ui.timer
	ui.timer = time.Now()

	tick := time.Now()
	elapsedMS := float64(tick.Sub(ui.timer)) / float64(time.Millisecond)
	if sleep := ticksPerFrame - elapsedMS; sleep > 0 {
		d := time.Duration(sleep)
		sdl.Delay(uint32(d))
	}
}
