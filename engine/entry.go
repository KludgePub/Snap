package engine

import (
	"runtime"

	"github.com/LinMAD/SnapEngine/engine/core"
	"github.com/LinMAD/SnapEngine/engine/platform"
)

// Entry point for application
type Entry struct {
	sc      *platform.ScreenConfiguration
	isDebug bool
}

func init() {
	runtime.LockOSThread()
}

// New engine entry point for application
func New(screenConfig *platform.ScreenConfiguration, isDebugMode bool) Entry {
	return Entry{
		sc:      screenConfig,
		isDebug: isDebugMode,
	}
}

// TODO Allow to load any level in any time
// TODO Add built in level with loading information, init, loading user plugin, errors feedback

// Run engine work
func (e *Entry) Run() error {
	var frameStart uint32  // Initial time of one frame
	var frameTime int32    // Elapsed time to finish frame
	var frameDelay int32   // Delay time before next frame

	// Check frame lock
	if e.sc.FrameRateLock == 0 {
		e.sc.FrameRateLock = 60 // Common refresh rate of monitors
	}

	// Boot it up
	snapEngine := core.New(*e.sc, e.isDebug)
	if err := snapEngine.Init(); err != nil {
		return err
	}
	if err := snapEngine.LoadLevel(); err != nil {
		return err
	}

	frameDelay = int32(1000 / e.sc.FrameRateLock) // Max time between frames
	for snapEngine.IsRunning() {
		frameStart = snapEngine.DeltaTime()

		snapEngine.HandleEvents()
		snapEngine.HandleUpdate()
		if err := snapEngine.HandleRender(); err != nil {
			return err
		}

		snapEngine.FPS++
		frameTime = int32(snapEngine.DeltaTime() - frameStart)

		// Slow down render if the system can work too fast
		if frameDelay > frameTime {
			snapEngine.FPS = e.sc.FrameRateLock / uint16(frameTime + 1)
			snapEngine.Delay(uint32(frameDelay - frameTime))
		}
	}

	return nil
}
