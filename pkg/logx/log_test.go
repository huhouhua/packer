package logx

import (
	"testing"
)

func Test_Log(t *testing.T) {
	Trace("trace")
	Success("%s", "success")
}
