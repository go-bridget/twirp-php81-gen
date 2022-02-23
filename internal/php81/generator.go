package php81

import (
	"path"
	"strings"

	"github.com/apex/log"
	"github.com/emicklei/proto"

	"github.com/go-bridget/twirp-php81-gen/internal/php81/model"
)

type FileGenerator interface {
	Filename() string
	Bytes() []byte
}

type generator struct {
	options *Options

	hasRPC bool

	files []FileGenerator

	routes []*model.Route

	service     *model.Service
	serviceName string

	packageName string
}

func (g *generator) Files() []FileGenerator {
	router := model.NewRouter(g.serviceName, g.options.Namespace)
	if router.Prefix = strings.TrimRight(g.options.Prefix, "/"); router.Prefix == "" {
		router.Prefix = "/twirp"
	}
	router.ServiceName = g.serviceName
	router.PackageName = g.packageName
	router.RPCs = g.service.RPCs
	router.Routes = g.routes

	handler := model.NewHandler(g.serviceName, g.options.Namespace, g.service.RPCs)

	return append([]FileGenerator{g.service, router, handler}, g.files...)
}

func NewGenerator(options *Options) *generator {
	return &generator{
		options: options,
		files:   make([]FileGenerator, 0),
	}
}

func (g *generator) Handlers() []proto.Handler {
	return []proto.Handler{
		proto.WithPackage(g.Package),
		proto.WithService(g.Service),
		proto.WithOption(g.Option),
		proto.WithRPC(g.RPC),
		proto.WithMessage(g.Message),
		proto.WithImport(g.Import),
	}
}

func (g *generator) Package(pkg *proto.Package) {
	g.packageName = pkg.Name
}

func (g *generator) Service(service *proto.Service) {
	g.serviceName = service.Name

	svc := model.NewService(service.Name, g.options.Namespace)
	svc.Namespace = g.options.Namespace
	svc.Name = path.Join(g.options.Folder, service.Name)
	g.service = svc
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
	service, ok := rpc.Parent.(*proto.Service)
	if !ok {
		panic("parent is not proto.service")
	}
	g.hasRPC = true

	svc := g.service
	svc.Comment = comment(service.Comment)
	svc.AddRPC(rpc)
	svc.AddFunction([]string{
		"/** " + comment(rpc.Comment) + " */",
		"public function " + rpc.Name + "(" + rpc.RequestType + " $req): " + rpc.ReturnsType + ";",
	})

}

// Option collect grpc-proxy transcoding routes from .proto
func (g *generator) Option(option *proto.Option) {
	rpc, ok := option.Parent.(*proto.RPC)
	if !ok {
		return
	}
	// sniffing only for rpc options
	if option.Name != "(google.api.http)" {
		return
	}
	for _, v := range option.AggregatedConstants {
		method := strings.ToLower(v.Name)
		url := v.Literal.Source
		g.routes = append(g.routes, model.NewRoute(rpc, method, url))
	}
}

func (g *generator) Message(msg *proto.Message) {
	file := model.NewMessage(msg.Name, g.options.Namespace)

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
		file.AddField(model.NewField(field.Type, field.Name, repeated))
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
