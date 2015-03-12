package engine

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
	Run []string `yaml:"run"`
}

type BuilderS struct {
	Type string            `yaml:"type"`
	Attr map[string]string `yaml:"attr"`
}
