package graph

import (
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
)

type Graph interface {
	Client() *dgo.Dgraph
	SetSchema(schema string) error
	SetJSON(data []byte) (*api.Assigned, error)
	SetInterface(data interface{}) (*api.Assigned, error)
	DeleteJSON(data []byte) (*api.Assigned, error)
	DeleteInterface(data interface{}) (*api.Assigned, error)
	Query(query string) (*api.Response, error)
	QueryWithVars(query string, vars map[string]string) (*api.Response, error)
	NodeValueExists(node, value string) bool
}
