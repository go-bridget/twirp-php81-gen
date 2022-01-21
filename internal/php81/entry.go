package php81

import (
	"github.com/emicklei/proto"
)

// Load takes a .proto file, and generates multiple output File{} objects
func Load(filename string, options *Options) ([]File, error) {
	definition, err := loadProto(filename)
	if err != nil {
		return nil, err
	}

	gen := newGenerator(options)

	// main file for all the relevant info
	proto.Walk(definition, handlers()...)

	if gen.hasRPC {
		return gen.files, nil
	}
	return nil, nil
}
