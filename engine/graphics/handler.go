package graphics

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/LinMAD/Snap/engine/platform"
)

// CreateNativeWindow for engine
func CreateNativeWindow(cfg *platform.ScreenConfiguration) (win *sdl.Window, winErr error) {
	processScreenConfig(cfg)

	if initErr := sdl.Init(sdl.INIT_EVERYTHING); initErr != nil {
		return nil, fmt.Errorf("failed to initialise all subsystems for media layer: %s\n", initErr)
	}

	win, winErr = sdl.CreateWindow(
		cfg.Title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		int32(cfg.Width),
		int32(cfg.Height),
		parseWindowFlags(cfg),
	)
	if winErr != nil {
		return nil, fmt.Errorf("failed to create native window: %s\n", winErr)
	}

	return win, nil
}

// CreateRenderer for native window
func CreateRenderer(win *sdl.Window) (rend *sdl.Renderer, rendErr error) {
	rend, rendErr = sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED)
	if rendErr != nil {
		return nil, fmt.Errorf("failed to create renderer for window: %s\n", rendErr)
	}

	return rend, nil
}

func parseWindowFlags(cfg *platform.ScreenConfiguration) (flag uint32) {
	flag = sdl.WINDOW_SHOWN

	if cfg.IsFullScreen {
		flag |= sdl.WINDOW_FULLSCREEN
	}
	if cfg.IsResizeable {
		flag |= sdl.WINDOW_RESIZABLE
	}

	return
}

func processScreenConfig(cfg *platform.ScreenConfiguration) {
	if cfg.FrameRateLock == 0 {
		cfg.FrameRateLock = 60
	}
	if len(cfg.Title) == 0 {
		cfg.Title = "Snap Engine"
	}
	if cfg.Width == 0 {
		cfg.Width = 800
	}
	if cfg.Height == 0 {
		cfg.Height = 600
	}
}
