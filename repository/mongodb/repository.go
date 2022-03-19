package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/dealancer/validate.v2"

	"github.com/bhanupbalusu/gotestgrpcms/domain/application/logic"
	"github.com/bhanupbalusu/gotestgrpcms/domain/interface/model"
	r "github.com/bhanupbalusu/gotestgrpcms/domain/interface/repo"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoRepository(mongoURL string, mongoDB string, mongoTimeout int) (r.ProductRepoInterface, error) {
	repo := &mongoRepository{
		database: mongoDB,
		timeout:  time.Duration(mongoTimeout) * time.Second,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, err
	}
	repo.client = client
	return repo, nil
}

func (m *mongoRepository) GetProducts() (*[]model.ProductModel, error) {
	var resultList []model.ProductModel

	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	collection := m.client.Database(m.database).Collection("product_quantity_schedular")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments || err == mongo.ErrNilValue {
			return nil, errors.Wrap(logic.ErrRedirectNotFound, "domain.repository.mongo.repository.GetProducts")
		}
		return nil, errors.Wrap(err, "domain.repository.mongo.repository.GetProducts")
	}

	if err = cursor.All(ctx, &resultList); err != nil {
		errors.Wrap(err, "domain.repository.mongo.repository.GetProducts.cursor.All")
		log.Fatal(err)
	}
	fmt.Println(resultList)

	return &resultList, nil
}

func (m *mongoRepository) GetProductByID(id string) (*model.ProductModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	result := &model.ProductModel{}
	collection := m.client.Database(m.database).Collection("product_quantity_schedular")

	filter := bson.M{"product_id": id}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(logic.ErrRedirectNotFound, "domain.repository.mongo.repository.GetProductByID")
		}
		return nil, errors.Wrap(err, "domain.repository.mongo.repository.GetProductByID")
	}
	return result, nil

}

func (m *mongoRepository) CreateProduct(pm *model.ProductModel) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	collection := m.client.Database(m.database).Collection("product_quantity_schedular")

	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"pre_order_request_id": pm.PreOrderRequestId,
			"customer_id":          pm.CustomerId,
			"product_details": bson.M{
				"product_name": pm.ProductDetails.ProductName,
				"description":  pm.ProductDetails.Description,
				"ImageUrl":     pm.ProductDetails.ImageUrl,
			},
			"quantity_details": bson.M{
				"bulk_quantity": bson.M{
					"volume": pm.QuantityDetails.BulkQuantity.Volume,
					"units":  pm.QuantityDetails.BulkQuantity.Units,
				},
				"price": bson.M{
					"amount":   pm.QuantityDetails.Price.Amount,
					"currency": pm.QuantityDetails.Price.Currency,
					"per_unit": pm.QuantityDetails.Price.PerUnit,
					"units":    pm.QuantityDetails.Price.Units,
				},
			},
			"schedular": bson.M{
				"start_date": pm.Schedular.StartDate,
				"end_date":   pm.Schedular.EndDate,
			},
			"created_at": pm.CreatedAt,
			"updated_at": pm.UpdatedAt,
		},
	)
	if err != nil {
		return "", errors.Wrap(err, "domain.repository.mongo.repository.CreateProduct")
	}
	return pm.ProductID.Hex(), nil
}

func (m *mongoRepository) UpdateProduct(pm *model.ProductModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	collection := m.client.Database(m.database).Collection("product_quantity_schedular")

	if err := validate.Validate(pm); err != nil {
		return errors.Wrap(logic.ErrRedirectInvalid, "domain.repository.mongo.repository.UpdateProduct")
	}

	filter := bson.M{"_id": pm.ProductID}

	update := bson.M{
		"$set": bson.M{
			"pre_order_request_id": pm.PreOrderRequestId,
			"customer_id":          pm.CustomerId,
			"product_details": bson.M{
				"product_name": pm.ProductDetails.ProductName,
				"description":  pm.ProductDetails.Description,
				"ImageUrl":     pm.ProductDetails.ImageUrl,
			},
			"quantity_details": bson.M{
				"bulk_quantity": bson.M{
					"volume": pm.QuantityDetails.BulkQuantity.Volume,
					"units":  pm.QuantityDetails.BulkQuantity.Units,
				},
				"price": bson.M{
					"amount":   pm.QuantityDetails.Price.Amount,
					"currency": pm.QuantityDetails.Price.Currency,
					"per_unit": pm.QuantityDetails.Price.PerUnit,
					"units":    pm.QuantityDetails.Price.Units,
				},
			},
			"schedular": bson.M{
				"start_date": pm.Schedular.StartDate,
				"end_date":   pm.Schedular.EndDate,
			},
			"created_at": pm.CreatedAt,
			"updated_at": pm.UpdatedAt,
		},
	}

	_, err := collection.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		if err == mongo.ErrNoDocuments || err == mongo.ErrNilValue {
			return errors.Wrap(logic.ErrRedirectNotFound, "domain.repository.mongo.repository.UpdateProduct")
		}
		return errors.Wrap(err, "domain.repository.mongo.repository.UpdateProduct")
	}

	return nil
}

func (m *mongoRepository) DeleteProduct(pm *model.ProductModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	collection := m.client.Database(m.database).Collection("product_quantity_schedular")

	if err := validate.Validate(pm); err != nil {
		return errors.Wrap(logic.ErrRedirectInvalid, "domain.repository.mongo.repository.DeleteProduct")
	}

	_, err := collection.DeleteOne(ctx, bson.M{"_id": pm.ProductID})
	if err != nil {
		if err == mongo.ErrNoDocuments || err == mongo.ErrNilValue {
			return errors.Wrap(logic.ErrRedirectNotFound, "domain.repository.mongo.repository.DeleteProduct")
		}
		return errors.Wrap(err, "domain.repository.mongo.repository.DeleteProduct")
	}

	return nil
}
