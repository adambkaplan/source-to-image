package providers

import (
	"fmt"
	"runtime"

	"github.com/openshift/source-to-image/pkg/api"
	"github.com/openshift/source-to-image/pkg/build/providers/buildah"
	"github.com/openshift/source-to-image/pkg/build/providers/docker"
	"github.com/openshift/source-to-image/pkg/build/providers/podman"
	"github.com/openshift/source-to-image/pkg/util/cmd"
)

// BuildProvider is an interface to a container image building tool.
type BuildProvider interface {
	BuildImage(contextDir string, dockerfile string, tag string) error
}

// GetBuildProvider creates a BuildProvider instance for the given image building tool.
func GetBuildProvider(provider api.ProviderRuntime) (BuildProvider, error) {
	switch provider {
	case api.BuildahRuntime:
		return buildah.NewBuildProviderBuildah(), nil
	case api.DockerRuntime:
		return docker.NewBuildProviderDocker(), nil
	case api.PodmanRuntime:
		return podman.NewBuildProviderPodman(), nil
	}
	return nil, fmt.Errorf("build provider %s is not supported", provider)
}

// FindProviderRuntime searches the local environment for an appropriate build provider runtime.
func FindProviderRuntime() (api.ProviderRuntime, error) {
	if runtime.GOOS == "windows" {
		return api.DockerRuntime, nil
	}
	exec := cmd.NewCommandRunner()
	searchList := []api.ProviderRuntime{}
	if runtime.GOOS == "darwin" {
		searchList = append(searchList, api.PodmanRuntime, api.DockerRuntime)
	}
	if runtime.GOOS == "linux" {
		searchList = append(searchList, api.PodmanRuntime, api.BuildahRuntime, api.DockerRuntime)
	}
	for _, provider := range searchList {
		err := runWhich(provider, exec)
		if err == nil {
			return provider, nil
		}
	}
	return "", fmt.Errorf("could not find a build provider - searched %s", searchList)
}

func runWhich(provider api.ProviderRuntime, runner cmd.CommandRunner) error {
	return runner.Run("which", string(provider))
}
