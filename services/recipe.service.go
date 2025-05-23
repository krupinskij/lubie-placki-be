package services

import (
	"context"
	"errors"
	"time"

	"github.com/lubie-placki-be/configs"
	"github.com/lubie-placki-be/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const LIMIT int = 10

func GetRecipeById(id string) (models.Recipe, error) {
	coll := configs.Client.Database("database").Collection("recipes")

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return models.Recipe{}, err
	}

	var result models.Recipe
	if err := coll.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result); err != nil {
		return models.Recipe{}, err
	}

	return result, nil
}

func GetAllRecipes(page int) (models.Paginated[models.Recipe], error) {
	coll := configs.Client.Database("database").Collection("recipes")

	limit := int64(LIMIT)
	skip := int64((page - 1) * LIMIT)

	filter := bson.D{}
	opts := options.Find().SetLimit(limit).SetSkip(skip).SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		return models.Paginated[models.Recipe]{}, err
	}

	count, err := coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		return models.Paginated[models.Recipe]{}, err
	}

	var results []models.Recipe = []models.Recipe{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return models.Paginated[models.Recipe]{}, err
	}

	countPages := (count-1)/int64(LIMIT) + 1

	return models.Paginated[models.Recipe]{Data: results, Count: count, CountPages: countPages}, nil
}

func GetRandomId() (bson.ObjectID, error) {
	coll := configs.Client.Database("database").Collection("recipes")

	sampleStage := bson.D{
		{Key: "$sample", Value: bson.D{
			{Key: "size", Value: 1},
		}}}
	cursor, err := coll.Aggregate(context.TODO(), mongo.Pipeline{sampleStage})
	if err != nil {
		return bson.ObjectID{}, err
	}

	var results []models.Recipe = []models.Recipe{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return bson.ObjectID{}, err
	}

	if len(results) == 0 {
		return bson.ObjectID{}, errors.New("no recipes")
	}

	return results[0].ID, nil
}

func CreateRecipe(newRecipe models.Recipe) (bson.ObjectID, error) {
	me, err := GetMe()
	if err != nil {
		return bson.ObjectID{}, err
	}

	var recipe = models.Recipe{
		Title:             newRecipe.Title,
		ImageId:           newRecipe.ImageId,
		Time:              newRecipe.Time,
		IngredientsGroups: newRecipe.IngredientsGroups,
		MethodsGroups:     newRecipe.MethodsGroups,
		Author:            me,
		CreatedAt:         time.Now().Unix(),
	}

	coll := configs.Client.Database("database").Collection("recipes")
	result, err := coll.InsertOne(context.TODO(), recipe)
	if err != nil {
		return bson.ObjectID{}, err
	}

	insertedId, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return bson.ObjectID{}, errors.New("id of wrong type")
	}

	return insertedId, nil
}
