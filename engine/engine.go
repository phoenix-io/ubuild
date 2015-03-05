package engine

// A ContainerEngine is engine object
//
// This engine will be responsible to perform various
// functions
type ContainerEngine interface {
	//Verify the JSON file, it can be used to create
	//Container on given engine
	Verify()

	//Test run mock build, without creating any container
	Test()

	//Create will create the container on given backend.
	Create()
}
