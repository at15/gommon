package generator

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/noodle"
	"github.com/dyweb/gommon/util/fsutil"
	"github.com/dyweb/gommon/util/genutil"
)

// Generate walks all sub directories and generate files based on gommon.yml
func Generate(root string) error {
	var files []string
	// TODO: limit level
	ignores := DefaultIgnores()
	if fsutil.FileExists(filepath.Join(root, GommonIgnoreFile)) {
		userIgnores, err := fsutil.ReadIgnoreFile(filepath.Join(root, GommonIgnoreFile))
		if err != nil {
			return err
		}
		ignores = userIgnores
	}
	err := fsutil.Walk(root, ignores, func(path string, info os.FileInfo) {
		//log.Trace(path + "/" + info.Name())
		if info.Name() == GommonConfigFile {
			files = append(files, join(path, info.Name()))
		}
	})
	if err != nil {
		return err
	}
	for _, file := range files {
		if err := GenerateSingle(file); err != nil {
			return err
		}
	}
	return nil
}

// GenerateSingle generates based on a single gommon.yml
func GenerateSingle(file string) error {
	dir := filepath.Dir(file)
	segments := strings.Split(dir, string(os.PathSeparator))
	pkg := segments[len(segments)-1]
	cfg := NewConfigFile(pkg, file)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return errors.Wrap(err, "can't read config file")
	}
	// NOTE: not using Unmarshal strict so new binary still works with old config with deprecated fields
	if err = yaml.Unmarshal(b, &cfg); err != nil {
		return errors.Wrap(err, "can't decode config file as YAML")
	}
	if cfg.GoPackage != "" {
		pkg = cfg.GoPackage
	}

	// gommon
	var body bytes.Buffer
	for _, l := range cfg.Loggers {
		b, err := l.Render()
		if err != nil {
			return err
		}
		body.Write(b)
	}
	if body.Len() != 0 {
		var header bytes.Buffer
		header.WriteString(genutil.DefaultHeader(file))
		header.WriteString("package " + pkg + "\n\n")
		// FIXME: (piguo) the import is hard coded
		header.WriteString("import dlog \"github.com/dyweb/gommon/log\"")
		header.Write(body.Bytes())
		formatted, err := format.Source(header.Bytes())
		if err != nil {
			return errors.Wrap(err, "error format generated go code")
		}
		if fsutil.WriteFile(join(dir, DefaultGeneratedFile), formatted); err != nil {
			return errors.Wrap(err, "error write generated file to disk")
		}
		log.Debugf("generated %s from %s", join(dir, DefaultGeneratedFile), file)
	}

	// noodle
	dstIndex := make(map[string][]noodle.EmbedConfig)
	for _, cfg := range cfg.Noodles {
		// update src and dst because the cwd is different, user may write gommon.yaml in assets folder
		// but run gommon in project root, using os.Chdir will make the logic hard to parallel
		cfg.Src = join(dir, cfg.Src)
		cfg.Dst = join(dir, cfg.Dst)
		sameDst, ok := dstIndex[cfg.Dst]
		if !ok {
			dstIndex[cfg.Dst] = []noodle.EmbedConfig{cfg}
		} else {
			dstIndex[cfg.Dst] = append(sameDst, cfg)
		}
	}
	// all the config that has same dst will be generated together
	// TODO: maybe should have put this logic in noodle package ...
	for dst, cfgs := range dstIndex {
		b, err := noodle.GenerateEmbedBytes(cfgs)
		if err != nil {
			return errors.Wrap(err, "error generate assets bundle using noodle")
		}
		if err := fsutil.WriteFile(dst, b); err != nil {
			return errors.Wrap(err, "error write generated file to disk")
		}
		// TODO: log all the sources
		log.Debugf("noodle generated %s from %d folders", dst, len(cfgs))
	}

	// gotmpl
	for _, tpl := range cfg.GoTemplates {
		if err := tpl.Render(dir); err != nil {
			return err
		}
	}

	// shell
	for _, s := range cfg.Shells {
		if err := s.Render(dir); err != nil {
			return err
		}
	}

	return nil
}
