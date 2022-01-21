package php81

import (
	"github.com/emicklei/proto"
	"github.com/apex/log"
)

type generator struct {
	options *Options

	hasRPC bool

	files []*File
}

func NewGenerator(options *Options) *generator {
	return &generator{
		options: options,
		files: []*File{}
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
	log.Infof("importing %s", i.Filename)

	definition, err := loadProto(i.Filename)
	if err != nil {
		log.Infof("Can't load %s, err=%s, ignoring (want to make PR?)", i.Filename, err)
		return
	}

	handlers := []proto.Handler{
		proto.WithImport(g.Import),
		proto.WithMessage(g.Message),
	}
	proto.Walk(definition, handlers...)
}

func (g *generator) RPC(rpc *proto.RPC) {
	parent, ok := rpc.Parent.(*proto.Service)
	if !ok {
		panic("parent is not proto.service")
	}
	g.hasRPC = true
}

func (g *generator) Message(msg *proto.Message) {
	file := NewFile(msg.Name + ".php", g.options.namespace)
	g.files = append(g.files, file)
}
