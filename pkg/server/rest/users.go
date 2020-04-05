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
package rest

import (
	"fmt"
	"net/http"
)

type UserRequest struct {
	ID         string      `json:"id"`          // User id assigned to the user in your app.
	Name       string      `json:"name"`        // Name of the new user.
	AvatarURL  string      `json:"avatar_url"`  // A link to the user’s photo/image.
	CustomData interface{} `json:"custom_data"` // Custom data to associate with a user.
}

func (u *UserRequest) Bind(r *http.Request) error {
	if u.ID == "" {
		return fmt.Errorf("`id` is required")
	}
	if u.Name == "" {
		return fmt.Errorf("`name` is required")
	}
	return nil
}

type BatchUsersRequest []*UserRequest

func (u BatchUsersRequest) Bind(r *http.Request) error {
	for idx, uu := range u {
		if uu.ID == "" {
			return fmt.Errorf("%d element: `id` is required", idx)
		}
		if uu.Name == "" {
			return fmt.Errorf("%d element: `name` is required", idx)
		}
	}
	return nil
}
