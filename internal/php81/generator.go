package php81

import (
	"path"

	"github.com/apex/log"
	"github.com/emicklei/proto"
)

type FileGenerator interface {
	Name() string
	Bytes() []byte
}

type generator struct {
	options *Options

	hasRPC bool

	files []FileGenerator

	service *Service
}

func NewGenerator(options *Options) *generator {
	return &generator{
		options: options,
		files:   make([]FileGenerator, 0),
	}
}

func (g *generator) Handlers() []proto.Handler {
	return []proto.Handler{
		proto.WithRPC(g.RPC),
		proto.WithMessage(g.Message),
		proto.WithImport(g.Import),
	}
}

func (g *generator) Import(i *proto.Import) {
	log.Debugf("importing %s", i.Filename)

	definition, err := loadProto(i.Filename)
	if err != nil {
		log.Debugf("Can't load %s, err=%s, ignoring", i.Filename, err)
		return
	}

	handlers := []proto.Handler{
		proto.WithImport(g.Import),
		proto.WithMessage(g.Message),
	}
	proto.Walk(definition, handlers...)
}

// RPC marks if a service is defined in the .proto file.
// The service part is mandatory, in order to generate
// relevant rpc request and response structures.
func (g *generator) RPC(rpc *proto.RPC) {
	parent, ok := rpc.Parent.(*proto.Service)
	if !ok {
		panic("parent is not proto.service")
	}
	g.hasRPC = true

	if g.service == nil {
		filename := path.Join(g.options.Folder, parent.Name+".php")
		g.service = NewService(filename, g.options.Namespace)
		g.files = append(g.files, g.service)
	}

	svc := g.service
	svc.comment = comment(parent.Comment)
	svc.addFunction([]string{
		"/** " + comment(rpc.Comment) + " */",
		"public function " + rpc.Name + "(" + rpc.RequestType + " $req): " + rpc.ReturnsType + ";",
	});
}

func (g *generator) Message(msg *proto.Message) {
	filename := path.Join(g.options.Folder, msg.Name+".php")
	file := NewMessage(filename, g.options.Namespace)

	allFields := msg.Elements

	for _, element := range msg.Elements {
		switch val := element.(type) {
		case *proto.Oneof:
			// We're unpacking val.Elements into the field list,
			// which may or may not be correct. The oneof semantics
			// likely bring in edge-cases.
			allFields = append(allFields, val.Elements...)
		default:
			// No need to unpack for *proto.NormalField,...
			log.Debugf("prepare: uknown field type: %T", element)
		}
	}

	addField := func(field *proto.Field, repeated bool) {
		file.addField(Field{
			Name:     field.Name,
			Type:     field.Type,
			Repeated: repeated,
		})
	}

	for _, element := range allFields {
		switch val := element.(type) {
		case *proto.Comment:
		case *proto.Oneof:
			// Nothing.
		case *proto.OneOfField:
			addField(val.Field, false)
		case *proto.MapField:
			addField(val.Field, false)
		case *proto.NormalField:
			addField(val.Field, val.Repeated)
		default:
			log.Infof("Unknown field type: %T", element)
		}
	}

	g.files = append(g.files, file)
}
