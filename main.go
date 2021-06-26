package main

import (
	"fmt"
	"os"

	"github.com/LinMAD/SnapEngine/engine"
	"github.com/LinMAD/SnapEngine/engine/platform"
)

func main() {
	snapEngine := engine.New(new(platform.ScreenConfiguration), true)
	if err := snapEngine.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
