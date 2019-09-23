package graph

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/buger/jsonparser"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

type Dgraph struct {
	Dgraph *dgo.Dgraph
	Conn   *grpc.ClientConn
	mutex  sync.Mutex
}

func OpenDgraph(addr string) (Graph, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	return &Dgraph{
		Dgraph: client,
		Conn:   conn,
	}, nil
}

func (db *Dgraph) Client() *dgo.Dgraph {
	return db.Dgraph
}

func (db *Dgraph) SetSchema(schema string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	return db.Dgraph.Alter(context.Background(), &api.Operation{
		Schema: schema,
	})
}

func (db *Dgraph) SetJSON(data []byte) (*api.Response, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	mut := &api.Mutation{
		CommitNow: true,
		SetJson:   data,
	}
	return db.Dgraph.NewTxn().Mutate(context.Background(), mut)
}

func (db *Dgraph) SetInterface(data interface{}) (*api.Response, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	plainJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mut := &api.Mutation{
		CommitNow: true,
		SetJson:   plainJSON,
	}
	return db.Dgraph.NewTxn().Mutate(context.Background(), mut)
}

func (db *Dgraph) DeleteJSON(data []byte) (*api.Response, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	mut := &api.Mutation{
		CommitNow:  true,
		DeleteJson: data,
	}
	return db.Dgraph.NewTxn().Mutate(context.Background(), mut)
}

func (db *Dgraph) DeleteInterface(data interface{}) (*api.Response, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	plainJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mut := &api.Mutation{
		CommitNow:  true,
		DeleteJson: plainJSON,
	}
	return db.Dgraph.NewTxn().Mutate(context.Background(), mut)
}

func (db *Dgraph) Query(query string) (*api.Response, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	return db.Dgraph.NewTxn().Query(context.Background(), query)
}

func (db *Dgraph) QueryWithVars(query string, vars map[string]string) (*api.Response, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	return db.Dgraph.NewTxn().QueryWithVars(context.Background(), query, vars)
}

func (db *Dgraph) NodeValueExists(node, value string) bool {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	resp, err := db.Dgraph.NewTxn().Query(context.Background(), `{
		exists(func: eq(`+node+`, "`+value+`")) {
			uid
		}
	}`)
	if err != nil {
		return false
	}
	if _, err = jsonparser.GetString(resp.GetJson(), "exists", "[0]", "uid"); err != nil {
		return false
	}
	return true
}
