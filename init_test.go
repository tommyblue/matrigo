package matrigo

import (
	"testing"
)

func Test_sdlWrapper_initTimers(t *testing.T) {
	t.Run("Initialize timers", func(t *testing.T) {
		ui := &sdlWrapper{}
		if !ui.timer.IsZero() {
			t.Errorf("timer already initialized")
		}
		if !ui.previousTimer.IsZero() {
			t.Errorf("previousTimer already initialized")
		}
		if err := ui.initTimers(); err != nil {
			t.Errorf("sdlWrapper.initTimers() error = %v", err)
		}
		if ui.timer.IsZero() {
			t.Errorf("timer isn't initialized")
		}
		if ui.previousTimer.IsZero() {
			t.Errorf("previousTimer isn't initialized")
		}
		if ui.timer != ui.previousTimer {
			t.Errorf("previousTimer incorrectly initialized")
		}
	})
}

func Test_sdlWrapper_initFonts(t *testing.T) {
	t.Run("Initialize fonts", func(t *testing.T) {
		fontsConfig := map[string]FontConfig{
			"first": {
				Path: getAbsolutePath("./fixtures/fonts/mono.ttf"),
				Size: 12,
			},
			"second": {
				Path: getAbsolutePath("./fixtures/fonts/mono.ttf"),
				Size: 14,
			},
		}
		ui := &sdlWrapper{
			conf: &Conf{
				Fonts: fontsConfig,
			},
		}
		if fonts != nil {
			t.Errorf("fonts should be nil at this point")
		}
		err := ui.initFonts()
		if err != nil {
			t.Errorf("Error occurring initializing fonts: %v", err)
		}
		if fonts == nil {
			t.Errorf("fonts shouldn't be nil at this point")
		}
		if len(fonts) != 2 {
			t.Errorf("Wrong number of fonts loaded")
		}
		if defaultFont == nil {
			t.Errorf("The first loaded font should also be the default one")
		}
	})
}

func Test_sdlWrapper_initWindow(t *testing.T) {
	t.Run("Initialize SDL window", func(t *testing.T) {
		ui := &sdlWrapper{
			conf: &Conf{
				Title: "Test",
				Window: &Window{
					Width:  100,
					Height: 100,
				},
			},
		}
		if ui.window != nil {
			t.Errorf("ui.window should be nil at this point")
		}
		err := ui.initWindow()
		if err != nil {
			t.Errorf("Error occurring initializing ui.window: %v", err)
		}
		if ui.window == nil {
			t.Errorf("ui.window shouldn't be nil at this point")
		}
	})
}

func Test_sdlWrapper_initRenderer(t *testing.T) {
	t.Run("Initialize SDL renderer", func(t *testing.T) {
		ui := &sdlWrapper{
			conf: &Conf{
				Title: "Test",
				Window: &Window{
					Width:  100,
					Height: 100,
				},
			},
		}
		err := ui.initWindow()
		if err != nil {
			t.Errorf("Error occurring initializing ui.window: %v", err)
		}
		if ui.renderer != nil {
			t.Errorf("ui.window should be nil at this point")
		}
		err = ui.initRenderer()
		if err != nil {
			t.Errorf("Error occurring initializing ui.renderer: %v", err)
		}
		if ui.renderer == nil {
			t.Errorf("ui.renderer shouldn't be nil at this point")
		}
	})
}
