package iai

import "github.com/spacepal/Spacepal-AI/internal/model/imodel"

// Handler is an interface of TurnHandler
type Handler interface {
	Handle(in imodel.InGetter)
	Start()
}
