package nosql

import (
	"context"

	"cloud.google.com/go/firestore"
)

type FireStore struct {
	FireStore *firestore.Client
}

func OpenFireStore(ctx context.Context, projectID string) (NoSQL, error) {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &FireStore{
		FireStore: client,
	}, nil
}

func (db *FireStore) Client() *firestore.Client {
	return db.FireStore
}

func (db *FireStore) Add(ctx context.Context, collection string, data map[string]interface{}) error {
	if _, _, err := db.FireStore.Collection(collection).Add(ctx, data); err != nil {
		return err
	}
	return nil
}

func (db *FireStore) Set(ctx context.Context, collection, doc string, data map[string]interface{}) error {
	if _, err := db.FireStore.Collection(collection).Doc(doc).Set(ctx, data); err != nil {
		return err
	}
	return nil
}

func (db *FireStore) Delete(ctx context.Context, collection, doc string) error {
	if _, err := db.FireStore.Collection(collection).Doc(doc).Delete(ctx); err != nil {
		return err
	}
	return nil
}

func (db *FireStore) DeleteField(ctx context.Context, collection, doc, path string) error {
	if _, err := db.FireStore.Collection(collection).Doc(doc).Update(ctx, []firestore.Update{
		{
			Path:  path,
			Value: firestore.Delete,
		},
	}); err != nil {
		return err
	}
	return nil
}
