package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aferen/firecracker-api/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type VMRepositoryI interface {
	Create(ctx context.Context, req model.VM) error
}

type VMRepository struct {
	mongoDB *mongo.Database
}

func NewVMepository(mongo *mongo.Database) VMRepositoryI {
	return &VMRepository{
		mongoDB: mongo,
	}
}

func (vm VMRepository) Create(ctx context.Context, req model.VM) (err error) {

	dataReq := bson.M{
		"ip":      req.Ip,
		"vmid":    req.VmID,
		"machine": req.Machine,
	}

	query, err := vm.mongoDB.Collection("vms").InsertOne(ctx, dataReq)
	if err != nil {
		log.Println("error")
	}

	if oid, ok := query.InsertedID.(primitive.ObjectID); ok {
		productID := oid.Hex()
		vmID := bson.M{"_id": oid}
		log.Printf("inserted %s %s", productID, vmID)
	} else {
		err = errors.New(fmt.Sprint("can't get inserted ID ", err))
		log.Println("error")
	}

	return nil
}
