package util

import "fmt"

func FileSizeToFormat(size int64) string {
	format := []string{"B", "KB", "MB", "GB", "TB"}
	slen := float32(size)
	order := 0
	for slen >= 1024 && order < len(format)-1 {
		order++
		slen = slen / 1024
	}
	return fmt.Sprintf("%.2f %s", slen, format[order])
}
