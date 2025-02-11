package repository

import (
	"context"

	user "user/domain"
	"user/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UploadRepo struct {
	UserRepository repository.UserRepository
}

func NewUploadRepository(user_repo repository.UserRepository) *UploadRepo {
	return &UploadRepo{
		UserRepository: user_repo,
	}
}

func (repo *UploadRepo) AddProfile(media user.Media, id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	media.ID = primitive.NewObjectID()
	filter := bson.D{{Key: "_id", Value: objID}}
	data := bson.D{{Key: "profile_picture", Value: media}}
	setter := bson.D{{Key: "$set", Value: data}}

	_, err := repo.UserRepository.Collection.UpdateOne(context.TODO(), filter, setter)

	return err
}
