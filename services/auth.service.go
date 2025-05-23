package services

import (
	"context"

	"github.com/google/go-github/v72/github"
	"github.com/lubie-placki-be/configs"
	"github.com/lubie-placki-be/middlewares"
	"github.com/lubie-placki-be/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetMe() (models.User, error) {
	coll := configs.Client.Database("database").Collection("users")

	user, _, err := middlewares.GithubClient.Users.Get(context.TODO(), "")
	if err != nil {
		return models.User{}, err
	}

	var me models.User
	if err := coll.FindOne(context.TODO(), bson.M{"github_id": user.ID}).Decode(&me); err != nil {
		return models.User{}, err
	}

	return me, nil
}

func CreateUser(githubClient *github.Client) error {
	coll := configs.Client.Database("database").Collection("users")
	opts := options.UpdateOne().SetUpsert(true)

	user, _, err := githubClient.Users.Get(context.TODO(), "")
	if err != nil {
		return err
	}

	var me = models.User{
		GithubID: *user.ID,
		Login:    *user.Login,
		Name:     *user.Name,
	}

	filter := bson.D{
		{Key: "github_id", Value: me.GithubID},
		{Key: "login", Value: me.Login},
		{Key: "name", Value: me.Name},
	}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "github_id", Value: me.GithubID},
			{Key: "login", Value: me.Login},
			{Key: "name", Value: me.Name},
		}},
	}

	if _, err := coll.UpdateOne(context.TODO(), filter, update, opts); err != nil {
		return err
	}

	return nil
}
