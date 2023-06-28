package handler

import (
	"github.com/piperipheral/Reloader-Custom/internal/pkg/util"
)

// ResourceHandler handles the creation and update of resources
type ResourceHandler interface {
	Handle() error
	GetConfig() (util.Config, string)
}
