/**
*@program: mongo-info
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-25 15:55
 */
package examples

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDB struct {
	mgo *mongo.Client
	db *mongo.Database
}
type MongoCollection struct {
	collection *mongo.Collection
	mgo *mongo.Client
	db *mongo.Database
}

func New(uri string,db string) (*MongoDB,error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil,err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil,err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil,err
	}

	return &MongoDB{
		mgo:client,
		db:client.Database(db),
	},nil
}

func (m *MongoDB) Collection(collection string,opts ...*options.CollectionOptions) *MongoCollection {
	return &MongoCollection{
		collection:m.db.Collection(collection,opts...),
		db:m.db,
		mgo:m.mgo,
	}
}

func (c *MongoCollection) Insert(data interface{}) (*mongo.InsertOneResult, error) {
	return c.collection.InsertOne(context.TODO(),data)
}

func (c *MongoCollection) Select(filter interface{}) (*mongo.SingleResult,error)  {
	si := c.collection.FindOne(context.TODO(),filter)
	return si,si.Err()
}
