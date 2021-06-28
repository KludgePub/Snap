package entity

import "github.com/LinMAD/SnapEngine/engine/graphics/data"

// SceneObject interface represents as actor in the scene, can be anything
type SceneObject interface {
	// OnUpdate event called on each tick to update state
	// Make actual data updates for object here
	OnUpdate()

	// TODO Input event

	// GetDrawableInformation about object
	GetDrawableInformation() DrawableInformation
	// GetPosition in the scene
	GetPosition() Position
}

// DrawableInformation asset data
type DrawableInformation struct {
	// Width of image
	Width uint32
	// Height of image
	Height uint32

	// IsFlipped image flipped horizontally ?
	IsFlipped bool

	// TextureData about image
	TextureData data.TextureData
}

// Position in screen
type Position struct {
	// X coordinate on the screen
	X int32
	// Y coordinate on the screen
	Y int32
}

