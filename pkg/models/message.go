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

type Message struct {
	ID        int64              `json:"id" bson:"_id"`
	CreatedAt string             `json:"created_at" bson:"created_at"`
	Parts     []MessagePart      `json:"parts" bson:"parts"`
	RoomID    primitive.ObjectID `json:"room_id" bson:"room_id"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
}

type MessagePart struct {
	Content    string     `json:"content" bson:"content"`
	Type       string     `json:"type" bson:"type"`
	URL        string     `json:"url" bson:"url"`
	Attachment Attachment `json:"attachment" bson:"attachment"`
}

type Attachment struct {
	ID          primitive.ObjectID `json:"id" bson:"id"`
	CustomData  json.RawMessage    `json:"custom_data" bson:"custom_data"`
	DownloadURL string             `json:"download_url" bson:"download_url"`
	Expiration  string             `json:"expiration" bson:"expiration"`
	Name        string             `json:"name" bson:"name"`
	RefreshURL  string             `json:"refresh_url" bson:"refresh_url"`
	Size        int64              `json:"size" bson:"size"`
}
