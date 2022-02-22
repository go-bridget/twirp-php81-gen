package model

import (
	"github.com/emicklei/proto"
)

type Route struct {
	RPC    *proto.RPC
	Method string
	URL    string
}

func NewRoute(rpc *proto.RPC, method, url string) *Route {
	return &Route{
		RPC:    rpc,
		Method: method,
		URL:    url,
	}
}
