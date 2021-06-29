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

// TextureData ...
type TextureData struct {
	// ImageFilePath which will be used on screen
	ImageFilePath string
	// ID unique name
	ID string
}

// NewTexturesContainer for loaded textures
func NewTexturesContainer(r *sdl.Renderer) *TextureContainer {
	return &TextureContainer{
		renderer: r,
		loaded:   map[string]*sdl.Texture{},
	}
}

// LoadFromFile texture
func (t *TextureContainer) LoadFromFile(d *TextureData) error {
	if d == nil {
		return nil
	}

	// TODO Validate path to file
	image, err := img.Load(d.ImageFilePath)
	if err != nil {
		return fmt.Errorf("failed to load texture from file (%s): %s\n", d.ImageFilePath, err)
	}

	texture, err := t.renderer.CreateTextureFromSurface(image)
	if err != nil {
		return fmt.Errorf("failed to create texture: %s\n", err.Error())
	}

	t.loaded[d.ID] = texture

	return nil
}

// Get texture by id
func (t *TextureContainer) Get(d *TextureData) (*sdl.Texture, error) {
	if texture, isFound := t.loaded[d.ID]; isFound {
		return texture, nil
	}

	return nil, fmt.Errorf("texture by id (%s) not found", d.ID)
}

// GetAll all loaded textures
func (t *TextureContainer) GetAll() map[string]*sdl.Texture {
	return t.loaded
}
