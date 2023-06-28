package main

import (
	"github.com/piperipheral/Reloader-Custom/internal/pkg/app"
	"os"
)

func main() {
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
