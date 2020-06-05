package controller

import (
	"github.com/ramanprasanth/go-operator/pkg/controller/deploypostgres"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, deploypostgres.Add)
}
