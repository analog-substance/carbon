package image

import (
	builder "github.com/NoF0rte/cmd-builder"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"syscall"
)

type Image struct {
	Path string `json:"path"`
}

func (img *Image) Name() string {
	return filepath.Base(img.Path)
}

func (img *Image) ExecBuild() error {
	path, _ := exec.LookPath("ssh")
	args := []string{
		"build",
		".",
	}

	//args = append(args, additionalArgs...)
	if //goland:noinspection GoBoolExpressions
	runtime.GOOS == "windows" {
		return builder.Cmd(args[0], args[1:]...).Interactive().Run()
	}
	return syscall.Exec(path, args, os.Environ())
}

func ListBuilds() {

}

func ListImages() {

}
