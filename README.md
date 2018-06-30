# MatriGo

[![Documentation](https://godoc.org/github.com/tommyblue/matrigo?status.svg)](http://godoc.org/github.com/tommyblue/matrigo)

[![CircleCI](https://circleci.com/gh/tommyblue/matrigo/tree/master.svg?style=svg)](https://circleci.com/gh/tommyblue/matrigo/tree/master)

[![Coverage Status](https://coveralls.io/repos/github/tommyblue/matrigo/badge.svg?branch=master)](https://coveralls.io/github/tommyblue/matrigo?branch=master)

[![Go Report Card](https://goreportcard.com/badge/github.com/tommyblue/matrigo)](https://goreportcard.com/report/github.com/tommyblue/matrigo)

MatriGo is an abstraction layer library over [go-sdl2](https://github.com/veandco/go-sdl2) for matrix-based games

API Documentation: https://godoc.org/github.com/tommyblue/matrigo

## Features

- Draw a custom size matrix of images
- Preload images (cache)
- Receive mouse input from the user (left and right click) and send it to the client by callbacks

## To do

- Make simple shapes
- Color shapes (also support gradients)
- Write stuff
- Use different fonts for text (now it only loads different ttf files)
- Fullscreen window
- Keyboard events
- Draw the matrix by an offset (e.g. move it) into the window and manage the mouse click accordingly

## Setup

Run `./script/setup`, that's all.
The file will install sdl2 dependencies for Mac or Debian-based distros.

### Other o.s.

`go-sdl2` package is the only dependency, check its README file to know how to install it and its
dependencies (`sdl2`): https://github.com/veandco/go-sdl2

## Usage

TODO

## Examples

```go
import "github.com/tommyblue/matrigo"

func InitMyUI() *UI {
    matrigoConf := &matrigo.Conf{
		TileSize: 40,
		Window: &matrigo.Window{
			Width:  800,
			Height: 600,
		},
		Debug: true,
		Title: "Minesweeper",
		Fonts: map[string]matrigo.FontConfig{
			"mono": {
				Path: getAbsolutePath("../assets/fonts/mono.ttf"),
				Size: 14,
			},
		},
		BackgroundColor: &[4]uint8{255, 255, 255, 255},
		BackgroundImage: getAbsolutePath("../assets/images/bg.jpg"),
		ImagesToCache:   getImagesToCache(),
		SyncFPS:         true,
    }
    err := matrigo.Init(matrigoConf)
	if err != nil {
		panic(err)
    }

    ui := &UI{}
	mapInputToFn(ui)

	return ui
}

var tileImages = map[Tile]string{
	Empty:     "empty",
	Mine:      "mine",
	Flag:      "flag",
}

func getImagesToCache() *map[string]string {
	ret := make(map[string]string)
	for _, v := range tileImages {
		ret[v] = getAbsolutePath(fmt.Sprintf("../assets/images/%s.png", v))
	}
	return &ret
}

type matrigoInputInterface struct {
	mouseLeftClickDownFn  func(x, y int32)
	mouseRightClickDownFn func(x, y int32)
	quitFn                func()
}

var input *matrigoInputInterface

func mapInputToFn(ui *UI) {
	if input == nil {
		input = &matrigoInputInterface{
			quitFn:                ui.stopRunning,
			mouseLeftClickDownFn:  ui.mouseLeftClickAt,
			mouseRightClickDownFn: ui.mouseRightClickAt,
		}
	}
}

func (i *matrigoInputInterface) MouseLeftClickDown(x, y int32) {
	i.mouseLeftClickDownFn(x, y)
}
func (i *matrigoInputInterface) MouseRightClickDown(x, y int32) {
	i.mouseRightClickDownFn(x, y)
}
func (i *matrigoInputInterface) Quit() {
	i.quitFn()
}
```

## Changelog

- 2018-06-29: super-mega-alpha
