package docker

import (
	"bytes"
	"os"

	"github.com/openshift/source-to-image/pkg/util/cmd"
	"github.com/openshift/source-to-image/pkg/util/log"
)

var klog = log.StderrLog

// Docker is the s2i build provider for docker
type Docker struct {
	runner cmd.CommandRunner
}

// NewBuildProviderDocker returns the s2i build provider for Docker
func NewBuildProviderDocker() *Docker {
	return &Docker{
		runner: cmd.NewCommandRunner(),
	}
}

// BuildImage uses Docker to build the container image with the provided context dir and Dockerfile.
// The resulting image is then given the provided tag.
func (d *Docker) BuildImage(contextDir string, dockerfile string, tag string) error {
	klog.V(0).Infof("Building container image %s with docker", tag)
	klog.V(5).Infof("Tag: %s, Context: %s, Dockerfile: %s", tag, contextDir, dockerfile)
	opts := cmd.CommandOpts{
		Stderr: os.Stderr,
		Stdout: os.Stdout,
		Dir:    contextDir,
	}
	args := []string{"build", "-t", tag, "-f", dockerfile}
	if klog.Is(1) {
		opts.Stdout = &bytes.Buffer{}
		opts.Stderr = &bytes.Buffer{}
	}
	args = append(args, ".")

	return d.runner.RunWithOptions(opts, "docker", args...)
}
