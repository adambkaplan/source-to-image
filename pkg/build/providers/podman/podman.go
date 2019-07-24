package podman

import (
	"bytes"
	"os"

	"k8s.io/klog"

	"github.com/openshift/source-to-image/pkg/util/cmd"
)

// Podman is the s2i build provider for Podman
type Podman struct {
	runner cmd.CommandRunner
}

// NewBuildProviderPodman returns the s2i build provider for Podman
func NewBuildProviderPodman() *Podman {
	return &Podman{
		runner: cmd.NewCommandRunner(),
	}
}

// BuildImage uses podman to build the container image with the provided context dir and Dockerfile.
// The resulting image is then given the provided tag.
func (p *Podman) BuildImage(contextDir string, dockerfile string, tag string) error {
	klog.V(0).Infof("Building container image %s with podman", tag)
	klog.V(5).Infof("Tag: %s, Context: %s, Dockerfile: %s", tag, contextDir, dockerfile)
	opts := cmd.CommandOpts{
		Stderr: os.Stderr,
		Stdout: os.Stdout,
		Dir:    contextDir,
	}
	args := []string{"build", "-t", tag, "-f", dockerfile}
	if !klog.V(1) {
		opts.Stdout = &bytes.Buffer{}
		opts.Stderr = &bytes.Buffer{}
		// args = append(args, "-q")
	}
	args = append(args, ".")

	return p.runner.RunWithOptions(opts, "podman", args...)
}
