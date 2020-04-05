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
package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/render"

	"github.com/neonxp/chatcloud/pkg"
	"github.com/neonxp/chatcloud/pkg/models"
	mw "github.com/neonxp/chatcloud/pkg/server/middleware"
	"github.com/neonxp/chatcloud/pkg/server/rest"
)

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	req := new(rest.UserRequest)
	if err := render.Bind(r, req); err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err)
		return
	}
	u, err := s.userManager.CreateUser(
		req.ID,
		req.Name,
		req.AvatarURL,
		req.CustomData,
	)
	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err)
		return
	}
	render.JSON(w, r, u)
}

func (s *Server) BatchCreateUsers(w http.ResponseWriter, r *http.Request) {
	req := new(rest.BatchUsersRequest)
	if err := render.Bind(r, req); err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err)
		return
	}
	var resp []*models.User
	//TODO insert many
	for _, request := range *req {
		u, err := s.userManager.CreateUser(
			request.ID,
			request.Name,
			request.AvatarURL,
			request.CustomData,
		)
		if err != nil {
			pkg.WriteError(w, http.StatusBadRequest, err)
			return
		}
		resp = append(resp, u)
	}

	render.JSON(w, r, resp)
}

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	user := mw.UserFromRequest(r)
	render.JSON(w, r, user)
}

func (s *Server) ListUsers(w http.ResponseWriter, r *http.Request) {
	fromTs := r.URL.Query().Get("from_ts")
	limit := r.URL.Query().Get("limit")
	var iFromTs time.Time
	var iLimit int
	var err error
	if fromTs != "" {
		if iFromTs, err = time.Parse(time.RFC3339, fromTs); err != nil {
			pkg.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}
	if limit != "" {
		if iLimit, err = strconv.Atoi(limit); err != nil {
			pkg.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}
	resp, err := s.userManager.Find(iFromTs, iLimit)
	if err != nil {
		pkg.WriteError(w, http.StatusServiceUnavailable, err)
		return
	}
	render.JSON(w, r, resp)
}

func (s *Server) ListUsersByIds(w http.ResponseWriter, r *http.Request) {
	ids := r.URL.Query()["id"]
	resp, err := s.userManager.FindByIDs(ids)
	if err != nil {
		pkg.WriteError(w, http.StatusServiceUnavailable, err)
		return
	}
	render.JSON(w, r, resp)
}
