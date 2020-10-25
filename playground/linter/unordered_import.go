package linter

import (
	"github.com/dyweb/gommon/errors"
	"github.com/dyweb/gommon/util/fsutil"
	"golang.org/x/tools/imports"
	"strings"
	"unicode"

	"bytes"
)

func foo() error {
	var _ = unicode.ToUpper('a')
	var msg = strings.Join([]string{"foo", "error"}, " ")
	_, err := imports.Process("", nil, nil)
	_, err = fsutil.CreateFileAndPath("foo", "boar.txt")
	var buf = bytes.Buffer{}
	buf.Write([]byte(msg))
	return errors.Wrap(err, msg)
}
