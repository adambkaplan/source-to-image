package podman

import (
	"bytes"
	"fmt"
	"io"
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
	stderr := &bytes.Buffer{}
	opts := cmd.CommandOpts{
		Stdout: nil,
		Stderr: stderr,
		Dir:    contextDir,
	}
	args := []string{"build", "-t", tag, "-f", dockerfile}

	if klog.V(2) {
		opts.Stdout = os.Stdout
		// hack around https://github.com/containers/libpod/issues/3642
		// stderr contains useful information, such as the current STEP
		// Using multiwriter to ensure STEP appears in combined output
		opts.Stderr = io.MultiWriter(stderr, os.Stderr)
	}
	args = append(args, ".")

	err := p.runner.RunWithOptions(opts, "podman", args...)
	if err != nil {
		err = fmt.Errorf("failed to build container image:\n%s", stderr.String())
		return err
	}

	return nil
}
