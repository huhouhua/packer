package statistics

import docker_registry "ruijie.com.cn/devops/packer/pkg/clients/docker-registry"

type Manifest struct {
	Success []docker_registry.Layer
	Fail    []docker_registry.Layer
	Skip    []docker_registry.Layer
}

func NewManifest() *Manifest {
	return &Manifest{
		Success: []docker_registry.Layer{},
		Fail:    []docker_registry.Layer{},
		Skip:    []docker_registry.Layer{},
	}
}
