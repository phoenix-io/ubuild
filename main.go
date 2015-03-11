package main

import (
	_"fmt"
	"github.com/codegangsta/cli"
	"os"
	"github.com/umbrellaops/ubuild/engine"
)

func main() {

	app := cli.NewApp()
	app.Name = "ubuild"
	app.Usage = "Tool to build Container(s)"
	app.Version = "0.1.0"
	app.Author = "Umbrella Ops"

	app.Commands = []cli.Command{
		{
			Name:  "verify",
			Usage: "Verify the json file",
			Action: func(c *cli.Context) {
				verifyJSON(c)
			},
		},
		{
			Name:  "test",
			Usage: "Mock the build procedure, but does not create any instance",
			Action: func(c *cli.Context) {
				testContainer(c)
			},
		},
		{
			Name:  "create",
			Usage: "Create the container at provided builder",
			Action: func(c *cli.Context) {
				createContainer(c)
			},
		},
	}

	app.Run(os.Args)
}

// Parse the file provided in cmd line argument.
// Do sanity check.
// Check for Builder and accordingly invoke Verify()
// using engine Interface
func verifyJSON(c *cli.Context) bool {
	var container engine.ContainerEngine
	container = &engine.ContainerConfig{Name:"Test Container"}
	container.Verify()
	return false
}


// Parse the file provided in cmd line argument.
// Do sanity check.
// Check for Builder and accordingly invoke Test()
// using engine Interface
func testContainer(c *cli.Context) bool {
	return false

}


// Parse the file provided in cmd line argument.
// Do sanity check.
// Check for Builder and accordingly invoke Create()
// using engine Interface
func createContainer(c *cli.Context) bool {
	return false
}
