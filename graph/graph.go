package graph

import (
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
)

type Graph interface {
	Client() *dgo.Dgraph
	SetSchema(schema string) error
	SetJSON(data []byte) (*api.Response, error)
	SetInterface(data interface{}) (*api.Response, error)
	DeleteJSON(data []byte) (*api.Response, error)
	DeleteInterface(data interface{}) (*api.Response, error)
	Query(query string) (*api.Response, error)
	QueryWithVars(query string, vars map[string]string) (*api.Response, error)
	NodeValueExists(node, value string) bool
}
