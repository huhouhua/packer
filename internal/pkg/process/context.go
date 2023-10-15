package process

import (
	"ruijie.com.cn/devops/packer/internal/pkg/statistics"
	docker_registry "ruijie.com.cn/devops/packer/pkg/clients/docker-registry"
	"ruijie.com.cn/devops/packer/pkg/clients/minio"
)

type Context struct {
	ManifestsPath   string
	Bucket          string
	Environment     string
	Platform        string
	MinioClientFunc minio.ClientFunc
}

type ManifestContext struct {
	ManifestsPath   string
	Environment     string
	Platform        string
	Manifests       docker_registry.ManifestSlice
	MinioClientFunc minio.ClientFunc
}

type PackageContext struct {
	Statistics      *statistics.Manifest
	Layer           docker_registry.Layer
	Repo            docker_registry.RepoTag
	Environment     string
	Platform        string
	MinioClientFunc minio.ClientFunc
	Registry        *docker_registry.Registry
}

func NewManifestContext(platform string, minioClientFunc minio.ClientFunc, manifestsPath string, environment string) *ManifestContext {
	return &ManifestContext{Platform: platform, MinioClientFunc: minioClientFunc, ManifestsPath: manifestsPath, Environment: environment, Manifests: docker_registry.ManifestSlice{}}
}

func (m *ManifestContext) UnmarshalWithManifest(data []byte) error {
	return m.Manifests.UnmarshalJSON(data)
}
