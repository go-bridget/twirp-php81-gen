package model

import (
	"bytes"
	"fmt"
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
		uses: []string{
			"Psr\\Http\\Message\\ResponseInterface as Response",
			"Psr\\Http\\Message\\ServerRequestInterface as Request",
			"Slim\\Routing\\RouteCollectorProxy",
		},
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
	f.print("\tpublic function Mount(\\Slim\\App $app, string $serviceClass)")
	f.print("\t{")

	urls := make([]string, len(f.routes))
	for k, v := range f.routes {
		urls[k] = v.URL
	}

	var prefix string
	if len(urls) > 0 {
		prefix = urls[0]
		for _, url := range urls {
			for prefix != "" && !strings.HasPrefix(url, prefix) {
				prefix = prefix[0 : len(prefix)-1]
			}
		}
	}
	prefix = strings.TrimSuffix(prefix, "/")

	if prefix != "" {
		f.print("\t\t$app->group(\"" + prefix + "\", function (RouteCollectorProxy $group) use ($serviceClass)")
		f.print("\t\t{")
		for _, v := range f.routes {
			var (
				methods = "[\"" + strings.ToUpper(v.Method) + "\"]"
				url     = "\"" + strings.TrimPrefix(v.URL, prefix) + "\""
				name    = "\"" + v.Name + "\""
			)
			f.print(fmt.Sprintf("\t\t\t$group->map(%s, %s, function (Request $request, Response $response, array $args) use ($serviceClass) {", methods, url))
			f.print("\t\t\t\t$service = new $serviceClass;")
			f.print(fmt.Sprintf("\t\t\t\t$handler = new %s($service);", handlerClassName))
			f.print(fmt.Sprintf("\t\t\t\treturn $handler->%s($request, $response, $args);", v.Name))
			f.print(fmt.Sprintf("\t\t\t})->setName(%s);", name))
		}
		f.print("\t\t});")
		f.print("\t}")
		f.print("}")
		return f.contents.Bytes()
	}

	for _, v := range f.routes {
		var (
			methods = "[\"" + strings.ToUpper(v.Method) + "\"]"
			url     = "\"" + strings.TrimPrefix(v.URL, prefix) + "\""
			name    = "\"" + v.Name + "\""
			handler = "$request" + v.Name
		)
		f.print("\t\t$request" + v.Name + " = function(Request $request, Response $response, array $args) {")
		f.print("\t\t\t$service = new $serviceClass;")
		f.print("\t\t\t$service->" + v.Name + "($request, $response, $args);")
		f.print("\t\t};")

		f.print(fmt.Sprintf("\t\t$app->map(%s, %s, %s)->setName(%s);", methods, url, handler, name))
	}

	f.print("\t}")
	f.print("}")

	return f.contents.Bytes()
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
