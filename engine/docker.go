package engine

import (
	"fmt"
	_ "log"
)

func (c *Container) Verify() bool {
	fmt.Println("Inside Docker Verify")
	return true
}

func (c *Container) Test() bool {
	fmt.Println("Inside Docker Test")
	return true
}

func (c *Container) Create() bool {
	fmt.Println("Inside Docker Create")
	return true
}
