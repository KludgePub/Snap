package platform

// ScreenConfiguration contains settings for window
type ScreenConfiguration struct {
	FrameRateLock uint16

	// Width of window
	Width uint
	// Height of window
	Height uint

	// Title of window
	Title string

	// IsResizeable window context
	IsResizeable bool
	// IsFullScreen window context
	IsFullScreen bool
}
