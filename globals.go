package matrigo

import (
	"time"

	"github.com/veandco/go-sdl2/ttf"
)

var ui *sdlWrapper

var fonts map[string]*ttf.Font
var defaultFont *ttf.Font

// fpsTarget is the frame rate target
const fpsTarget = 60

// ticksPerFrame is the number of ticks required to reach FPS
const ticksPerFrame = 1000.0 / fpsTarget

var lastTimestamp time.Time
var lastFPS int
var countSinceLast int
