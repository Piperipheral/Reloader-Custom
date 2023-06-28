package app

import (
	"github.com/piperipheral/Reloader-Custom/internal/pkg/cmd"
)

// Run runs the command
func Run() error {
	cmda := cmd.NewReloaderCommand()
	return cmda.Execute()
}
