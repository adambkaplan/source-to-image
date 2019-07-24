package docker

import (
	"bytes"
	"fmt"
	"os"

	"k8s.io/klog"

	"github.com/openshift/source-to-image/pkg/util/cmd"
)

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
	stderr := &bytes.Buffer{}
	opts := cmd.CommandOpts{
		Stdout: nil,
		Stderr: stderr,
		Dir:    contextDir,
	}
	args := []string{"build", "-t", tag, "-f", dockerfile}
	if klog.V(2) {
		opts.Stdout = os.Stdout
	}
	args = append(args, ".")
	err := d.runner.RunWithOptions(opts, "docker", args...)
	if err != nil {
		err = fmt.Errorf("failed to build container image: %s", stderr.String())
		return err
	}

	return nil
}
