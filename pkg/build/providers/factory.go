package providers

import (
	"fmt"

	"github.com/openshift/source-to-image/pkg/build/providers/docker"
	"github.com/openshift/source-to-image/pkg/build/providers/podman"
)

// BuildProvider is an interface to a container image building tool.
type BuildProvider interface {
	BuildImage(contextDir string, dockerfile string, tag string) error
}

// GetBuildProvider creates a BuildProvider instance for the given image building tool.
func GetBuildProvider(provider string) (BuildProvider, error) {
	switch provider {
	case "docker":
		return docker.NewBuildProviderDocker(), nil
	case "podman":
		return podman.NewBuildProviderPodman(), nil
	}
	return nil, fmt.Errorf("build provider %s is not supported", provider)
}
