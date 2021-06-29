package entity

// StaticObject not movable, not intractable just image on screen
type StaticObject struct {
	Position            *Position
	DrawableInformation *DrawableInformation
}

// NewStaticObject simple static object, can image, can be building
func NewStaticObject(position *Position, drawableInformation *DrawableInformation) *StaticObject {
	return &StaticObject{Position: position, DrawableInformation: drawableInformation}
}

// OnUpdate will do nothing, it's static object
func (s *StaticObject) OnUpdate() {}

// GetDrawableInformation about object
func (s *StaticObject) GetDrawableInformation() *DrawableInformation {
	return s.DrawableInformation
}

// GetPosition in the scene
func (s *StaticObject) GetPosition() *Position {
	return s.Position
}
