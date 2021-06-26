package data

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// TextureContainer data storage
type TextureContainer struct {
	renderer *sdl.Renderer
	loaded   map[string]*sdl.Texture
}

// NewTexturesContainer for loaded textures
func NewTexturesContainer(r *sdl.Renderer) *TextureContainer {
	return &TextureContainer{
		renderer: r,
		loaded:   map[string]*sdl.Texture{},
	}
}

// LoadFromFile texture
func (t *TextureContainer) LoadFromFile(path string, id string) error {
	image, err := img.Load(path)
	if err != nil {
		return fmt.Errorf("failed to load texture from file (%s): %s\n", path, err)
	}

	texture, err := t.renderer.CreateTextureFromSurface(image)
	if err != nil {
		return fmt.Errorf("failed to create texture: %s\n", err.Error())
	}

	t.loaded[id] = texture

	return nil
}

// Get texture by id
func (t *TextureContainer) Get(id string) (*sdl.Texture, error) {
	if texture, isFound := t.loaded[id]; isFound {
		return texture, nil
	}

	return nil, fmt.Errorf("texture by id (%s) not found", id)
}

// GetAll all loaded textures
func (t *TextureContainer) GetAll() map[string]*sdl.Texture {
	return t.loaded
}
