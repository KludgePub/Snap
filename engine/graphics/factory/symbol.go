package factory

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/LinMAD/Snap/engine/entity"
	"github.com/LinMAD/Snap/engine/graphics/data"
)

// SymbolFactory responsible for generating text
type SymbolFactory struct {
	fontContainer *data.FontContainer
	renderer      *sdl.Renderer

	textResources map[string]*resource
}

type resource struct {
	surface *sdl.Surface
	texture *sdl.Texture
}

// NewSymbolFactory instance
func NewSymbolFactory(f *data.FontContainer, r *sdl.Renderer) *SymbolFactory {
	return &SymbolFactory{
		fontContainer: f,
		renderer:      r,
		textResources: make(map[string]*resource, 0),
	}
}

// Draw new text surface
func (f *SymbolFactory) Draw(object entity.SceneObject) (err error) {
	drawInfo := object.GetDrawableInformation()

	if drawInfo.FontData == nil || drawInfo.Text == nil || len(drawInfo.Text.TextToPrint) == 0 {
		return nil
	}
	// TODO This is not optimal way to handle font on the screen

	var isLocatedInList bool
	var textResource *resource

	// refresh resource or create new one
	if textResource, isLocatedInList = f.textResources[object.GetDrawableInformation().FontData.ID]; isLocatedInList {
		textResource.surface.Free()
		if errDestroy := textResource.texture.Destroy(); errDestroy != nil {
			return fmt.Errorf("cannot remove font texture: %s", errDestroy.Error())
		}
	} else {
		textResource = &resource{}
		f.textResources[object.GetDrawableInformation().FontData.ID] = textResource
	}

	drawInfo.Color.Lock()
	defer drawInfo.Color.Unlock()

	// for text transparency enough to use RGB, 3x 0 will be transparent and 255 visible
	color := sdl.Color{
		R: drawInfo.Color.Red,
		G: drawInfo.Color.Green,
		B: drawInfo.Color.Blue,
		A: 255,
	}

	drawInfo.Text.Lock()
	defer drawInfo.Text.Unlock()

	textResource.surface, err = f.fontContainer.Font.RenderUTF8Blended(drawInfo.Text.TextToPrint, color)
	if err != nil {
		return fmt.Errorf("cannot create font surface: %s", err.Error())
	}

	object.GetPosition().Lock()
	defer object.GetPosition().Unlock()

	dst := &sdl.Rect{
		X: object.GetPosition().X,
		Y: object.GetPosition().Y,
		W: textResource.surface.W,
		H: textResource.surface.H,
	}

	textResource.texture, err = f.renderer.CreateTextureFromSurface(textResource.surface)
	if err != nil {
		return fmt.Errorf("cannot create font texture from surface: %s", err.Error())
	}

	err = f.renderer.CopyEx(textResource.texture, nil, dst, 0, new(sdl.Point), sdl.FLIP_NONE)
	if err != nil {
		return fmt.Errorf("unable copy font texture to GPU: %s", err.Error())
	}

	return nil
}
