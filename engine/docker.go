package engine

import (
	"fmt"
	_ "log"
)


func (c *ContainerConfig) Verify() bool {
	fmt.Println("Inside Docker Verify")
	return true
}

func (c *ContainerConfig) Test() bool {
	fmt.Println("Inside Docker Test")
	return true
}

func (c *ContainerConfig) Create() bool {
	fmt.Println("Inside Docker Create")
	return true
}
