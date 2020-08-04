package auth

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"khanhnguyen234/api-gateway/_mongo"
	"khanhnguyen234/api-gateway/common"
)

const (
	collection = "auth"
)

type AuthModel struct {
	ID     			primitive.ObjectID 	`bson:"_id,omitempty"`
	Phone  			string             	`bson:"phone,omitempty"`
	PasswordHash 	string             	`bson:"password_hash,omitempty"`
	Email   		string           	`bson:"email,omitempty"`
}

func (u *AuthModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

func (u *AuthModel) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func InsertAuth(data interface{}) error {
	return _mongo.InsertOne(collection, data)
}

func FindOneUser(email string) (AuthModel, error) {
	db := _mongo.ConnectMongo()
	c := db.Collection(collection)

	var authModel AuthModel
	//condition := bson.D{{"email", email}}
	condition := bson.M{"email": email}

	ctx := common.GetContext()
	err := c.FindOne(ctx, condition).Decode(&authModel)

	return authModel, err
}