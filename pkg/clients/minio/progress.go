package minio

import (
	"github.com/emirpasic/gods/sets/hashset"
	"ruijie.com.cn/devops/packer/pkg/logx"
)

type ProgressPar struct {
	percentDone *hashset.Set
	increment   int64
	FileSize    int64
}

func NewProgressPar(fileSize int64) *ProgressPar {
	return &ProgressPar{
		percentDone: hashset.New(),
		increment:   0,
		FileSize:    fileSize}
}

func (p2 *ProgressPar) Read(p []byte) (n int, err error) {
	p2.increment += int64(len(p))
	percent := (int)(p2.increment * 100 / p2.FileSize)
	if p2.percentDone.Contains(percent) {
		return
	}
	p2.print(percent)
	p2.percentDone.Add(percent)
	return
}
func (p2 *ProgressPar) print(done int) {
	switch done {
	case 1, 3, 5, 10, 20, 30, 40, 50, 60, 70, 80, 90, 96, 97, 98, 99, 100:
		logx.Trace("%d%s uploading....", done, "%")
		break
	default:
		break
	}
}
