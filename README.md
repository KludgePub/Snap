# Snap

> Simple educational 2D micro engine, designed with Golang and SDL2 bindings to allow a bit easier to build primitive 2D games or scenes to play with.

___

## Usage

Build command can be different on your OS but that shouldn't be difficult.
Since SDL2 can be compiled on any platform.

If you are using Linux Debian 10, then you can execute this:
```text
env CGO_ENABLED=1 CC=gcc GOOS=linux GOARCH=amd64 go build -tags -gcflags="all=-N -l" -o Snap --race main.go
```

Example of main file:
```go
package main

import (
	"github.com/LinMAD/Snap/engine"
	"github.com/LinMAD/Snap/engine/entity"
	"github.com/LinMAD/Snap/engine/graphics/data"
	"github.com/LinMAD/Snap/engine/platform"
)

func main() {
	isDebugMode := true
	actors := []entity.SceneObject{
		&entity.Static{
			Position: entity.Position{X: 150, Y: 50},
			DrawableInformation: entity.DrawableInformation{
				Width:  500,
				Height: 500,
				TextureData: data.TextureData{
					ImageFilePath: "assets/snap_engine_logo.png",
					ID:            "snap_logo",
				},
			},
		},
	}

	snapEngine := engine.New(new(platform.ScreenConfiguration), isDebugMode)
	snapEngine.LoadSceneObjects(actors)

	if err := snapEngine.Run(); err != nil {
		panic(err.Error())
	}
}
```

### Dependencies

Debian like:
- `apt install libsdl2{,-image,-gfx}-dev`

Darwin:
- `brew install sdl2{,_image,_gfx} pkg-config`

[More explanations can be found in bindings if needed](https://github.com/veandco/go-sdl2#requirements)

___
### License MIT

It was designed to create code challenge and workshop in a more fun way instead of writing TODO, or APIs. Feel free to
contribute, fork it or use it as you like, happy coding.

![SnapEngineLogo](assets/snap_engine_logo.png "Logo")
