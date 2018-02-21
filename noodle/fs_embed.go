package noodle

import (
	"archive/zip"
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/dyweb/gommon/util/fsutil"
	"github.com/pkg/errors"
)

type EmbedBox struct {
	Dirs  map[string]EmbedDir
	Files map[string]EmbedFile
}

type EmbedFile struct {
	FileInfo
	Data []byte
}

type EmbedDir struct {
	FileInfo
	Entries []FileInfo
}

var _ os.FileInfo = (*FileInfo)(nil)

// the File prefix is to export the field but avoid conflict with os.FileInfo interface ...
type FileInfo struct {
	FileName    string
	FileSize    int64
	FileMode    os.FileMode
	FileModTime time.Time
	FileIsDir   bool
}

func GenerateEmbed(root string) error {
	var (
		err     error
		ignores *fsutil.Ignores
		dirs    = make(map[string]*EmbedDir)
		files   = make(map[string][]*EmbedFile)
		lastErr error
	)
	if rootStat, err := os.Stat(root); err != nil {
		return errors.Wrap(err, "can't get stat of root folder")
	} else {
		log.Infof("root %s rootStat name %s", root, rootStat.Name())
		dirs[root] = newEmbedDir(rootStat)
	}
	if ignores, err = readIgnoreFile(root); err != nil {
		return err
	}
	fsutil.Walk(root, ignores, func(path string, info os.FileInfo) {
		//log.Info(path)
		// TODO: allow config if ignore file should be included
		//if info.Name() == ignoreFileName {
		//	return
		//}
		if info.IsDir() {
			log.Infof("add %s to path %s", info.Name(), path)
			dirInfo := newEmbedDir(info)
			dirs[join(path, info.Name())] = dirInfo
			dirs[path].Entries = append(dirs[path].Entries, dirInfo.FileInfo)
			return
		}
		if file, err := newEmbedFile(path, info); err != nil {
			// TODO: aggregate the error, errors group?
			log.Warn(err)
			lastErr = err
		} else {
			files[path] = append(files[path], file)
		}
	})
	log.Infof("total dirs (including root) %d", len(dirs))
	log.Infof("dirs %d", len(files))
	updateDirectoryInfo(dirs, files)
	//err = writeZipFiles("t.zip", root, files)
	err = renderTemplate(dirs)
	log.Warn(err)
	return lastErr
}

func readIgnoreFile(root string) (*fsutil.Ignores, error) {
	var err error
	ignores := fsutil.NewIgnores(nil, nil)
	ignoreFile := join(root, ignoreFileName)
	if fsutil.FileExists(ignoreFile) {
		log.Debugf("found ignore file %s", ignoreFile)
		if ignores, err = ReadIgnoreFile(ignoreFile); err != nil {
			return ignores, err
		}
		// set common prefix so ignore path would work
		ignores.SetPathPrefix(root)
		log.Debugf("ignore patterns %v", ignores.Patterns())
	}
	return ignores, nil
}

func newEmbedDir(info os.FileInfo) *EmbedDir {
	return &EmbedDir{
		FileInfo: FileInfo{
			FileName:    info.Name(),
			FileSize:    info.Size(),
			FileMode:    info.Mode(),
			FileModTime: info.ModTime(),
			FileIsDir:   info.IsDir(),
		},
	}
}

func newEmbedFile(path string, info os.FileInfo) (*EmbedFile, error) {
	if b, err := ioutil.ReadFile(join(path, info.Name())); err != nil {
		return nil, errors.Wrap(err, "can't read file from disk")
	} else {
		return &EmbedFile{
			FileInfo: FileInfo{
				FileName:    info.Name(),
				FileSize:    info.Size(),
				FileMode:    info.Mode(),
				FileModTime: info.ModTime(),
				FileIsDir:   info.IsDir(),
			},
			Data: b,
		}, nil
	}
}

func updateDirectoryInfo(dirs map[string]*EmbedDir, flatFiles map[string][]*EmbedFile) {
	//folders := make(map[string][]FileInfo, len(flatFiles) + 1)
	for path, files := range flatFiles {
		log.Infof("path %s files %d", path, len(files))
		for _, f := range files {
			log.Infof("add %s to path %s", f.Name(), path)
			dirs[path].Entries = append(dirs[path].Entries, f.FileInfo)
		}
	}
}

func writeZipFiles(dst string, root string, flatFiles map[string][]*EmbedFile) error {
	var lastErr error
	buf := &bytes.Buffer{}
	w := zip.NewWriter(buf)
	for path, files := range flatFiles {
		for _, f := range files {
			log.Infof("write file %s FileSize %d", f.FileName, len(f.Data))
			lastErr = writeZipFile(w, root, path, f)
		}
	}

	if lastErr != nil {
		return lastErr
	}
	if err := w.Close(); err != nil {
		return errors.Wrap(err, "can't close zip writer")
	}
	if err := ioutil.WriteFile(dst, buf.Bytes(), 0666); err != nil {
		return errors.Wrap(err, "can't write zip file")
	}
	return nil
}

func writeZipFile(w *zip.Writer, root string, path string, file *EmbedFile) error {
	header, err := zip.FileInfoHeader(&file.FileInfo)
	if err != nil {
		return errors.Wrap(err, "can't create file header")
	}
	header.Method = zip.Deflate
	header.Name = strings.TrimLeft(join(path, file.FileInfo.Name()), root)
	f, err := w.CreateHeader(header)
	if err != nil {
		return errors.Wrap(err, "can't add file to zip")
	}
	if _, err := f.Write(file.Data); err != nil {
		return errors.Wrap(err, "can't write zip file content")
	}
	return nil
}

// TODO: trim root
func renderTemplate(dirs map[string]*EmbedDir) error {
	t, err := template.New("noodleembed").Parse(embedTemplate)
	if err != nil {
		return errors.Wrap(err, "can't parse embed template")
	}
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, map[string]interface{}{
		"pkg": "main", // TODO: allow config package FileName
		"dir": dirs,
	}); err != nil {
		return errors.Wrap(err, "can't execute template")
	}
	log.Info(buf.String())
	if b, err := format.Source(buf.Bytes()); err != nil {
		return errors.Wrap(err, "can't format go code")
	} else {
		ioutil.WriteFile("_examples/embed/t.go", b, 0666)
	}
	return nil
}

func (i *FileInfo) Name() string {
	return i.FileName
}

func (i *FileInfo) Size() int64 {
	return i.FileSize
}

func (i *FileInfo) Mode() os.FileMode {
	return i.FileMode
}

func (i *FileInfo) ModTime() time.Time {
	return i.FileModTime
}

func (i *FileInfo) IsDir() bool {
	return i.FileIsDir
}

func (i *FileInfo) Sys() interface{} {
	return nil
}