package main

import (
	"os"
	"path/filepath"
	"strings"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	basedir := os.Args[1]
	query := os.Args[2]
	dirs, err := filepath.Glob(basedir + "/*/*/*")
	if err != nil {
		wf.FatalError(err)
	}
	for _, dir := range dirs {
		content := strings.Replace(dir, basedir+"/", "", 1)
		arr := strings.Split(content, "/")
		host, owner, repo := arr[0], arr[1], arr[2]
		wf.NewItem(content).Arg(basedir, host, owner, repo).Valid(true)
	}
	wf.Filter(query)
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
