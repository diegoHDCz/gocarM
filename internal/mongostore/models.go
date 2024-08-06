package mongostore

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ads struct {
	ID    primitive.ObjectID `bson:"_id"`
	Code  int64              `bson:"code"`
	Title string             `bson:"title"`
	Texts []AdContent        `bson:"texts"`
}

type AdContent struct {
	Code       int64          `bson:"code"`
	Image      string         `bson:"image"`
	Subcontent []AdSubContent `bson:"subtexts"`
	Text       string         `bson:"text"`
	Url        string         `bson:"url"`
}

type AdSubContent struct {
	Code  int    `bson:"code"`
	Image string `bson:"image"`
	Text  string `bson:"text"`
	Url   string `bson:"url"`
}

type EmailResult struct {
	Email string `bson:"email"`
}

type PlateResult struct {
	Plate string `bson:"plate"`
}
