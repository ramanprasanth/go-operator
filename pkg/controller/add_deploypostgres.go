package controller

import (
	"github.com/talat-shaheen/postgres-go/pkg/controller/deploypostgres"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, deploypostgres.Add)
}
