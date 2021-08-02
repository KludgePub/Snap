package engine

import (
	"github.com/LinMAD/Snap/engine/logger"
	"runtime"

	"github.com/LinMAD/Snap/engine/core"
	"github.com/LinMAD/Snap/engine/entity"
	"github.com/LinMAD/Snap/engine/platform"
)

var log *logger.Logger

// Entry point for application
type Entry struct {
	config       *platform.ScreenConfiguration
	sceneObjects []entity.SceneObject

	// frameDelay time before next frame
	frameDelay uint32

	isDebug bool
}

func init() {
	runtime.LockOSThread()
}

// New engine entry point for application
func New(screenConfig *platform.ScreenConfiguration, isDebugMode bool) Entry {
	entry := Entry{
		config:  screenConfig,
		isDebug: isDebugMode,
	}

	// Check frame lock
	if entry.config.FrameRateLock == 0 {
		entry.config.FrameRateLock = 60 // Common refresh rate of monitors
	}

	// Max time between frames
	entry.frameDelay = uint32(1000 / entry.config.FrameRateLock)

	return entry
}

// LoadSceneObjects will be used during the scene
func (e *Entry) LoadSceneObjects(sceneObjects []entity.SceneObject) {
	e.sceneObjects = sceneObjects
}

// Run engine work
func (e *Entry) Run() error {
	var frameStart uint32 // Initial time of one frame
	var frameTime uint32  // Elapsed time to finish frame

	// Boot it up
	log.Log("Initializing Snap engine...")
	snapEngine := core.New(*e.config, e.isDebug)
	if err := snapEngine.Init(); err != nil {
		return err
	}

	log.Log("Loading components...")
	if err := snapEngine.LoadComponents(e.sceneObjects); err != nil {
		return err
	}

	log.Log("Executing...")
	for snapEngine.IsRunning() {
		frameStart = snapEngine.DeltaTime()

		snapEngine.HandleEvents()
		snapEngine.HandleUpdate()
		if err := snapEngine.HandleRender(); err != nil {
			return err
		}

		// Slow down render if the system can work too fast
		frameTime = snapEngine.DeltaTime() - frameStart
		if e.frameDelay > frameTime {
			snapEngine.SetDelay(e.frameDelay - frameTime)
		}
	}

	return nil
}
