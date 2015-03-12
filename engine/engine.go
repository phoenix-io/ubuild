package engine

import "github.com/umbrellaops/ubuild/lib"

// A ContainerEngine is engine object
//
// This engine will be responsible to perform various
// functions
type ContainerEngine interface {
	//Verify the JSON file, it can be used to create
	//Container on given engine
	Verify() bool

	//Test run mock build, without creating any container
	Test() bool

	//Create will create the container on given backend.
	Create() bool
}

// Json representation of config file for genric container
type Container struct {
	ID     string
	Config *parser.ContainerConfig
}
