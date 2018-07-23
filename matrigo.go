/*
Package matrigo provides an abstraction layer over go-sdl2.

It works for tile-based words, like famous games Sokoban, Minesweeper and Tetris.

The client sends a configuration object during the initialization and this package
exposes few methods to draw tiles and manage input from the user.

For usage examples see the Readme file in the package repository: https://github.com/tommyblue/matrigo
*/
package matrigo

// Init the package. The client must call this function with a valid
// configuration at the beginning of the ui setup
func Init(c *Conf) error {
	ui = &sdlWrapper{conf: c}

	err := ui.initSdl()
	if err != nil {
		return err
	}

	err = ui.initFonts()
	if err != nil {
		return err
	}

	err = ui.initWindow()
	if err != nil {
		return err
	}

	err = ui.initRenderer()
	if err != nil {
		return err
	}

	err = ui.initTimers()
	if err != nil {
		return err
	}

	err = ui.warmCache()

	return err
}

// Close everything. To be called by the client before quitting.
func Close() error {
	err := ui.closeRenderer()
	if err != nil {
		return err
	}

	err = ui.closeWindow()
	if err != nil {
		return err
	}

	ui.closeSdl()
	return nil
}

// Draw what needs to be drawn down.
// It draws the background image and color, all the images and then syncs
// the frame rate to the value in the fpsTarget variable
func Draw(matrix *Matrix) {
	ui.drawBackground()

	for _, tile := range matrix.Tiles {
		ui.drawImage(*tile, getMatrixOffset(matrix))
	}

	if ui.conf.SyncFPS {
		ui.syncFPS()
	}
	debugFPS()

	ui.renderer.Present()
	if ui.conf.BackgroundColor != nil {
		ui.renderer.SetDrawColor(
			ui.conf.BackgroundColor[0],
			ui.conf.BackgroundColor[1],
			ui.conf.BackgroundColor[2],
			ui.conf.BackgroundColor[3],
		)
	}
	ui.renderer.Clear()
}

// Calculate draw offset of the matrix, currently commented out because mouse events must be
// re-calculated based on offset
func getMatrixOffset(matrix *Matrix) *imageOffset {
	// matrixSide := int32(math.Sqrt(float64(len(*matrix.Tiles))))
	// matrixWidth := matrixSide * ui.conf.TileSize
	// xOffset := int32((ui.conf.Window.Width - matrixWidth) / 2)
	// offset := &imageOffset{
	// 	x: xOffset,
	// 	y: 0,
	// }
	offset := &imageOffset{
		x: 0,
		y: 0,
	}
	return offset
}
