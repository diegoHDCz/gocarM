package mongostore

import (
	"context"
	"errors"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/bibipe2024/goautdoor/internal/api/spec"
	"github.com/bibipe2024/goautdoor/internal/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	db *mongo.Client
}

func New(db *mongo.Client) *Store {
	return &Store{db: db}
}

func (s *Store) GetAllAds(c *gin.Context) (*Ads, error) {
	opts := options.Find()
	coll := s.db.Database(os.Getenv("MONGO_NAME")).Collection("AdContent")

	ads, err := coll.Find(context.Background(), bson.D{{}}, opts)
	if err != nil {
		log.Fatal("mongo could not read or responde : %w", err)
		return nil, err
	}
	var doc Ads
	for ads.Next(context.Background()) {
		if err := ads.Decode(&doc); err != nil {
			log.Fatal("Error decoding document : %w", err)
			continue
		}

	}
	return &doc, nil
}

func (s *Store) CreateUser(c *gin.Context) (*spec.DefaulCreatedResponse, error) {
	coll := s.db.Database(os.Getenv("MONGO_NAME")).Collection("User")

	var body spec.CreateUserRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Fatal("Error decoding document : %w", err)
		return nil, err
	}
	userExists, err := s.CheckIfUserExistsByEmail(string(*body.Email))
	if err != nil {
		return nil, err
	}
	if userExists {
		return nil, errors.New("user email already exists")
	}

	result, err := coll.InsertOne(context.Background(), body)

	res, err := utils.PrepareIDResponse(result)

	if err != nil {
		log.Fatal("Erro inserting object : %w", err)
		return nil, err
	}
	return res, nil
}

func (s *Store) CheckIfUserExistsByEmail(email string) (bool, error) {
	filter := bson.M{"email": email}
	coll := s.db.Database(os.Getenv("MONGO_NAME")).Collection("User")
	var body EmailResult
	err := coll.FindOne(context.Background(), filter).Decode(&body)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		log.Fatalf("Error Reading document")
		return false, errors.New("Error Reading document")
	}

	return true, nil
}

func (s *Store) InsertCars(body *spec.CreateCarRequest) (*spec.DefaulCreatedResponse, error) {
	coll := s.db.Database(os.Getenv("MONGO_NAME")).Collection("Car")
	t := time.Now()
	body.CreatedAt = &t
	body.UpdatedAt = &t
	plateExists, err := s.CarExistsByPlate(body.Plate)
	if err != nil {
		return nil, err
	}
	if plateExists {
		return nil, errors.New("user Plate already exists")
	}

	result, err := coll.InsertOne(context.Background(), body)
	if err != nil {
		log.Fatal("Error inserting data : %w", err)
		return nil, err
	}
	res, err := utils.PrepareIDResponse(result)

	if err != nil {
		log.Fatal("Erro inserting object : %w", err)
		return nil, err
	}
	return res, nil

}

func (s *Store) CarExistsByPlate(plate *string) (bool, error) {
	filter := bson.M{"plate": plate}
	coll := s.db.Database(os.Getenv("MONGO_NAME")).Collection("Car")
	var body PlateResult
	err := coll.FindOne(context.Background(), filter).Decode(&body)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		log.Fatalf("Error Reading document")
		return false, errors.New("Error Reading document")
	}
	return true, err
}

func (s *Store) FindCarById(objectID *primitive.ObjectID) (*spec.GetCarResponse, error) {
	coll := s.db.Database(os.Getenv("MONGO_NAME")).Collection("Car")

	var res spec.GetCarResponse
	err := coll.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil

}

func (s *Store) ListAllCars(params spec.GetCarsParams) (*spec.GetCarsListResponse, error) {

	coll := s.db.Database(os.Getenv("MONGO_NAME")).Collection("Car")

	filter := bson.D{}
	if params.IsAvailable != nil {
		filter = append(filter, bson.E{Key: "isAvailable", Value: *params.IsAvailable})
	}
	if params.ValueDateFrom != nil && params.ValueDateTo != nil {
		valueDateFrom, err := time.Parse(time.RFC3339, params.ValueDateFrom.String())
		if err != nil {
			return nil, err
		}
		filter = append(filter, bson.E{Key: "valueDate", Value: bson.D{{"$gte", valueDateFrom}}})
		valueDateTo, err := time.Parse(time.RFC3339, params.ValueDateTo.String())
		if err != nil {
			return nil, err
		}
		filter = append(filter, bson.E{Key: "valueDate", Value: bson.D{{"$lte", valueDateTo}}})

	}

	var result spec.GetCarsListResponse

	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var carsList []spec.GetCarResponse
	for cursor.Next(context.Background()) {
		var car spec.GetCarResponse
		err := cursor.Decode(&car)
		if err != nil {
			return nil, err
		}
		carsList = append(carsList, car)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	result.Data = &carsList
	return &result, nil
}

func (s *Store) DeleteCarDB(objectID *primitive.ObjectID) error {
	coll := s.db.Database(os.Getenv("MONGO_NAME")).Collection("Car")

	filter := bson.D{{Key: "_id", Value: objectID}}
	var body spec.GetCarResponse
	if err := coll.FindOne(context.Background(), filter).Decode(&body); err != nil {
		return err
	}

	result, err := coll.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return err
	}
	return nil

}

func (s *Store) UpdateCarEntity(objectID *primitive.ObjectID, body *spec.UpdateCarRequest) error {
	coll := s.db.Database(os.Getenv("MONGO_NAME")).Collection("Car")

	updateDoc := bson.D{}
	val := reflect.ValueOf(body).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		fieldName := fieldType.Tag.Get("bson")

		if fieldName == "" || fieldName == "-" {
			continue
		}
		if !field.IsZero() {
			updateDoc = append(updateDoc, bson.E{Key: fieldName, Value: field.Interface()})

		}
	}
	updateDoc = append(updateDoc, bson.E{Key: "updated_at", Value: time.Now()})
	if len(updateDoc) > 0 {
		_, err := coll.UpdateOne(
			context.Background(),
			bson.M{"_id": objectID},
			bson.D{{"$set", updateDoc}},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Store) UpdateUserEntity(objectID *primitive.ObjectID, user *spec.UpdateUserRequest) error {
	coll := s.db.Database(os.Getenv("MONGO_NAME")).Collection("User")

	updateDoc := bson.D{}
	val := reflect.ValueOf(user).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		fieldName := fieldType.Tag.Get("bson")

		if fieldName == "" || fieldName == "-" {
			continue
		}
		if !field.IsZero() {
			updateDoc = append(updateDoc, bson.E{Key: fieldName, Value: field.Interface()})
		}
	}
	updateDoc = append(updateDoc, bson.E{Key: "updated_at", Value: time.Now()})

	if len(updateDoc) > 0 {
		_, err := coll.UpdateOne(
			context.Background(),
			bson.M{"_id": objectID},
			bson.D{{Key: "$set", Value: updateDoc}},
		)
		if err != nil {
			return err
		}
	}
	return nil
}
