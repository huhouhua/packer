package main

import (
	"math/rand"
	"ruijie.com.cn/devops/packer/internal/packerserver"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	packerserver.NewApp().Run()
}
