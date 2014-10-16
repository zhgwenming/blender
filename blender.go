package main

import (
	"github.com/docker/docker/reexec"
	_ "github.com/zhgwenming/blender/docker"
)

func main() {
	if reexec.Init() {
		return
	}
	//dockercli.Main()
}
