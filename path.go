package logger

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

// defaultLogPath
func defaultLogPath(path string) string {
	var err error
	var p string
	p, err = getCurrentExecDir()
	if err != nil {
		p, err = os.UserHomeDir()
		if err != nil {
			panic("can't find home path")
		}
	}

	logPath := p + "/" + path

	afs := afero.NewOsFs()
	check, _ := afero.DirExists(afs, logPath)
	if !check {
		err = afs.MkdirAll(logPath, 0o755)
		if err != nil {
			return p
			// panic("can't make path" + logPath)
		}
	}

	return logPath
}

// getCurrentExecDir

func getCurrentExecDir() (dir string, err error) {
	var path string
	path, err = exec.LookPath(os.Args[0])
	if err != nil {
		// fmt.Printf("exec.LookPath(%s), err: %s\n", os.Args[0], err)
		return "", err
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		// fmt.Printf("filepath.Abs(%s), err: %s\n", path, err)
		return "", err
	}
	dir = filepath.Dir(absPath)
	return dir, nil
}

// StrBuilder strings builder
func StrBuilder(args ...string) string {
	if len(args) == 0 {
		return ""
	}
	var str strings.Builder

	for _, v := range args {
		str.WriteString(v)
	}
	return str.String()
}
