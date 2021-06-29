package data

import (
	"fmt"
	"sync"

	"github.com/veandco/go-sdl2/ttf"
)

// FontContainer data storage
type FontContainer struct {
	// TODO Flyweight pattern re-use same fonts or add new
	Font *ttf.Font

	isLoaded bool
}

// FontData that will be used for in scene
type FontData struct {
	sync.Mutex

	// ID unique name
	ID string
	// FontFilePath loaded once and cannot be changed later
	FontFilePath string
	// Size of font
	Size uint32
}

// NewFontContainer initialise new font container
func NewFontContainer() (*FontContainer, error) {
	if err := ttf.Init(); err != nil {
		return nil, fmt.Errorf("unable to load font module: %s", err.Error())
	}

	return new(FontContainer), nil
}

// LoadFromFile a font
func (f *FontContainer) LoadFromFile(d *FontData) (err error) {
	if f.isLoaded || d == nil {
		return nil
	}

	f.Font, err = ttf.OpenFont(d.FontFilePath, int(d.Size))
	if err != nil {
		return fmt.Errorf(
			"unable to open font file(%s) error: %s",
			d.FontFilePath,
			err.Error(),
		)
	}

	f.isLoaded = true

	return nil
}

// Clear font data container
func (f *FontContainer) Clear() {
	f.Font.Close()
	ttf.Quit()
}
