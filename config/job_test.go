package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJobName(t *testing.T) {
	job := &job{
		name: "foo",
	}
	require.Equal(t, job.name, job.Name())
}

func TestJobPrimaryContainer(t *testing.T) {
	job := &job{
		primaryContainer: &container{
			ContainerName: "foo",
		},
	}
	require.Equal(t, job.primaryContainer, job.PrimaryContainer())
}

func TestJobSicecarContainers(t *testing.T) {
	job := &job{
		sidecarContainers: []Container{
			&container{
				ContainerName: "foo",
			},
			&container{
				ContainerName: "bar",
			},
		},
	}
	require.Equal(t, job.sidecarContainers, job.SidecarContainers())
}

func TestJobSourceMountMode(t *testing.T) {
	job := &job{
		sourceMountMode: SourceMountModeReadOnly,
	}
	require.Equal(t, job.sourceMountMode, job.SourceMountMode())
}
