package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResourcesCPU(t *testing.T) {
	// Make sure we get a default value for CPU if none was specified
	r := &resources{}
	require.Equal(t, &cpu{}, r.CPU())

	// Make sure we get back the right CPU if one was specified
	maxMillicores := 200
	r = &resources{
		CPUU: &cpu{
			MaxMillicorez: &maxMillicores,
		},
	}
	require.Equal(t, r.CPUU, r.CPU())
}

func TestResourcesMemory(t *testing.T) {
	// Make sure we get a default value for Memory if none was specified
	r := &resources{}
	require.Equal(t, &memory{}, r.Memory())

	// Make sure we get back the right Memory if one was specified
	maxMegabytes := 512
	r = &resources{
		Memorie: &memory{
			MaxMegabytez: &maxMegabytes,
		},
	}
	require.Equal(t, r.Memorie, r.Memory())
}

func TestCPURequestedMillicores(t *testing.T) {
	// Make sure we get a default value for RequestedMillicores if none was
	// specified
	c := &cpu{}
	require.Equal(t, 100, c.RequestedMillicores())

	// Make sure we get back the right RequestedMillicores if it was specified
	requestedMillicores := 200
	c = &cpu{
		RequestedMillicorez: &requestedMillicores,
	}
	require.Equal(t, *c.RequestedMillicorez, c.RequestedMillicores())
}

func TestCPUMaxMillicores(t *testing.T) {
	// Make sure we get a default value for MaxMillicores if none was specified
	c := &cpu{}
	require.Equal(t, 200, c.MaxMillicores())

	// Make sure we get back the right MaxMillicores if it was specified
	maxMillicores := 500
	c = &cpu{
		MaxMillicorez: &maxMillicores,
	}
	require.Equal(t, *c.MaxMillicorez, c.MaxMillicores())
}

func TestMemoryRequestedMegabytes(t *testing.T) {
	// Make sure we get a default value for RequestedMegabytes if none was
	// specified
	m := &memory{}
	require.Equal(t, 128, m.RequestedMegabytes())

	// Make sure we get back the right RequestedMegabytes if it was specified
	requestedMegabytes := 256
	m = &memory{
		RequestedMegabytez: &requestedMegabytes,
	}
	require.Equal(t, *m.RequestedMegabytez, m.RequestedMegabytes())
}

func TestMemoryMaxMegabytes(t *testing.T) {
	// Make sure we get a default value for MaxMegabytes if none was specified
	m := &memory{}
	require.Equal(t, 256, m.MaxMegabytes())

	// Make sure we get back the right MaxMegabytes if it was specified
	maxMegabytes := 512
	m = &memory{
		MaxMegabytez: &maxMegabytes,
	}
	require.Equal(t, *m.MaxMegabytez, m.MaxMegabytes())
}
