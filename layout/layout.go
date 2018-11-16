package layout

// ProjectLayout provides set of functions for describing C++ project
type ProjectLayout interface {
	IsLibrary(root string) (bool, error)
	IsApplication(root string) (bool, error)
}
