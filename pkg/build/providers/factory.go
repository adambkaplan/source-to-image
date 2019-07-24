package providers

import "fmt"

// BuildProvider is an interface to a container image building tool.
type BuildProvider interface {
	BuildImage(contextDir string, dockerfile string, tag string) error
}

// GetBuildProvider creates a BuildProvider instance for the given image building tool.
func GetBuildProvider(provider string) (BuildProvider, error) {
	return nil, fmt.Errorf("build provider %s is not supported", provider)
}
