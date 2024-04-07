// A generated module for DockerfileSecrets functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/dockerfile-secrets/internal/dagger"
)

type DockerfileSecrets struct{}

func (m *DockerfileSecrets) Build(
	ctx context.Context,
	// Build context
	buildContext *Directory,
	// Dockerfile path
	dockerfile *File,
	// Secret file path
	secretFile *Secret,
) *Container {
	return buildContext.
		WithFile("Dockerfile", dockerfile).
		DockerBuild(dagger.DirectoryDockerBuildOpts{
			Secrets: []*Secret{secretFile},
		})
}
