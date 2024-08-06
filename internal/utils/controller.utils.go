package utils

import (
	"errors"
	"strings"

	"github.com/bibipe2024/goautdoor/internal/api/spec"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func PrepareIDResponse(result *mongo.InsertOneResult) (*spec.DefaulCreatedResponse, error) {
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("could not get id")
	}
	hexString := oid.Hex()
	res := spec.DefaulCreatedResponse{
		Id: &hexString,
	}

	return &res, nil
}

func splitBytesToStrings(byteArray []byte) []string {
	str := string(byteArray)
	return strings.Split(str, ",")
}
