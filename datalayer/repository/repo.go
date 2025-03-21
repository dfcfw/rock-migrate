package repository

import (
	"context"
	"iter"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository[T any] interface {
	// InsertMany executes an insert command to insert multiple documents into the collection. If write errors occur
	// during the operation (e.g. duplicate key error), this method returns a BulkWriteException error.
	//
	// The documents parameter must be a slice of documents to insert. The slice cannot be nil or empty. The elements must
	// all be non-nil. For any document that does not have an _id field when transformed into BSON, one will be added
	// automatically to the marshalled document. The original document will not be modified. The _id values for the inserted
	// documents can be retrieved from the InsertedIDs field of the returned InsertManyResult.
	//
	// The opts parameter can be used to specify options for the operation (see the options.InsertManyOptions documentation.)
	//
	// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/insert/.
	InsertMany(ctx context.Context, docs []*T, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)

	All(ctx context.Context, filter any, limit int64, opts ...*options.FindOptions) iter.Seq2[[]*T, error]

	Indexes() mongo.IndexView

	FindOne(ctx context.Context, filter any, opts ...*options.FindOneOptions) (*T, error)

	IndexCreator
}

type IndexCreator interface {
	CreateIndex(ctx context.Context) error
}

func CreateIndex(ctx context.Context, indexes []IndexCreator) error {
	for _, idx := range indexes {
		if err := idx.CreateIndex(ctx); err != nil {
			return err
		}
	}

	return nil
}
