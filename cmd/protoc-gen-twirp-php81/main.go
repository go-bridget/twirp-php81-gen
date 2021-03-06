package main

import (
	"flag"
	"path"

	"github.com/apex/log"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/go-bridget/twirp-php81-gen/internal/php81"
)

var _ = spew.Dump

func init() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	var (
		flags flag.FlagSet

		namespace = flags.String("namespace", "Twirp", "PHP namespace for generated messages")
		folder    = flags.String("folder", "", "Output folder for generated .php files")
		prefix    = flags.String("prefix", "/twirp", "Prefix for RPC apis (default: /twirp)")
		router    = flags.Bool("router", false, "Output generated gRPC transcoding router")
	)
	opts := protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(func(gen *protogen.Plugin) error {
		options := &php81.Options{
			Namespace: *namespace,
			Folder:    *folder,
			Router:    *router,
			Prefix:    *prefix,
		}

		for _, f := range gen.Files {
			in := f.Desc.Path()
			log.Debugf("generating: %q", in)

			if !f.Generate {
				log.Debugf("skip generating: %q", in)
				continue
			}

			files, err := php81.Load(in, options)
			if err != nil {
				return err
			}

			for _, file := range files {
				g := gen.NewGeneratedFile(path.Join(*folder, file.Filename()), f.GoImportPath)
				if _, err := g.Write(file.Bytes()); err != nil {
					return err
				}
			}

		}
		return nil
	})
}
