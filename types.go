package matrigo

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// FontConfig is used to configure the fonts to use. Each struct must contain the absolute path
// to a ttf font file and its size (in points)
type FontConfig struct {
	Path string // Absolute path to a ttf font file
	Size int    // Font size in points
}

// Window defines the size of the window to create (in pixels)
type Window struct {
	Width  int32 // Width of the window
	Height int32 // Height of the window
}

// Conf is the global configuration of the library
type Conf struct {
	TileSize        int32                 // Size of the side of the tiles (tiles are squared)
	Title           string                // Title of the window
	Fonts           map[string]FontConfig // Fonts to use
	Debug           bool                  // Enable or disable debugging
	Window          *Window               // Window to draw
	BackgroundColor *[4]uint8             // Background color of the window
	BackgroundImage string                // Background image. Must be of the same size of the Window (optional)
	ImagesToCache   *map[string]string    // All the images to add to the cache, as key => abs path of the image
	SyncFPS         bool                  // Sync the frame rate to 60fps
}

// Tile is an image to be drawn down. Each Tile has an absolute path to the image file and its
// position (x and y) in the matrix
type Tile struct {
	ImageID string // Absolute path to the image file
	PosX    int32  // X position (column) of the tile in the matrix
	PosY    int32  // Y position (row) of the tile in the matrix
	OffsetX int32  // Offset in the X direction in pixels
	OffsetY int32  // Offset in the Y direction in pixels
}

// Matrix represents the whole tiles (i.e. images) to draw in the window
type Matrix struct {
	Tiles *[]Tile // Array of tiles
}

type imageStruct struct {
	id    string
	image *sdl.Texture
	rect  sdl.Rect
}

type imageOffset struct {
	x int32
	y int32
}

type sdlWrapper struct {
	cache    *map[string]imageStruct
	conf     *Conf
	window   *sdl.Window
	renderer *sdl.Renderer

	timer         time.Time
	previousTimer time.Time
}
