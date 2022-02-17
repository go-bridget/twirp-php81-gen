package php81

import (
	"fmt"

	"github.com/emicklei/proto"
)

// Load takes a .proto file, and generates multiple output File{} objects
func Load(filename string, options *Options) ([]FileGenerator, error) {
	definition, err := loadProto(filename)
	if err != nil {
		return nil, fmt.Errorf("got error loading file %s: %w", filename, err)
	}

	gen := NewGenerator(options)

	// main file for all the relevant info
	proto.Walk(definition, gen.Handlers()...)

	if gen.hasRPC {
		return gen.Files(), nil
	}
	return nil, nil
}
