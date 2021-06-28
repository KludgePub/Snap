package entity

// Static not movable, not intractable just image on screen
type Static struct {
	Position Position
	DrawableInformation DrawableInformation
	reachRight bool
}

// OnUpdate will do nothing, it's static object
func (s *Static) OnUpdate() {}

// GetDrawableInformation about object
func (s *Static) GetDrawableInformation() DrawableInformation {
	return s.DrawableInformation
}

// GetPosition in the scene
func (s *Static) GetPosition() Position {
	return s.Position
}
