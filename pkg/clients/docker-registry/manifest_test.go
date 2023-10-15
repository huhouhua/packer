package docker_registry

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_manifes_v2(t *testing.T) {
	hub, err := New(_url, username, password, func(format string, args ...interface{}) {
		t.Logf(format, args)
	})
	require.NoError(t, err)
	manifest, err := hub.ManifestV2("huhouhua/redis", "6.2.4-alpine3.13")
	assert.NoError(t, err)
	assert.NotNil(t, manifest)
	_, a, err := manifest.Payload()
	assert.NoError(t, err)
	t.Logf("%s", string(a))
}

func Test_manifes(t *testing.T) {
	hub, err := New(_url, username, password, func(format string, args ...interface{}) {
		t.Logf(format, args)
	})
	require.NoError(t, err)
	manifest, err := hub.Manifest("huhouhua/redis", "6.2.4-alpine3.13")

	assert.NoError(t, err)
	assert.NotNil(t, manifest)
	_, a, err := manifest.Payload()

	assert.NoError(t, err)
	t.Logf("%s", string(a))
}

func Test_manifes_digest(t *testing.T) {
	hub, err := New(_url, username, password, func(format string, args ...interface{}) {
		t.Logf(format, args)
	})
	require.NoError(t, err)
	digest, err := hub.ManifestDigest("huhouhua/redis", "6.2.4-alpine3.13")
	assert.NoError(t, err)
	assert.NotNil(t, digest)
	t.Logf("%s", digest.String())
}
func Test_manifest_2(t *testing.T) {
	hub, err := New(_url, "", "", func(format string, args ...interface{}) {
		t.Logf(format, args)
	})
	require.NoError(t, err)
	manifest, err := hub.Manifest("huhouhua/redis", "6.2.4-alpine3.13")

	assert.NoError(t, err)
	assert.NotNil(t, manifest)
	_, a, err := manifest.Payload()

	assert.NoError(t, err)
	t.Logf("%s", string(a))
}
