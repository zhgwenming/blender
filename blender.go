package main

import (
	version "github.com/docker/docker/dockerversion"
	"github.com/docker/docker/reexec"
	_ "github.com/zhgwenming/blender/docker"
	"os"
)

func main() {
	version.IAMSTATIC = true

	// bridge "blender docker" to docker
	if len(os.Args) > 1 && os.Args[1] == "docker" {
		os.Args = os.Args[1:]
	}

	if reexec.Init() {
		return
	}
}
