package engine

import (
	"runtime"

	"github.com/LinMAD/Snap/engine/core"
	"github.com/LinMAD/Snap/engine/entity"
	"github.com/LinMAD/Snap/engine/platform"
)

// Entry point for application
type Entry struct {
	config       *platform.ScreenConfiguration
	sceneObjects []entity.SceneObject
	isDebug      bool
}

func init() {
	runtime.LockOSThread()
}

// New engine entry point for application
func New(screenConfig *platform.ScreenConfiguration, isDebugMode bool) Entry {
	return Entry{
		config:  screenConfig,
		isDebug: isDebugMode,
	}
}

// LoadSceneObjects will be used during the scene
func (e *Entry) LoadSceneObjects(sceneObjects []entity.SceneObject) {
	e.sceneObjects = sceneObjects
}

// Run engine work
func (e *Entry) Run() error {
	var frameStart uint32 // Initial time of one frame
	var frameTime int32   // Elapsed time to finish frame
	var frameDelay int32  // SetDelay time before next frame

	// Check frame lock
	if e.config.FrameRateLock == 0 {
		e.config.FrameRateLock = 60 // Common refresh rate of monitors
	}

	// Boot it up
	snapEngine := core.New(*e.config, e.isDebug)
	if err := snapEngine.Init(); err != nil {
		return err
	}
	if err := snapEngine.LoadComponents(e.sceneObjects); err != nil {
		return err
	}

	frameDelay = int32(1000 / e.config.FrameRateLock) // Max time between frames
	for snapEngine.IsRunning() {
		frameStart = snapEngine.DeltaTime()

		snapEngine.HandleEvents()
		snapEngine.HandleUpdate()
		if err := snapEngine.HandleRender(); err != nil {
			return err
		}

		frameTime = int32(snapEngine.DeltaTime() - frameStart)

		// Slow down render if the system can work too fast
		if frameDelay > frameTime {
			snapEngine.SetFps(e.config.FrameRateLock / uint16(frameTime+1))
			snapEngine.SetDelay(uint32(frameDelay - frameTime))
		}
	}

	return nil
}
