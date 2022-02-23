package model

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/apex/log"
	"github.com/emicklei/proto"
)

type Handler struct {
	// Name (full path) for the generated file
	Name string

	// Namespace to use in file
	Namespace string

	// Used objects from other namespaces
	uses []string

	// Routes
	RPCs []*proto.RPC

	// Contents for final generated file
	contents bytes.Buffer
}

func NewHandler(name, namespace string, rpcs []*proto.RPC) *Handler {
	return &Handler{
		Name:      name,
		Namespace: namespace,
		RPCs:      rpcs,
	}
}

func (f *Handler) Filename() string {
	return f.Name + "Handlers.php"
}

func (f *Handler) Bytes() []byte {
	funcMap := template.FuncMap{
		"comment": comment,
	}
	tpl, err := template.New("render-php81-handler").Funcs(funcMap).Parse(handlerTemplate)
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

func (f *Handler) print(lines ...string) {
	if len(lines) == 0 {
		f.contents.WriteString("\n")
		return
	}
	for _, line := range lines {
		f.contents.WriteString(line + "\n")
	}
}

const handlerTemplate = `
<?php

namespace {{ .Namespace }};

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

abstract class {{ .Name }}Handlers implements {{ .Name }}
{
{{- range .RPCs }}
	/** {{ .Comment | comment }} */
	public function handle{{ .Name }}(Request $request, Response $response, array $args): Response
	{
		$params = new {{ .RequestType }}($request);
		$data = $this->{{ .Name }}($params);
		$response->getBody()->write(json_encode($data));
		return $response->withHeader('Content-Type', 'application/json');
	}
-{{- end }}
}
`
