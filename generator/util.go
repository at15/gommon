package generator

import (
	"io/ioutil"

	"path/filepath"
)

var DefaultIgnores ignores = []string{"testdata", "vendor", ".idea"}

type ignores []string

func (is *ignores) isIgnored(s string) bool {
	for _, i := range *is {
		if i == s {
			return true
		}
	}
	return false
}

// dfs to find all gommon files
// TODO: limit level
// TODO: bfs using recursion?
func Walk(root string, ignore ignores) []string {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Warn(err)
	}
	var gommonFiles []string
	for _, file := range files {
		name := file.Name()
		//log.Info(join(root, name))
		if file.IsDir() && !ignore.isIgnored(name) {
			gommonFiles = append(gommonFiles, Walk(join(root, name), ignore)...)
			continue
		}
		if name == "gommon.yml" {
			gommonFiles = append(gommonFiles, join(root, name))
		}
	}
	return gommonFiles
}

func join(s ...string) string {
	return filepath.Join(s...)
}
