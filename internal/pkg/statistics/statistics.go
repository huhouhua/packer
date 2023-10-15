package statistics

import (
	docker_registry "ruijie.com.cn/devops/packer/pkg/clients/docker-registry"
	"ruijie.com.cn/devops/packer/pkg/logx"
	"ruijie.com.cn/devops/packer/pkg/util"
)

type Statistics struct {
	manifestMap map[int]ValuePair
	IsSuccess   bool
}
type ValuePair struct {
	Manifest   docker_registry.Manifest
	Statistics *Manifest
}

func NewValuePair(manifest docker_registry.Manifest, statistics *Manifest) ValuePair {
	return ValuePair{Manifest: manifest, Statistics: statistics}
}

func NewStatistics() *Statistics {
	return &Statistics{
		manifestMap: make(map[int]ValuePair),
		IsSuccess:   true,
	}
}
func (s *Statistics) Add(i int, value ValuePair) {
	if _, exist := s.manifestMap[i]; exist {
		return
	}
	s.manifestMap[i] = value
}
func (s *Statistics) Print() {
	for _, v := range s.manifestMap {
		Print(v.Manifest, "upload statistics ", false)
		skip := len(v.Statistics.Skip)
		fail := len(v.Statistics.Fail)

		if skip == len(v.Manifest.Layers)+1 {
			logx.InfoWithGreen("success:%d", skip)
		} else {
			logx.InfoWithGreen("success:%d", len(v.Statistics.Success))
		}
		logx.InfoWithMagenta("skip:%d", skip)
		logx.InfoWithRed("fail:%d", fail)
		if fail > 0 {
			s.IsSuccess = false
		}
	}
}
func sum(layers []docker_registry.Layer) int64 {
	var sum int64 = 0
	for i := 0; i < len(layers); i++ {
		sum += layers[i].Size
	}
	return sum
}
func Print(manifest docker_registry.Manifest, sub string, end bool) {
	logx.Info("====================== %s =======================", sub)
	logx.InfoWithCyan("schemaVersion:%d", manifest.SchemaVersion)
	logx.InfoWithCyan("externalRegistry:%s", manifest.RepoTag.ExternalRegistry)
	logx.InfoWithCyan("repo:%s/%s", manifest.RepoTag.RepoName, manifest.RepoTag.Tag)
	logx.InfoWithCyan("layerCount: %d  total: %s", len(manifest.Layers), util.FileSizeToFormat(sum(manifest.Layers)))
	logx.InfoWithCyan("configCount: 1 total: %s", util.FileSizeToFormat(manifest.Config.Size))
	if end {
		logx.Info("================================================================")
	}
}
