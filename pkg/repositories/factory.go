package repositories

import (
	"context"
	"warehouse-management/pkg/database"
	"warehouse-management/pkg/repositories/product"
)

func NewFactory(mainDataSource database.Instance, dbName string) Factory {
	return Factory{
		mainDataSource: mainDataSource,
		dbName:         dbName,
	}
}

type Factory struct {
	mainDataSource database.Instance
	dbName         string
}

func (f Factory) MainDB() database.Instance {
	return f.mainDataSource
}

func (f Factory) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	return f.MainDB().WithTransaction(ctx, fn)
}

func (f Factory) WithSession(ctx context.Context, fn func(context.Context) error) error {
	return f.MainDB().WithSession(ctx, fn)
}

func (f Factory) ProductRepo() repositories.ProductRepo {
	return repositories.NewProductRepo(f.MainDB().Database(f.dbName))
}
