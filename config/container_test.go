package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainerName(t *testing.T) {
	container := &container{
		ContainerName: "foo",
	}
	require.Equal(t, container.ContainerName, container.Name())
}

func TestContainerImage(t *testing.T) {
	container := &container{
		Img: "debian:stretch",
	}
	require.Equal(t, container.Img, container.Image())
}

func TestContainerImagePullPolicy(t *testing.T) {
	c := &container{
		ImgPullPolicy: ImagePullPolicyAlways,
	}
	require.Equal(t, c.ImgPullPolicy, c.ImagePullPolicy())

	// Test default value
	c = &container{}
	require.Equal(t, ImagePullPolicyIfNotPresent, c.ImagePullPolicy())
}

func TestContainerEnvironment(t *testing.T) {
	container := &container{
		Env: []string{
			"FOO=bar",
			"BAT=baz",
		},
	}
	require.Equal(t, container.Env, container.Environment())
}

func TestContainerWorkingDirectory(t *testing.T) {
	container := &container{
		WorkDir: "/home/krancour",
	}
	require.Equal(t, container.WorkDir, container.WorkingDirectory())
}

func TestContainerCommand(t *testing.T) {
	container := &container{
		Cmd: []string{"echo"},
	}
	require.Equal(t, container.Cmd, container.Command())
}

func TestContainerArgs(t *testing.T) {
	container := &container{
		Arguments: []string{"foo"},
	}
	require.Equal(t, container.Arguments, container.Args())
}

func TestContainerTTY(t *testing.T) {
	container := &container{
		IsTTY: true,
	}
	require.Equal(t, container.IsTTY, container.TTY())
}

func TestContainerPrivileged(t *testing.T) {
	container := &container{
		IsPrivileged: true,
	}
	require.Equal(t, container.IsPrivileged, container.Privileged())
}

func TestContainerMountDockerSocket(t *testing.T) {
	container := &container{
		ShouldMountDockerSocket: true,
	}
	require.Equal(
		t,
		container.ShouldMountDockerSocket,
		container.MountDockerSocket(),
	)
}

func TestContainerSourceMountPath(t *testing.T) {
	container := &container{
		SrcMountPath: "/var/src/app",
	}
	require.Equal(t, container.SrcMountPath, container.SourceMountPath())
}

func TestContainerSharedStorageMountPath(t *testing.T) {
	container := &container{
		SharedStrgMountPath: "/var/shared",
	}
	require.Equal(
		t,
		container.SharedStrgMountPath,
		container.SharedStorageMountPath(),
	)
}

func TestContainerResources(t *testing.T) {
	// Make sure we get a default value for Resources if none was specified
	c := &container{}
	require.Equal(t, &resources{}, c.Resources())

	// Make sure we get back the right Resources if one was specified
	maxMillicores := 200
	maxMegabytes := 512
	c = &container{
		Resourcez: &resources{
			CPUU: &cpu{
				MaxMillicorez: &maxMillicores,
			},
			Memorie: &memory{
				MaxMegabytez: &maxMegabytes,
			},
		},
	}
}
