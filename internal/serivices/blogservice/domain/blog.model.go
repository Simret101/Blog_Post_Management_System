package domain

import (
	"time"
	"user/domain"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	Admin   Role = "admin"
	Creator Role = "creator"
	Viewer  Role = "viewer"
)

type Blog struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Content   string             `json:"content" bson:"content"`
	Tag       []string           `json:"tag" bson:"tag"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	Owner     interface{}        `json:"owner" bson:"owner"`
}
type PostBlog struct {
	Owner   ResponseUser `json:"owner,omitempty" bson:"owner,omitempty"`
	Title   string       `json:"title" bson:"title"`
	Content string       `json:"content" bson:"content"`
	Tag     []string     `json:"tag" bson:"tag"`
}
type ResponseUser struct {
	ID             string       `json:"_id" bson:"_id"`
	UserName       string       `json:"username" bson:"username"`
	Bio            string       `json:"bio,omitempty" bson:"bio,omitempty"`
	ProfilePicture domain.Media `json:"profile_picture,omitempty" bson:"profile_picture,omitempty"`
	Email          string       `json:"email" bson:"email"`
	Role           Role         `json:"role" bson:"role"`
}

type Claims struct {
	UserID string `json:"user_id"`
	Role   Role   `json:"role"`
	jwt.RegisteredClaims
}

type Tag struct {
	ID string `bson:"_id" json:"id"`
}
type View struct {
	BlogID primitive.ObjectID `bson:"blogid" json:"blogid"`
	User   string             `bson:"user" json:"user"`
}
