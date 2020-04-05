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

type Room struct {
	ID                            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name                          string             `json:"name" bson:"name"`
	Private                       bool               `json:"private" bson:"private"`
	PushNotificationTitleOverride string             `json:"push_notification_title_override" bson:"push_notification_title_override"`
	CreatedByID                   string             `json:"created_by_id" bson:"created_by_id"`
	LastMessageAt                 primitive.DateTime `json:"last_message_at" bson:"last_message_at"`
	UpdatedAt                     primitive.DateTime `json:"updated_at" bson:"updated_at"`
	CreatedAt                     primitive.DateTime `json:"created_at" bson:"created_at"`
	CustomData                    json.RawMessage    `json:"custom_data" bson:"custom_data"`
}

type Membership struct {
	RoomID  primitive.ObjectID `json:"room_id"`
	UserIds []string           `json:"user_ids"`
}

type RS struct {
	Cursor struct {
		CursorType int64  `json:"cursor_type"`
		Position   int64  `json:"position"`
		RoomID     string `json:"room_id"`
		UpdatedAt  string `json:"updated_at"`
		UserID     string `json:"user_id"`
	} `json:"cursor"`
	RoomID      string `json:"room_id"`
	UnreadCount int64  `json:"unread_count"`
}
