package model

import (
	"bytes"
	"fmt"
)

type Handler struct {
	// Name (full path) for the generated file
	name string

	// Namespace to use in file
	namespace string

	// Used objects from other namespaces
	uses []string

	// Routes
	routes []*Route

	// Contents for final generated file
	contents bytes.Buffer
}

func NewHandler(name, namespace string, routes ...*Route) *Handler {
	return &Handler{
		name:      name,
		namespace: namespace,
		routes:    routes,
		uses: []string{
			"Psr\\Http\\Message\\ResponseInterface as Response",
			"Psr\\Http\\Message\\ServerRequestInterface as Request",
		},
	}
}

func (f *Handler) Filename() string {
	return f.name + "Handler.php"
}

func (f *Handler) Bytes() []byte {
	f.print("<?php")
	f.print()
	f.print("namespace " + f.namespace + ";")
	f.print()
	for _, use := range f.uses {
		f.print("use " + use + ";")
	}
	if len(f.uses) > 0 {
		f.print()
	}
	className := f.name + "Handler"
	serviceClassName := f.name

	f.print("class " + className)
	f.print("{")
	f.print(fmt.Sprintf("\tprivate %s $service;", serviceClassName))
	f.print()

	f.print(fmt.Sprintf("\tpublic function __construct(%s $service)", serviceClassName))
	f.print("\t{")
	f.print("\t\t$this->service = $service;")
	f.print("\t}")

	for _, v := range f.routes {
		f.print()
		f.print(fmt.Sprintf("\tpublic function %s(Request $request, Response $response, array $args): Response", v.Name))
		f.print("\t{")
		f.print(fmt.Sprintf("\t\t$serviceRequest = new %sRequest($request);", v.Name))
		f.print(fmt.Sprintf("\t\t$response->writeJSON($this->service->%s($serviceRequest));", v.Name))
		f.print("\t}")
	}
	f.print("}")

	return f.contents.Bytes()
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
