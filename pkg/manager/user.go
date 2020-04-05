/*
Copyright © 2020 Alexander Kiryukhin <a.kiryukhin@mail.ru>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package manager

import (
	"context"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/neonxp/chatcloud/pkg/db"
	"github.com/neonxp/chatcloud/pkg/models"
)

type User struct {
	manager *db.Manager
}

func NewUser(collection *mongo.Collection) (*User, error) {
	manager, err := db.NewManager(collection, nil)
	if err != nil {
		return nil, err
	}
	return &User{
		manager: manager,
	}, nil
}

func (m *User) CreateUser(id string, name string, avatarUrl string, customData interface{}) (*models.User, error) {
	bCustomData, err := json.Marshal(customData)
	if err != nil {
		return nil, err
	}
	u := &models.User{
		ID:         id,
		Name:       name,
		AvatarURL:  avatarUrl,
		CustomData: bCustomData,
		CreatedAt:  primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:  primitive.NewDateTimeFromTime(time.Now()),
	}
	if _, err := m.manager.Add(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (m *User) FindByID(id string) (*models.User, error) {
	u := new(models.User)
	return u, m.manager.FindOne(bson.M{"_id": id}, u)
}
func (m *User) FindByIDs(ids []string) ([]*models.User, error) {
	cur, err := m.manager.Find(
		bson.M{"_id": ids},
		map[string]int{"created_at": -1},
		db.Pagination{Offset: 0, Limit: 0},
	)
	if err != nil {
		return nil, err
	}
	if cur == nil {
		return nil, nil
	}
	defer cur.Close(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var users []*models.User
	for cur.Next(ctx) {
		u := new(models.User)
		if err := cur.Decode(u); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (m *User) Find(fromTS time.Time, limit int) ([]*models.User, error) {
	filter := bson.M{}
	if !fromTS.IsZero() {
		filter = bson.M{
			"created_at": bson.M{
				"$gt": primitive.NewDateTimeFromTime(fromTS),
			},
		}
	}
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	cur, err := m.manager.Find(
		filter,
		map[string]int{"created_at": -1},
		db.Pagination{
			Offset: 0,
			Limit:  int64(limit),
		},
	)
	if err != nil {
		return nil, err
	}
	if cur == nil {
		return nil, nil
	}
	defer cur.Close(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var users []*models.User
	for cur.Next(ctx) {
		u := new(models.User)
		if err := cur.Decode(u); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
