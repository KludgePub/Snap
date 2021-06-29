package factory

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/LinMAD/Snap/engine/entity"
	"github.com/LinMAD/Snap/engine/graphics/data"
)

// SpriteFactory produces sprites
type SpriteFactory struct {
	renderer *sdl.Renderer
	textures *data.TextureContainer
}

// NewSpriteFactory instance
func NewSpriteFactory(r *sdl.Renderer, t *data.TextureContainer) *SpriteFactory {
	return &SpriteFactory{renderer: r, textures: t}
}

// Draw new sprite
func (f *SpriteFactory) Draw(object entity.SceneObject, flip sdl.RendererFlip) error {
	drawInfo := object.GetDrawableInformation()
	if drawInfo.TextureData == nil {
		return nil
	}

	src := sdl.Rect{
		W: int32(drawInfo.Width),
		H: int32(drawInfo.Height),
	}

	object.GetPosition().Lock()
	defer object.GetPosition().Unlock()

	dst := sdl.Rect{
		X: object.GetPosition().X,
		Y: object.GetPosition().Y,
		W: int32(drawInfo.Width),
		H: int32(drawInfo.Height),
	}

	t, tErr := f.textures.Get(drawInfo.TextureData)
	if tErr != nil {
		return fmt.Errorf("unable to get texture data back: %s", tErr.Error())
	}

	drawInfo.Color.Lock()
	defer drawInfo.Color.Unlock()

	cErr := t.SetColorMod(
		drawInfo.Color.Red,
		drawInfo.Color.Green,
		drawInfo.Color.Blue,
	)
	if cErr != nil {
		return fmt.Errorf("cannot set color for texture: %s", tErr.Error())
	}

	if err := f.renderer.CopyEx(t, &src, &dst, 0, new(sdl.Point), flip); err != nil {
		return fmt.Errorf("unable copy texture to GPU: %s", err.Error())
	}

	return nil
}
