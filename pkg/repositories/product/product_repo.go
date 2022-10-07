package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"warehouse-management/pkg/database"
	"warehouse-management/pkg/models/product"
	"warehouse-management/pkg/repositories/internal"
)

type ProductRepo interface {
	internal.BaseRepo
	FindById(ctx context.Context, productId string) (*product.Product, error)
}

type productStruct struct {
	internal.BaseRepoImpl
}

func NewProductRepo(db database.Database) ProductRepo {
	repo := &productStruct{}
	repo.Init(db, "product")
	return repo
}

func (rcv productStruct) FindById(ctx context.Context, productId string) (*product.Product, error) {
	filter := bson.M{"product_id": productId}
	//var product *product.Product
	product := new(product.Product)
	err := rcv.BaseFindOne(ctx, filter).Decode(&product)
	if err != nil {
		//log.Logger(ctx).Error("fail to find", zap.Error(err))
		return nil, err
	}

	return product, nil
}
