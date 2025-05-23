package services

import (
	"context"
	"errors"

	"github.com/lubie-placki-be/configs"
	"github.com/lubie-placki-be/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

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

func GetAllRecipes() ([]models.Recipe, error) {
	coll := configs.Client.Database("database").Collection("recipes")

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return []models.Recipe{}, err
	}

	var results []models.Recipe = []models.Recipe{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return []models.Recipe{}, err
	}

	return results, nil
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
