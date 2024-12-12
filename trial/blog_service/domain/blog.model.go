package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"practice/user_service/domain"
)

type Blog struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Content   string             `json:"content" bson:"content"`
	Tag       []string           `json:"tag" bson:"tag"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	Owner     string             `json:"owner" bson:"owner"`
}
type PostBlog struct {
	Owner   domain.ResponseUser `json:"owner,omitempty" bson:"owner,omitempty "`
	Title   string   `json:"title" bson:"title"`
	Content string   `json:"content" bson:"content"`
	Tag     []string `json:"tag" bson:"tag"`
}
type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
