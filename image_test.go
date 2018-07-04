package matrigo

import (
	"testing"
)

func Test_warmCache(t *testing.T) {
	t.Run("Cache gets created", func(t *testing.T) {
		ui := &sdlWrapper{
			conf: &Conf{
				Debug:         false,
				ImagesToCache: new(map[string]string),
			},
		}
		if ui.cache != nil {
			t.Errorf("ui.cache should be nil at this point")
		}
		ui.warmCache()
		if ui.cache == nil {
			t.Errorf("ui.cache shouldn't be nil at this point")
		}
	})

	t.Run("Background image is cached", func(t *testing.T) {
		ui := &sdlWrapper{
			conf: &Conf{
				Window: &Window{
					Width:  100,
					Height: 100,
				},
				Debug:           false,
				ImagesToCache:   new(map[string]string),
				BackgroundImage: getAbsolutePath("./fixtures/bg.jpg"),
			},
		}
		ui.initSdl()
		ui.initWindow()
		ui.initRenderer()
		err := ui.warmCache()
		if err != nil {
			t.Errorf("Error occurred caching bg: %v", err)
		}
		if (*ui.cache)["matrigo-bg"].image == nil || (*ui.cache)["matrigo-bg"].id != "matrigo-bg" {
			t.Errorf("The background images isn't cached")
		}
	})

	t.Run("Images are cached", func(t *testing.T) {
		ui := &sdlWrapper{
			conf: &Conf{
				Window: &Window{
					Width:  100,
					Height: 100,
				},
				Debug: false,
				ImagesToCache: &map[string]string{
					"img1": getAbsolutePath("./fixtures/bg.jpg"),
					"img2": getAbsolutePath("./fixtures/bg.jpg"),
				},
				BackgroundImage: getAbsolutePath("./fixtures/bg.jpg"),
			},
		}
		ui.initSdl()
		ui.initWindow()
		ui.initRenderer()
		err := ui.warmCache()
		if err != nil {
			t.Errorf("Error occurred caching bg: %v", err)
		}
		if (*ui.cache)["img1"].image == nil || (*ui.cache)["img2"].image == nil {
			t.Errorf("Images aren't cached")
		}
	})
}
