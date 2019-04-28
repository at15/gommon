package main

import (
	"os"
	"time"

	dlog "github.com/dyweb/gommon/log"
	"github.com/dyweb/gommon/log/handlers/cli"
)

var logReg = dlog.NewRegistry()
var log = logReg.Logger()

func main() {
	dlog.SetHandler(cli.New(os.Stderr, true))

	if len(os.Args) > 1 {
		if os.Args[1] == "nocolor" || os.Args[1] == "no" {
			dlog.SetHandler(cli.NewNoColor(os.Stderr))
		}
	}

	log.Info("hi")
	log.Infof("open file %s", "foo.yml")
	log.InfoF("open",
		dlog.Str("file", "foo.yml"),
		dlog.Int("mode", 0666),
	)
	log.Warn("I am yellow")
	func() {
		defer func() {
			if r := recover(); r != nil {
				log.Info("recovered", r)
			}
		}()
		log.Panic("I just want to panic")
	}()
	dlog.SetLevel(dlog.DebugLevel)
	log.Debug("I will sleep for a while")
	time.Sleep(500 * time.Millisecond)
	log.Fatal("I am red")
}
