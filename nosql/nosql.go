package nosql

import (
	"context"

	"cloud.google.com/go/firestore"
)

type NoSQL interface {
	Client() *firestore.Client
	Add(ctx context.Context, collection string, data map[string]interface{}) error
	Set(ctx context.Context, collection, doc string, data map[string]interface{}) error
	Delete(ctx context.Context, collection, doc string) error
	DeleteField(ctx context.Context, collection, doc, path string) error
}
