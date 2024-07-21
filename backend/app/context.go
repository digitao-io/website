package app

import (
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/mongo"
)

type Context struct {
	Configuration *Configuration
	Database      *mongo.Database
	Objstorage    *minio.Client
}
