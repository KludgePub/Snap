package factory

import (
	"github.com/LinMAD/SnapEngine/engine/entity"
	"github.com/LinMAD/SnapEngine/engine/graphics/data"
	"github.com/veandco/go-sdl2/sdl"
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
// textureID in texture container
// x, y coordinates on screen
// w - width, h - height of texture
func (f *SpriteFactory) Draw(object entity.SceneObject, flip sdl.RendererFlip) error {
	src := sdl.Rect{
		W: int32(object.GetDrawableInformation().Width),
		H: int32(object.GetDrawableInformation().Height),
	}
	dst := sdl.Rect{
		X: object.GetPosition().X,
		Y: object.GetPosition().Y,
		W: int32(object.GetDrawableInformation().Width),
		H: int32(object.GetDrawableInformation().Height),
	}

	t, tErr := f.textures.Get(object.GetDrawableInformation().TextureData)
	if tErr != nil {
		return tErr
	}

	if err := f.renderer.CopyEx(t, &src, &dst, 0, new(sdl.Point), flip); err != nil {
		return err
	}

	return nil
}
