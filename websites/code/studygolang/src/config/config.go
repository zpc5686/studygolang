package config

import (
	"process"
	"path"
)

// 项目根目录
var ROOT string

func init() {
	binDir, err := process.ExecutableDir()
	if err != nil {
	    panic(err)
	}
	ROOT = path.Dir(binDir)
}

