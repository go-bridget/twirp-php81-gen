package model

import (
	"bytes"
	"path"
	"strings"
)

type Router struct {
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

func NewRouter(name, namespace string, routes ...*Route) *Router {
	return &Router{
		name:      name,
		namespace: namespace,
		routes:    routes,
	}
}

func (f *Router) AddRoute(n *Route) {
	f.routes = append(f.routes, n)
}

func (f *Router) Name() string {
	return f.name
}

func (f *Router) Bytes() []byte {
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
	className := strings.TrimSuffix(path.Base(f.name), ".php")

	handlerClassName := strings.ReplaceAll(className, "Router", "Handler")

	f.print("class " + className)
	f.print("{")
	f.print("\tpublic function Mount(\\Slim\\App $app)")
	f.print("\t{")
	for _, v := range f.routes {
		handlerCall := "\\" + f.namespace + "\\" + handlerClassName + ":" + v.Name
		f.print("\t\t$app->" + strings.ToLower(v.Method) + "('" + v.URL + "', '" + handlerCall + "')->setName('" + v.Name + "');")
	}
	f.print("\t}")
	f.print("}")

	return f.contents.Bytes()
}

func (f *Router) use(name string) {
	f.uses = append(f.uses, name)
}

func (f *Router) print(lines ...string) {
	if len(lines) == 0 {
		f.contents.WriteString("\n")
		return
	}
	for _, line := range lines {
		f.contents.WriteString(line + "\n")
	}
}
