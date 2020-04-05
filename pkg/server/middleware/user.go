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
package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/neonxp/chatcloud/pkg"
	"github.com/neonxp/chatcloud/pkg/manager"
	"github.com/neonxp/chatcloud/pkg/models"
)

const userUrlParam = "user_id"
const userCtxKey = "user"

func User(m *manager.User) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			uid := chi.URLParam(r, userUrlParam)
			if uid != "" {
				u, err := m.FindByID(uid)
				if err != nil {
					if err == mongo.ErrNoDocuments {
						pkg.WriteError(w, http.StatusNotFound, fmt.Errorf("user %s not found", uid))
						return
					}
					pkg.WriteError(w, http.StatusInternalServerError, err)
				}
				r = r.WithContext(context.WithValue(
					r.Context(),
					userCtxKey,
					u,
				))
			}
			next.ServeHTTP(w, r)
		})
	}
}

func UserFromRequest(r *http.Request) *models.User {
	return r.Context().Value(userCtxKey).(*models.User)
}
