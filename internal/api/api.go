package api

import (
	"net/http"

	"github.com/bibipe2024/goautdoor/internal/api/spec"
	"github.com/bibipe2024/goautdoor/internal/mongostore"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type store interface {
	GetAllAds(c *gin.Context) (*mongostore.Ads, error)
	CreateUser(c *gin.Context) (*spec.DefaulCreatedResponse, error)
	InsertCars(body *spec.CreateCarRequest) (*spec.DefaulCreatedResponse, error)
	FindCarById(objectID *primitive.ObjectID) (*spec.GetCarResponse, error)
	ListAllCars(params spec.GetCarsParams) (*spec.GetCarsListResponse, error)
	DeleteCarDB(objectID *primitive.ObjectID) error
	UpdateCarEntity(objectID *primitive.ObjectID, body *spec.UpdateCarRequest) error
	UpdateUserEntity(objectID *primitive.ObjectID, user *spec.UpdateUserRequest) error
}

type API struct {
	store  store
	logger *zap.Logger
}

func NewAPI(pool *mongo.Client, logger *zap.Logger) API {
	return API{mongostore.New(pool), logger}
}

// Get all list texts of ads.
// (GET /ads)
func (api API) GetAds(c *gin.Context) {
	res, err := api.store.GetAllAds(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, spec.Error{Message: "Could not get ads"})
		return
	}
	c.JSON(200, res)

}

// create user auth login
// (POST /auth/create-login)
func (api API) PostAuthCreateLogin(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

// logs user in
// (POST /auth/login)
func (api API) PostAuthLogin(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

// Get list of All cars
// (GET /cars)
func (api API) GetCars(c *gin.Context, params spec.GetCarsParams) {

	api.logger.Debug("params ", zap.Any("IsAvailable", params.IsAvailable))
	api.logger.Debug("params ", zap.Any("PageSize", params.PageSize))
	api.logger.Debug("params ", zap.Any("ValueDateFrom", params.ValueDateFrom))
	api.logger.Debug("params ", zap.Any("ValueDateTo", params.ValueDateTo))
	res, err := api.store.ListAllCars(params)
	if err != nil {
		api.logger.Error("Erro na consulta", zap.Error(err))
		c.JSON(http.StatusBadGateway, spec.Error{Message: "can not get data"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Store a car in db
// (POST /cars)
func (api API) PostCars(c *gin.Context) {

	var body spec.CreateCarRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		api.logger.Error("Error decoding document ", zap.Error(err))
		c.JSON(http.StatusBadRequest, spec.Error{Message: "could not decode document"})
		return
	}
	result, err := api.store.InsertCars(&body)
	if err != nil {
		api.logger.Error("Error storing car ", zap.Error(err))
		c.JSON(http.StatusBadRequest, spec.Error{Message: "could not insert car"})
		return
	}
	c.JSON(200, result)
}

// delete a car by id
// (DELETE /cars/{carId})
func (api API) DeleteCarsCarId(c *gin.Context, carId string) {
	stringID := c.Param("carId")
	objectID, err := primitive.ObjectIDFromHex(stringID)
	if err != nil {
		c.JSON(http.StatusBadRequest, spec.Error{Message: "ID error, could not encode"})
		return
	}
	if err := api.store.DeleteCarDB(&objectID); err != nil {
		api.logger.Error("Error deleting car", zap.Any("CarID", objectID))
		c.JSON(http.StatusBadGateway, zap.Error(err))
		return
	}

	c.JSON(http.StatusNoContent, nil)

}

// get a car by id
// (GET /cars/{carId})
func (api API) GetCarsCarId(c *gin.Context, carId string) {
	stringID := c.Param("carId")
	api.logger.Debug("key ", zap.String("carId", stringID))
	objectID, err := primitive.ObjectIDFromHex(stringID)
	api.logger.Debug("MongoDB ObjectID", zap.Any("objectID", objectID))
	if err != nil {
		c.JSON(http.StatusBadRequest, spec.Error{Message: "ID error, could not encode"})
		return
	}
	result, err := api.store.FindCarById(&objectID)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, spec.Error{Message: "Not Found"})
			return
		}
		api.logger.Error("Server error ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, spec.Error{Message: "Could not respond"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// update a car by id
// (PATCH /cars/{carId})
func (api API) PatchCarsCarId(c *gin.Context, carId string) {
	stringID := c.Param("carId")
	api.logger.Debug("key ", zap.String("carId", stringID))
	objectID, err := primitive.ObjectIDFromHex(stringID)
	if err != nil {
		c.JSON(http.StatusBadRequest, spec.Error{Message: "ID error, could not encode"})
		return
	}
	var body spec.UpdateCarRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		api.logger.Error("Error decoding document ", zap.Error(err))
		c.JSON(http.StatusBadRequest, spec.Error{Message: "could not decode document"})
		return
	}
	err = api.store.UpdateCarEntity(&objectID, &body)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, spec.Error{Message: "Not Found"})
			return
		}
		api.logger.Error("Server error ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, spec.Error{Message: "Could not respond"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// Upload multiple images
// (POST /cars/{carId}/upload-multiple)
func (api API) PostCarsCarIdUploadMultiple(c *gin.Context, carId string) {
	panic("not implemented") // TODO: Implement
}

// create user with all
// (POST /users)
func (api API) PostUsers(c *gin.Context) {
	res, err := api.store.CreateUser(c)
	if err != nil {
		api.logger.Error("bad request", zap.Error(err))
		c.JSON(http.StatusBadRequest, spec.Error{Message: err.Error()})
		return
	}
	c.JSON(200, res)
}

// update user
// (PUT /users/{userId})
func (api API) PutUsersUserId(c *gin.Context, userId string) {
	stringID := c.Param("userId")
	api.logger.Debug("key ", zap.String("carId", stringID))
	objectID, err := primitive.ObjectIDFromHex(stringID)
	if err != nil {
		c.JSON(http.StatusBadRequest, spec.Error{Message: "ID error, could not encode"})
		return
	}
	var user spec.UpdateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		api.logger.Error("Error decoding document ", zap.Error(err))
		c.JSON(http.StatusBadRequest, spec.Error{Message: "could not decode document"})
		return
	}
	err = api.store.UpdateUserEntity(&objectID, &user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, spec.Error{Message: "Not Found"})
			return
		}
		api.logger.Error("Server error ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, spec.Error{Message: "Could not respond"})
		return
	}
	c.JSON(http.StatusNoContent, nil)

}

// upload user's image
// (POST /users/{userId}/image/upload)
func (api API) PostUsersUserIdImageUpload(c *gin.Context, userId string) {
	panic("not implemented") // TODO: Implement
}
