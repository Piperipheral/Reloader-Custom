package app

import (
	"fmt"
	"github.com/piperipheral/Reloader-Custom/internal/pkg/cmd"
)

// Run runs the command
func Run() error {
	fmt.Print("hello")
	cmda := cmd.NewReloaderCommand()
	return cmda.Execute()
}
