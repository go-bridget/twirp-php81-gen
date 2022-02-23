package model

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/apex/log"
	"github.com/emicklei/proto"
)

type Router struct {
	// Name (full path) for the generated file
	Name string

	// Namespace to use in file
	Namespace string

	// Routes
	Routes []*Route

	// RPCs
	RPCs []*proto.RPC

	// Twirp specific options
	Prefix, PackageName, ServiceName string
}

func NewRouter(name, namespace string) *Router {
	return &Router{
		Name:      name,
		Namespace: namespace,
		Prefix:    "/twirp",
	}
}

func (f *Router) Filename() string {
	return f.Name + "Router.php"
}

func (f *Router) Bytes() []byte {
	funcMap := template.FuncMap{
		"twirpRoute": func(rpc *proto.RPC) string {
			// /<prefix>/<package>.<service>/<call>
			return fmt.Sprintf("%s/%s.%s/%s", f.Prefix, f.PackageName, f.ServiceName, rpc.Name)
		},
	}
	tpl, err := template.New("render-php81-router").Funcs(funcMap).Parse(routerTemplate)
	if err != nil {
		log.WithError(err).Errorf("error loading template for %s", f.Filename())
		return nil
	}

	out := new(bytes.Buffer)
	err = tpl.Execute(out, f)
	if err != nil {
		log.WithError(err).Errorf("error rendering template for %s", f.Filename())
		return nil
	}

	sOut := out.String()
	sOut = strings.ReplaceAll(sOut, "\n\n\n", "\n\n")
	sOut = strings.ReplaceAll(sOut, "-\n", "")
	sOut = strings.ReplaceAll(sOut, "\n-", "")
	sOut = strings.TrimSpace(sOut) + "\n"

	return []byte(sOut)
}

const routerTemplate = `
<?php

namespace {{ .Namespace }};

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Routing\RouteCollectorProxy;

class {{ .Name }}Router
{
	public function Mount(\Slim\App $app, string $serviceClass)
	{
{{- range .RPCs }}
		$app->post("{{ . | twirpRoute }}", $serviceClass . ":handle{{ .Name }}");
-{{- end }}
{{- range .Routes }}
		$app->{{ .Method }}("{{ .URL }}", $serviceClass . ":handle{{ .RPC.Name }}");
{{- end }}
	}
}
`
