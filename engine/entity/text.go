package entity

import "github.com/LinMAD/Snap/engine/graphics/data"

// TextObject simple letter object
type TextObject struct {
	textField string

	Position            *Position
	DrawableInformation *DrawableInformation
}

// NewTextObject to show in scene
func NewTextObject(position *Position, font *data.FontData, color *Color) *TextObject {
	return &TextObject{
		Position: position,
		DrawableInformation: &DrawableInformation{
			FontData: font,
			Color:    color,
			Text:     &Text{},
		},
	}
}

// SetTextField of object
func (t *TextObject) SetTextField(textField string) {
	t.DrawableInformation.Text.Lock()
	defer t.DrawableInformation.Text.Unlock()

	t.textField = textField
}

// OnUpdate update text
func (t *TextObject) OnUpdate() {
	t.DrawableInformation.Text.Lock()
	defer t.DrawableInformation.Text.Unlock()

	t.DrawableInformation.Text.TextToPrint = t.textField
}

// GetDrawableInformation about object
func (t *TextObject) GetDrawableInformation() *DrawableInformation {
	return t.DrawableInformation
}

// GetPosition in the scene
func (t *TextObject) GetPosition() *Position {
	return t.Position
}
