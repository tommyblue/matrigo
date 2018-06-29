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
