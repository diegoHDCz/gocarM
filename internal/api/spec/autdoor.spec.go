// Package spec provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package spec

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for CreateCarRequestCategory.
const (
	CreateCarRequestCategoryClassic CreateCarRequestCategory = "Classic"
	CreateCarRequestCategoryLuxury  CreateCarRequestCategory = "Luxury"
	CreateCarRequestCategoryMoto    CreateCarRequestCategory = "Moto"
	CreateCarRequestCategoryPopular CreateCarRequestCategory = "Popular"
	CreateCarRequestCategorySUV     CreateCarRequestCategory = "SUV"
	CreateCarRequestCategorySports  CreateCarRequestCategory = "Sports"
)

// Defines values for CreateCarRequestWeekDaysAvailable.
const (
	CreateCarRequestWeekDaysAvailableFRIDAY    CreateCarRequestWeekDaysAvailable = "FRIDAY"
	CreateCarRequestWeekDaysAvailableMONDAY    CreateCarRequestWeekDaysAvailable = "MONDAY"
	CreateCarRequestWeekDaysAvailableSATURDAY  CreateCarRequestWeekDaysAvailable = "SATURDAY"
	CreateCarRequestWeekDaysAvailableSUNDAY    CreateCarRequestWeekDaysAvailable = "SUNDAY"
	CreateCarRequestWeekDaysAvailableTHURSDAY  CreateCarRequestWeekDaysAvailable = "THURSDAY"
	CreateCarRequestWeekDaysAvailableTUESDAY   CreateCarRequestWeekDaysAvailable = "TUESDAY"
	CreateCarRequestWeekDaysAvailableWEDNESDAY CreateCarRequestWeekDaysAvailable = "WEDNESDAY"
)

// Defines values for CreateUserRequestDriversPermitCategory.
const (
	CreateUserRequestDriversPermitCategoryA CreateUserRequestDriversPermitCategory = "A"
	CreateUserRequestDriversPermitCategoryB CreateUserRequestDriversPermitCategory = "B"
	CreateUserRequestDriversPermitCategoryC CreateUserRequestDriversPermitCategory = "C"
	CreateUserRequestDriversPermitCategoryD CreateUserRequestDriversPermitCategory = "D"
)

// Defines values for CreateUserRequestRole.
const (
	ADMIN  CreateUserRequestRole = "ADMIN"
	HOSTER CreateUserRequestRole = "HOSTER"
	RENTER CreateUserRequestRole = "RENTER"
)

// Defines values for GetCarResponseCategory.
const (
	GetCarResponseCategoryClassic GetCarResponseCategory = "Classic"
	GetCarResponseCategoryLuxury  GetCarResponseCategory = "Luxury"
	GetCarResponseCategoryMoto    GetCarResponseCategory = "Moto"
	GetCarResponseCategoryPopular GetCarResponseCategory = "Popular"
	GetCarResponseCategorySUV     GetCarResponseCategory = "SUV"
	GetCarResponseCategorySports  GetCarResponseCategory = "Sports"
)

// Defines values for GetCarResponseWeekDaysRent.
const (
	GetCarResponseWeekDaysRentFRIDAY    GetCarResponseWeekDaysRent = "FRIDAY"
	GetCarResponseWeekDaysRentMONDAY    GetCarResponseWeekDaysRent = "MONDAY"
	GetCarResponseWeekDaysRentSATURDAY  GetCarResponseWeekDaysRent = "SATURDAY"
	GetCarResponseWeekDaysRentSUNDAY    GetCarResponseWeekDaysRent = "SUNDAY"
	GetCarResponseWeekDaysRentTHURSDAY  GetCarResponseWeekDaysRent = "THURSDAY"
	GetCarResponseWeekDaysRentTUESDAY   GetCarResponseWeekDaysRent = "TUESDAY"
	GetCarResponseWeekDaysRentWEDNESDAY GetCarResponseWeekDaysRent = "WEDNESDAY"
)

// Defines values for UpdateCarRequestWeekDaysAvailable.
const (
	FRIDAY    UpdateCarRequestWeekDaysAvailable = "FRIDAY"
	MONDAY    UpdateCarRequestWeekDaysAvailable = "MONDAY"
	SATURDAY  UpdateCarRequestWeekDaysAvailable = "SATURDAY"
	SUNDAY    UpdateCarRequestWeekDaysAvailable = "SUNDAY"
	THURSDAY  UpdateCarRequestWeekDaysAvailable = "THURSDAY"
	TUESDAY   UpdateCarRequestWeekDaysAvailable = "TUESDAY"
	WEDNESDAY UpdateCarRequestWeekDaysAvailable = "WEDNESDAY"
)

// Defines values for UpdateUserRequestDriversPermitCategory.
const (
	UpdateUserRequestDriversPermitCategoryA UpdateUserRequestDriversPermitCategory = "A"
	UpdateUserRequestDriversPermitCategoryB UpdateUserRequestDriversPermitCategory = "B"
	UpdateUserRequestDriversPermitCategoryC UpdateUserRequestDriversPermitCategory = "C"
	UpdateUserRequestDriversPermitCategoryD UpdateUserRequestDriversPermitCategory = "D"
)

// AdContent defines model for AdContent.
type AdContent struct {
	Code       *int            `json:"code,omitempty"`
	Image      *string         `json:"image,omitempty"`
	Subcontent *[]AdSubContent `json:"subcontent,omitempty"`
	Text       *string         `json:"text,omitempty"`
	Url        *string         `json:"url,omitempty"`
}

// AdSubContent defines model for AdSubContent.
type AdSubContent struct {
	Code  *int    `json:"code,omitempty"`
	Image *string `json:"image,omitempty"`
	Text  *string `json:"text,omitempty"`
	Url   *string `json:"url,omitempty"`
}

// CreateAuthUserRequest defines model for CreateAuthUserRequest.
type CreateAuthUserRequest struct {
	Email    *openapi_types.Email `bson:"email" json:"email,omitempty"`
	Name     *string              `bson:"name" json:"name,omitempty"`
	Password *string              `bson:"password" json:"password,omitempty"`
	Phone    *[]string            `bson:"phone" json:"phone,omitempty"`
}

// CreateCarRequest defines model for CreateCarRequest.
type CreateCarRequest struct {
	Brand             *string                              `bson:"brand" json:"brand,omitempty"`
	CarModel          *string                              `bson:"car_model" json:"car_model,omitempty"`
	Category          *CreateCarRequestCategory            `bson:"category" json:"category,omitempty"`
	City              *string                              `bson:"city" json:"city,omitempty"`
	Color             *string                              `bson:"color" json:"color,omitempty"`
	CreatedAt         *time.Time                           `bson:"created_at" json:"created_at,omitempty"`
	ExtraFeatures     *[]string                            `bson:"extra_features" json:"extra_features,omitempty"`
	ImageUrl          *[]string                            `bson:"imageUrl" json:"imageUrl,omitempty"`
	KmQuantity        *float32                             `bson:"km_quantity" json:"km_quantity,omitempty"`
	Maximum60Days     *bool                                `bson:"maximum_60_days" json:"maximum_60_days,omitempty"`
	OwnerEmail        *openapi_types.Email                 `bson:"owner_email" json:"owner_email,omitempty"`
	OwnerId           *string                              `bson:"owner_id" json:"owner_id,omitempty"`
	Plate             *string                              `bson:"plate" json:"plate,omitempty"`
	State             *string                              `bson:"state" json:"state,omitempty"`
	UpdatedAt         *time.Time                           `bson:"updated_at" json:"updated_at,omitempty"`
	WeekDaysAvailable *[]CreateCarRequestWeekDaysAvailable `bson:"week_days_available" json:"week_days_available,omitempty"`
	YearFabrication   *string                              `bson:"year_fabrication" json:"year_fabrication,omitempty"`
	YearModel         *string                              `bson:"year_model" json:"year_model,omitempty"`
}

// CreateCarRequestCategory defines model for CreateCarRequest.Category.
type CreateCarRequestCategory string

// CreateCarRequestWeekDaysAvailable defines model for CreateCarRequest.WeekDaysAvailable.
type CreateCarRequestWeekDaysAvailable string

// CreateUserRequest defines model for CreateUserRequest.
type CreateUserRequest struct {
	CPF     *string `bson:"cpf" json:"CPF,omitempty"`
	Address *struct {
		CEP          *string `bson:"CEP" json:"CEP,omitempty"`
		City         *string `bson:"city" json:"city,omitempty"`
		Complement   *string `bson:"complement" json:"complement,omitempty"`
		Neighborhood *string `bson:"neighborhood" json:"neighborhood,omitempty"`
		Number       *int    `bson:"number" json:"number,omitempty"`
		State        *string `bson:"state" json:"state,omitempty"`
		Street       *string `bson:"street" json:"street,omitempty"`
	} `bson:"address" json:"address,omitempty"`
	DriversPermit *struct {
		Category     *[]CreateUserRequestDriversPermitCategory `bson:"category" json:"category,omitempty"`
		DueDate      *openapi_types.Date                       `bson:"due_date" json:"due_date,omitempty"`
		PermitNumber *string                                   `bson:"license" json:"permit_number,omitempty"`
	} `bson:"permit" json:"drivers_permit,omitempty"`
	Email      *openapi_types.Email     `bson:"email" json:"email,omitempty"`
	IdDocument *string                  `bson:"id_document" json:"id_document,omitempty"`
	ImageUrl   *string                  `bson:"imageUrl" json:"imageUrl,omitempty"`
	IsActive   *bool                    `bson:"is_active" json:"isActive,omitempty"`
	Name       *string                  `bson:"name" json:"name,omitempty" validate:"required,min=1,max=256"`
	Phone      *[]string                `json:"phone,omitempty"`
	Role       *[]CreateUserRequestRole `json:"role,omitempty"`
}

// CreateUserRequestDriversPermitCategory defines model for CreateUserRequest.DriversPermit.Category.
type CreateUserRequestDriversPermitCategory string

// CreateUserRequestRole defines model for CreateUserRequest.Role.
type CreateUserRequestRole string

// DefaulCreatedResponse defines model for DefaulCreatedResponse.
type DefaulCreatedResponse struct {
	Id *ObjectId `bson:"_id" json:"_id,omitempty"`
}

// Error Bad request
type Error struct {
	Message string `json:"message"`
}

// GetAllAdsResponse defines model for GetAllAdsResponse.
type GetAllAdsResponse struct {
	Id    *ObjectId    `bson:"_id" json:"_id,omitempty"`
	Texts *[]AdContent `json:"texts,omitempty"`
	Title *string      `json:"title,omitempty"`
}

// GetCarResponse defines model for GetCarResponse.
type GetCarResponse struct {
	Brand             *string                       `bson:"brand" json:"brand,omitempty"`
	CarModel          *string                       `bson:"car_model" json:"car_model,omitempty"`
	Category          *GetCarResponseCategory       `bson:"category" json:"category,omitempty"`
	City              *string                       `bson:"city" json:"city,omitempty"`
	Color             *string                       `bson:"color" json:"color,omitempty"`
	ExtraFeatures     *[]string                     `bson:"extra_features" json:"extra_features,omitempty"`
	ImageUrl          *[]string                     `bson:"imageUrl" json:"imageUrl,omitempty"`
	IsVailableForRent *bool                         `bson:"isVailableForRent" json:"isVailableForRent,omitempty"`
	KmQuantity        *float32                      `bson:"km_quantity" json:"km_quantity,omitempty"`
	OwnerEmail        *openapi_types.Email          `bson:"owner_email" json:"owner_email,omitempty"`
	OwnerId           *string                       `bson:"owner_id" json:"owner_id,omitempty"`
	Plate             *string                       `bson:"plate" json:"plate,omitempty"`
	ReservedDate      *[]openapi_types.Date         `bson:"reserved_date" json:"reserved_date,omitempty"`
	State             *string                       `bson:"state" json:"state,omitempty"`
	WeekDaysRent      *[]GetCarResponseWeekDaysRent `bson:"week_days_rent" json:"week_days_rent,omitempty"`
	YearFabrication   *string                       `bson:"year_fabrication" json:"year_fabrication,omitempty"`
	YearModel         *string                       `bson:"year_model" json:"year_model,omitempty"`
}

// GetCarResponseCategory defines model for GetCarResponse.Category.
type GetCarResponseCategory string

// GetCarResponseWeekDaysRent defines model for GetCarResponse.WeekDaysRent.
type GetCarResponseWeekDaysRent string

// GetCarsListResponse defines model for GetCarsListResponse.
type GetCarsListResponse struct {
	Data *[]GetCarResponse `bson:"data" json:"data,omitempty"`
}

// GetUserToken defines model for GetUserToken.
type GetUserToken struct {
	Email    *openapi_types.Email `bson:"email" json:"email,omitempty"`
	Password *string              `bson:"password" json:"password,omitempty"`
}

// InsertMultipleCarImages defines model for InsertMultipleCarImages.
type InsertMultipleCarImages struct {
	CarId  *ObjectId `bson:"_id" json:"carId,omitempty"`
	Images *[]struct {
		Binary *openapi_types.File `json:"binary,omitempty"`
		Name   *string             `json:"name,omitempty"`
	} `json:"images,omitempty"`
	OwnerEmail *openapi_types.Email `json:"owner_email,omitempty"`
}

// InsertUserImage defines model for InsertUserImage.
type InsertUserImage struct {
	Image *struct {
		Binary *openapi_types.File `json:"binary,omitempty"`
		Name   *string             `json:"name,omitempty"`
	} `json:"image,omitempty"`
	OwnerEmail *string   `json:"owner_email,omitempty"`
	UserId     *ObjectId `bson:"_id" json:"userId,omitempty"`
}

// ObjectId defines model for ObjectId.
type ObjectId = string

// UpdateCarRequest defines model for UpdateCarRequest.
type UpdateCarRequest struct {
	City                 *string                              `bson:"city" json:"city,omitempty"`
	ExtraFeatures        *[]string                            `bson:"extra_features" json:"extra_features,omitempty"`
	KmQuantity           *float32                             `bson:"km_quantity" json:"km_quantity,omitempty"`
	Maximum60Days        *bool                                `bson:"maximum_60_days" json:"maximum_60_days,omitempty"`
	UpdatedAt            *time.Time                           `bson:"updated_at" json:"updated_at,omitempty"`
	WeekDaysAvailable    *[]UpdateCarRequestWeekDaysAvailable `bson:"week_days_available" json:"week_days_available,omitempty"`
	AdditionalProperties map[string]interface{}               `json:"-"`
}

// UpdateCarRequestWeekDaysAvailable defines model for UpdateCarRequest.WeekDaysAvailable.
type UpdateCarRequestWeekDaysAvailable string

// UpdateUserRequest defines model for UpdateUserRequest.
type UpdateUserRequest struct {
	Address *struct {
		CEP          *string `bson:"CEP" json:"CEP,omitempty"`
		City         *string `bson:"city" json:"city,omitempty"`
		Complement   *string `bson:"complement" json:"complement,omitempty"`
		Neighborhood *string `bson:"neighborhood" json:"neighborhood,omitempty"`
		Number       *int    `bson:"number" json:"number,omitempty"`
		State        *string `bson:"state" json:"state,omitempty"`
		Street       *string `bson:"street" json:"street,omitempty"`
	} `bson:"address" json:"address,omitempty"`
	DriversPermit *struct {
		Category     *[]UpdateUserRequestDriversPermitCategory `bson:"category" json:"category,omitempty"`
		DueDate      *openapi_types.Date                       `bson:"due_date" json:"due_date,omitempty"`
		PermitNumber *string                                   `bson:"license" json:"permit_number,omitempty"`
	} `bson:"permit" json:"drivers_permit,omitempty"`
	IdDocument *string   `bson:"id_document" json:"id_document,omitempty"`
	IsActive   *bool     `bson:"is_active" json:"isActive,omitempty"`
	Name       *string   `bson:"name" json:"name,omitempty"`
	Phone      *[]string `bson:"phone" json:"phone,omitempty"`
}

// UpdateUserRequestDriversPermitCategory defines model for UpdateUserRequest.DriversPermit.Category.
type UpdateUserRequestDriversPermitCategory string

// UserTokenResponse defines model for UserTokenResponse.
type UserTokenResponse struct {
	Token *string `bson:"token" json:"token,omitempty"`
}

// IsAvailable defines model for IsAvailable.
type IsAvailable = bool

// PageSize defines model for PageSize.
type PageSize = int32

// ValueDateFrom defines model for ValueDateFrom.
type ValueDateFrom = openapi_types.Date

// ValueDateTo defines model for ValueDateTo.
type ValueDateTo = openapi_types.Date

// GetCarsParams defines parameters for GetCars.
type GetCarsParams struct {
	// PageSize the page size indicates the number of the results in each page
	PageSize      *PageSize      `form:"pageSize,omitempty" json:"pageSize,omitempty"`
	IsAvailable   *IsAvailable   `form:"isAvailable,omitempty" json:"isAvailable,omitempty"`
	ValueDateTo   *ValueDateTo   `form:"valueDateTo,omitempty" json:"valueDateTo,omitempty"`
	ValueDateFrom *ValueDateFrom `form:"valueDateFrom,omitempty" json:"valueDateFrom,omitempty"`
}

// PostAuthCreateLoginJSONRequestBody defines body for PostAuthCreateLogin for application/json ContentType.
type PostAuthCreateLoginJSONRequestBody = CreateAuthUserRequest

// PostAuthLoginJSONRequestBody defines body for PostAuthLogin for application/json ContentType.
type PostAuthLoginJSONRequestBody = GetUserToken

// PostCarsJSONRequestBody defines body for PostCars for application/json ContentType.
type PostCarsJSONRequestBody = CreateCarRequest

// PatchCarsCarIdJSONRequestBody defines body for PatchCarsCarId for application/json ContentType.
type PatchCarsCarIdJSONRequestBody = UpdateCarRequest

// PostCarsCarIdUploadMultipleMultipartRequestBody defines body for PostCarsCarIdUploadMultiple for multipart/form-data ContentType.
type PostCarsCarIdUploadMultipleMultipartRequestBody = InsertMultipleCarImages

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody = CreateUserRequest

// PutUsersUserIdJSONRequestBody defines body for PutUsersUserId for application/json ContentType.
type PutUsersUserIdJSONRequestBody = UpdateUserRequest

// PostUsersUserIdImageUploadMultipartRequestBody defines body for PostUsersUserIdImageUpload for multipart/form-data ContentType.
type PostUsersUserIdImageUploadMultipartRequestBody = InsertUserImage

// Getter for additional properties for UpdateCarRequest. Returns the specified
// element and whether it was found
func (a UpdateCarRequest) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for UpdateCarRequest
func (a *UpdateCarRequest) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for UpdateCarRequest to handle AdditionalProperties
func (a *UpdateCarRequest) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["city"]; found {
		err = json.Unmarshal(raw, &a.City)
		if err != nil {
			return fmt.Errorf("error reading 'city': %w", err)
		}
		delete(object, "city")
	}

	if raw, found := object["extra_features"]; found {
		err = json.Unmarshal(raw, &a.ExtraFeatures)
		if err != nil {
			return fmt.Errorf("error reading 'extra_features': %w", err)
		}
		delete(object, "extra_features")
	}

	if raw, found := object["km_quantity"]; found {
		err = json.Unmarshal(raw, &a.KmQuantity)
		if err != nil {
			return fmt.Errorf("error reading 'km_quantity': %w", err)
		}
		delete(object, "km_quantity")
	}

	if raw, found := object["maximum_60_days"]; found {
		err = json.Unmarshal(raw, &a.Maximum60Days)
		if err != nil {
			return fmt.Errorf("error reading 'maximum_60_days': %w", err)
		}
		delete(object, "maximum_60_days")
	}

	if raw, found := object["updated_at"]; found {
		err = json.Unmarshal(raw, &a.UpdatedAt)
		if err != nil {
			return fmt.Errorf("error reading 'updated_at': %w", err)
		}
		delete(object, "updated_at")
	}

	if raw, found := object["week_days_available"]; found {
		err = json.Unmarshal(raw, &a.WeekDaysAvailable)
		if err != nil {
			return fmt.Errorf("error reading 'week_days_available': %w", err)
		}
		delete(object, "week_days_available")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for UpdateCarRequest to handle AdditionalProperties
func (a UpdateCarRequest) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.City != nil {
		object["city"], err = json.Marshal(a.City)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'city': %w", err)
		}
	}

	if a.ExtraFeatures != nil {
		object["extra_features"], err = json.Marshal(a.ExtraFeatures)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'extra_features': %w", err)
		}
	}

	if a.KmQuantity != nil {
		object["km_quantity"], err = json.Marshal(a.KmQuantity)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'km_quantity': %w", err)
		}
	}

	if a.Maximum60Days != nil {
		object["maximum_60_days"], err = json.Marshal(a.Maximum60Days)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'maximum_60_days': %w", err)
		}
	}

	if a.UpdatedAt != nil {
		object["updated_at"], err = json.Marshal(a.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'updated_at': %w", err)
		}
	}

	if a.WeekDaysAvailable != nil {
		object["week_days_available"], err = json.Marshal(a.WeekDaysAvailable)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'week_days_available': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all list texts of ads.
	// (GET /ads)
	GetAds(c *gin.Context)
	// create user auth login
	// (POST /auth/create-login)
	PostAuthCreateLogin(c *gin.Context)
	// logs user in
	// (POST /auth/login)
	PostAuthLogin(c *gin.Context)
	// Get list of All cars
	// (GET /cars)
	GetCars(c *gin.Context, params GetCarsParams)
	// Store a car in db
	// (POST /cars)
	PostCars(c *gin.Context)
	// delete a car by id
	// (DELETE /cars/{carId})
	DeleteCarsCarId(c *gin.Context, carId string)
	// get a car by id
	// (GET /cars/{carId})
	GetCarsCarId(c *gin.Context, carId string)
	// update a car by id
	// (PATCH /cars/{carId})
	PatchCarsCarId(c *gin.Context, carId string)
	// Upload multiple images
	// (POST /cars/{carId}/upload-multiple)
	PostCarsCarIdUploadMultiple(c *gin.Context, carId string)
	// create user with all
	// (POST /users)
	PostUsers(c *gin.Context)
	// update user
	// (PUT /users/{userId})
	PutUsersUserId(c *gin.Context, userId string)
	// upload user's image
	// (POST /users/{userId}/image/upload)
	PostUsersUserIdImageUpload(c *gin.Context, userId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetAds operation middleware
func (siw *ServerInterfaceWrapper) GetAds(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetAds(c)
}

// PostAuthCreateLogin operation middleware
func (siw *ServerInterfaceWrapper) PostAuthCreateLogin(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostAuthCreateLogin(c)
}

// PostAuthLogin operation middleware
func (siw *ServerInterfaceWrapper) PostAuthLogin(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostAuthLogin(c)
}

// GetCars operation middleware
func (siw *ServerInterfaceWrapper) GetCars(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetCarsParams

	// ------------- Optional query parameter "pageSize" -------------

	err = runtime.BindQueryParameter("form", true, false, "pageSize", c.Request.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter pageSize: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "isAvailable" -------------

	err = runtime.BindQueryParameter("form", true, false, "isAvailable", c.Request.URL.Query(), &params.IsAvailable)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter isAvailable: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "valueDateTo" -------------

	err = runtime.BindQueryParameter("form", true, false, "valueDateTo", c.Request.URL.Query(), &params.ValueDateTo)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter valueDateTo: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "valueDateFrom" -------------

	err = runtime.BindQueryParameter("form", true, false, "valueDateFrom", c.Request.URL.Query(), &params.ValueDateFrom)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter valueDateFrom: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetCars(c, params)
}

// PostCars operation middleware
func (siw *ServerInterfaceWrapper) PostCars(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostCars(c)
}

// DeleteCarsCarId operation middleware
func (siw *ServerInterfaceWrapper) DeleteCarsCarId(c *gin.Context) {

	var err error

	// ------------- Path parameter "carId" -------------
	var carId string

	err = runtime.BindStyledParameterWithOptions("simple", "carId", c.Param("carId"), &carId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter carId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteCarsCarId(c, carId)
}

// GetCarsCarId operation middleware
func (siw *ServerInterfaceWrapper) GetCarsCarId(c *gin.Context) {

	var err error

	// ------------- Path parameter "carId" -------------
	var carId string

	err = runtime.BindStyledParameterWithOptions("simple", "carId", c.Param("carId"), &carId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter carId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetCarsCarId(c, carId)
}

// PatchCarsCarId operation middleware
func (siw *ServerInterfaceWrapper) PatchCarsCarId(c *gin.Context) {

	var err error

	// ------------- Path parameter "carId" -------------
	var carId string

	err = runtime.BindStyledParameterWithOptions("simple", "carId", c.Param("carId"), &carId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter carId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PatchCarsCarId(c, carId)
}

// PostCarsCarIdUploadMultiple operation middleware
func (siw *ServerInterfaceWrapper) PostCarsCarIdUploadMultiple(c *gin.Context) {

	var err error

	// ------------- Path parameter "carId" -------------
	var carId string

	err = runtime.BindStyledParameterWithOptions("simple", "carId", c.Param("carId"), &carId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter carId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostCarsCarIdUploadMultiple(c, carId)
}

// PostUsers operation middleware
func (siw *ServerInterfaceWrapper) PostUsers(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostUsers(c)
}

// PutUsersUserId operation middleware
func (siw *ServerInterfaceWrapper) PutUsersUserId(c *gin.Context) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "userId", c.Param("userId"), &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutUsersUserId(c, userId)
}

// PostUsersUserIdImageUpload operation middleware
func (siw *ServerInterfaceWrapper) PostUsersUserIdImageUpload(c *gin.Context) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "userId", c.Param("userId"), &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostUsersUserIdImageUpload(c, userId)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/ads", wrapper.GetAds)
	router.POST(options.BaseURL+"/auth/create-login", wrapper.PostAuthCreateLogin)
	router.POST(options.BaseURL+"/auth/login", wrapper.PostAuthLogin)
	router.GET(options.BaseURL+"/cars", wrapper.GetCars)
	router.POST(options.BaseURL+"/cars", wrapper.PostCars)
	router.DELETE(options.BaseURL+"/cars/:carId", wrapper.DeleteCarsCarId)
	router.GET(options.BaseURL+"/cars/:carId", wrapper.GetCarsCarId)
	router.PATCH(options.BaseURL+"/cars/:carId", wrapper.PatchCarsCarId)
	router.POST(options.BaseURL+"/cars/:carId/upload-multiple", wrapper.PostCarsCarIdUploadMultiple)
	router.POST(options.BaseURL+"/users", wrapper.PostUsers)
	router.PUT(options.BaseURL+"/users/:userId", wrapper.PutUsersUserId)
	router.POST(options.BaseURL+"/users/:userId/image/upload", wrapper.PostUsersUserIdImageUpload)
}
