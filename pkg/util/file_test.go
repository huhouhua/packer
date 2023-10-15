package util

import (
	"github.com/stretchr/testify/require"
	"os"
	docker_registry "ruijie.com.cn/devops/packer/pkg/clients/docker-registry"
	"testing"
)

func Test_SizeFormat(t *testing.T) {
	var size int64 = 198324622
	format := FileSizeToFormat(size)
	require.Equal(t, format, "189.14 MB")
}
func Test_Read(t *testing.T) {
	data, err := os.ReadFile("manifest.json")
	require.NoError(t, err)

	var manifest docker_registry.ManifestSlice
	err = manifest.UnmarshalJSON(data)

}
