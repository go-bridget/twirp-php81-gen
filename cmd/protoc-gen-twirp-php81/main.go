package main

import (
	"errors"
	"flag"

	"github.com/apex/log"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/go-bridget/twirp-gen/internal/php81"
)

var _ = spew.Dump

func init() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	var (
		flags flag.FlagSet

		namespace   = flags.String("namespace", "Twirp")
		packageName = flags.String("package", "")
		folder      = flags.String("folder", "")
	)
	opts := protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(func(gen *protogen.Plugin) error {
		options = &php81.Options{
			Namespace: *namespace,
			Package:   *packageName,
			Folder:    *folder,
		}

		for _, f := range gen.Files {
			in := f.Desc.Path()
			log.Debugf("generating: %q", in)

			if !f.Generate {
				log.Debugf("skip generating: %q", in)
				continue
			}

			// @todo, *php81.Options
			files, err := php81.Load(in, nil)
			if err != nil {
				return err
			}

			for _, file := range files {
				// todo: output folder?
				g := gen.NewGeneratedFile(file.Name())
				if _, err := g.Write(file.Bytes()); err != nil {
					return err
				}
			}

		}
		return nil
	})
}
