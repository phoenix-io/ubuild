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
	Name        string               `json:"name"`
	Builder     BuilderS             `json:"builder"`
	Attributes  ContainerAttributesS `json:"container-attrib"`
	Provisioner []ProvisionerS       `json:"provisioner"`
}

type ContainerAttributesS struct {
	Base    string   `json:"base-image"`
	Env     []string `json:"env"`
	WorkDir string   `json:"work-dir"`
	Cmd     string   `json:"cmd"`
}

type ProvisionerS struct {
	Type string   `json:"type"`
	Run  []string `json:"run"`
}

type BuilderS struct {
	Type string            `json:"type"`
	Attr map[string]string `json:"attr"`
}
