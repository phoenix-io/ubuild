// Package parser will have all functionality for
// managing container manifest.
package parser

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Structural representation of config file for genric container
type ContainerConfig struct {
	Name        string               `yaml:"name"`
	Builder     BuilderS             `yaml:"builder"`
	Attributes  ContainerAttributesS `yaml:"container-attrib"`
	Provisioner []ProvisionerS       `yaml:"provisioner"`
}

type ContainerAttributesS struct {
	Base    string   `yaml:"base-image"`
	Env     []string `yaml:"env"`
	WorkDir string   `yaml:"work-dir"`
	Cmd     string   `yaml:"cmd"`
}

type ProvisionerS struct {
	Type string   `yaml:"type"`
	Run  []string `yaml:"run"`
}

type BuilderS struct {
	Type string            `yaml:"type"`
	Attr map[string]string `yaml:"attr"`
}

func dumpConfig(config []byte) {
	fmt.Printf("\n%s\n", string(config))
}


//LoadConfig loads the config
func LoadConfig(fname string) (*ContainerConfig, error) {
	data, err := ioutil.ReadFile(fname)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("Cannot read config file :%v", err)
	}

	var c ContainerConfig
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse config file : %v", err)
	}

	return &c, nil

}

//SaveConfig saves the config
func SaveConfig(fname string, c *ContainerConfig) error {
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

	dumpConfig(data)

	err = os.Rename(fname+".new", fname)
	if err != nil {
		return fmt.Errorf("Cannot rename temp config file : %v ", err)
	}

	return nil
}

//ConvertToDockerfile will comvert ubuild yml configuration into a Dockerfile.
func (c *ContainerConfig) ConvertToDockerfile(fpath string) error {

	return fmt.Errorf("NOT IMPLEMENTED..")
}
