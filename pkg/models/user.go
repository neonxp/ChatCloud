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
package models

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         string             `json:"id" bson:"_id,omitempty"`
	AvatarURL  string             `json:"avatar_url" bson:"avatar_url"`
	CustomData json.RawMessage    `json:"custom_data" bson:"custom_data"`
	Name       string             `json:"name" bson:"name"`
	CreatedAt  primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt  primitive.DateTime `json:"updated_at" bson:"updated_at"`
}
