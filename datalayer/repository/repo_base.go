package repository

import (
	"context"
	"iter"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newBaseRepository[T any](db *mongo.Database, name string, opts ...*options.CollectionOptions) *baseRepository[T] {
	coll := db.Collection(name, opts...)
	return &baseRepository[T]{coll: coll}
}

type baseRepository[T any] struct {
	coll *mongo.Collection
}

func (br *baseRepository[T]) CreateIndex(ctx context.Context) error {
	return nil
}

func (br *baseRepository[T]) InsertMany(ctx context.Context, docs []*T, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	dats := make([]any, len(docs))
	for i, doc := range docs {
		dats[i] = doc
	}

	return br.coll.InsertMany(ctx, dats, opts...)
}

func (br *baseRepository[T]) Indexes() mongo.IndexView {
	return br.coll.Indexes()
}

func (br *baseRepository[T]) FindOne(ctx context.Context, filter any, opts ...*options.FindOneOptions) (*T, error) {
	return br.decodeSingleResult(br.coll.FindOne(ctx, filter, opts...))
}

func (br *baseRepository[T]) All(ctx context.Context, filter any, limit int64, opts ...*options.FindOptions) iter.Seq2[[]*T, error] {
	return func(yield func([]*T, error) bool) {
		var skip int64
		opt := options.Find().SetLimit(limit)
		for {
			opt.SetSkip(skip)
			optis := append(opts, opt)
			cur, err := br.coll.Find(ctx, filter, optis...)
			if err != nil {
				yield(nil, err)
				break
			}
			ts, err := br.decodeCursor(ctx, cur)
			if err != nil {
				yield(nil, err)
				break
			}
			num := len(ts)
			if num == 0 {
				break
			}
			skip += int64(num)
			if !yield(ts, nil) {
				break
			}
		}
	}
}

func (*baseRepository[T]) decodeSingleResult(ret *mongo.SingleResult) (*T, error) {
	t := new(T)
	if err := ret.Decode(t); err != nil {
		return nil, err
	}

	return t, nil
}

func (*baseRepository[T]) decodeCursor(ctx context.Context, cur *mongo.Cursor) ([]*T, error) {
	//goland:noinspection GoUnhandledErrorResult
	defer cur.Close(context.Background())

	ts := make([]*T, 0, 10)
	if err := cur.All(ctx, &ts); err != nil {
		return nil, err
	}

	return ts, nil
}
