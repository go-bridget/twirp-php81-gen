package php81

import (
	"os"

	"github.com/emicklei/proto"
)

func loadProto(filename string) (*proto.Proto, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	parser := proto.NewParser(reader)
	return parser.Parse()
}
