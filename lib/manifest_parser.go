// Package parser will have all functionality for
// managing container manifest.
package parser

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/umbrellaops/ubuild/engine"
	"gopkg.in/yaml.v2"
)

//LoadConfig loads the config
func LoadConfig(fname string) (*engine.ContainerConfig, error) {
	data, err := ioutil.ReadFile(fname)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("Cannot read config file :%v", err)
	}

	var c engine.ContainerConfig
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse config file : %v", err)
	}
	return &c, nil

}

//SaveConfig saves the config
func SaveConfig(fname string, c *engine.ContainerConfig) error {
	f, err := os.Create(fname + ".new")
	if err != nil {
		return fmt.Errorf("Cannot create config file: %v", err)
	}

	defer f.Close()
	defer os.Remove(fname + ".new")

	data, err := yaml.Marshal(c)
	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("Cannot write to config file: %v ", err)
	}

	f.Close()

	err = os.Rename(fname+".new", fname)
	if err != nil {
		return fmt.Errorf("Cannot rename temp config file : %v ", err)
	}

	return nil
}

func (c *engine.ContainerConfig) ConvertToDockerfile(fpath string) error {
	
}
