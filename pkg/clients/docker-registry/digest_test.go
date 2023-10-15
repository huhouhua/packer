package docker_registry

import (
	"github.com/opencontainers/go-digest"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Digest(t *testing.T) {
	s := "sha256:1d45a2d73068df42938dc6a64649a79a49ca0609c8aa0a31011dab48176674f7"
	d, err := digest.Parse(s)
	require.NoError(t, err)
	require.NoError(t, d.Validate())
	t.Log(d.Encoded())
	t.Log(d.String())
}
