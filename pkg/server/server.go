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
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/neonxp/chatcloud/pkg/config"
	"github.com/neonxp/chatcloud/pkg/manager"
	mw "github.com/neonxp/chatcloud/pkg/server/middleware"
)

type Server struct {
	db             *mongo.Database
	cfg            *config.Config
	rds            *redis.Client
	serv           *http.Server
	userManager    *manager.User
	roomManager    *manager.Room
	messageManager *manager.Message
}

func NewServer(db *mongo.Database, rds *redis.Client, cfg *config.Config) (*Server, error) {
	roomManager, err := manager.NewRoom(db.Collection("rooms"))
	if err != nil {
		return nil, err
	}
	userManager, err := manager.NewUser(db.Collection("users"))
	if err != nil {
		return nil, err
	}
	messageManager, err := manager.NewMessage(db.Collection("messages"), rds)
	if err != nil {
		return nil, err
	}
	return &Server{
		db:             db,
		cfg:            cfg,
		rds:            rds,
		serv:           nil,
		userManager:    userManager,
		roomManager:    roomManager,
		messageManager: messageManager,
	}, nil
}

func (s *Server) Init() {
	api := chi.NewRouter()
	api.Use(middleware.RequestID)
	api.Use(middleware.RealIP)
	api.Use(middleware.Logger)
	api.Use(middleware.Recoverer)
	api.Use(middleware.StripSlashes)
	api.Get("/", s.notImplemented)
	api.Route("/api", func(r chi.Router) {
		// Users
		r.Post("/batch_users", s.BatchCreateUsers)
		r.Get("/users_by_ids", s.ListUsersByIds)
		r.Route("/users", func(users chi.Router) {
			users.Use(mw.User(s.userManager))
			users.Get("/", s.ListUsers)
			users.Post("/", s.CreateUser)
			users.Get("/{user_id}", s.GetUser)
			users.Get("/{user_id}/joined_rooms", s.notImplemented)
			users.Get("/{user_id}/joinable_rooms", s.notImplemented)
			users.Post("/{user_id}/join", s.notImplemented)
			users.Post("/{user_id}/leave", s.notImplemented)
			users.Put("/{user_id}", s.notImplemented)
			users.Delete("/{user_id}", s.notImplemented)
			users.Put("/{user_id}/roles", s.notImplemented)
			users.Get("/{user_id}/roles", s.notImplemented)
			users.Delete("/{user_id}/roles", s.notImplemented)
			users.MethodFunc("SUBSCRIBE", "", s.notImplemented)
			users.MethodFunc("SUBSCRIBE", "/{user_id}", s.notImplemented)
			users.MethodFunc("SUBSCRIBE", "/{user_id}/register", s.notImplemented)
		})

		// Rooms
		r.Route("/rooms", func(rooms chi.Router) {
			rooms.Post("/", s.notImplemented)
			rooms.Get("/", s.notImplemented)
			rooms.Get("/{room_id}", s.notImplemented)
			rooms.Put("/{room_id}", s.notImplemented)
			rooms.Delete("/{room_id}", s.notImplemented)
			rooms.Put("/{room_id}/users/add", s.notImplemented)
			rooms.Put("/{room_id}/users/remove", s.notImplemented)
			rooms.Post("/{room_id}/typing_indicators", s.notImplemented)
			rooms.Post("/{room_id}/attachments", s.notImplemented)
			rooms.Get("/{room_id}/messages", s.notImplemented)
			rooms.Post("/{room_id}/messages", s.notImplemented)
			rooms.Get("/{room_id}/messages/{message_id}", s.notImplemented)
			rooms.Put("/{room_id}/messages/{message_id}", s.notImplemented)
			rooms.Delete("/{room_id}/messages/{message_id}", s.notImplemented)
			rooms.Get("/{room_id}/files/{file_name}", s.notImplemented)
			rooms.Delete("/{room_id}/files/{file_name}", s.notImplemented)
			rooms.Post("/{room_id}/users/{user_id}/files/{file_name}", s.notImplemented)
			rooms.Delete("/{room_id}/users/{user_id}/files", s.notImplemented)
			rooms.MethodFunc("SUBSCRIBE", "/{room_id}", s.notImplemented)
		})

		// Roles
		r.Route("/roles", func(roles chi.Router) {
			roles.Get("/", s.notImplemented)
			roles.Post("/", s.notImplemented)
			roles.Delete("/{role_name}/scope/{scope_type}", s.notImplemented)
			roles.Get("/{role_name}/scope/{scope_name}/permissions", s.notImplemented)
			roles.Put("/{role_name}/scope/{scope_name}/permissions", s.notImplemented)
		})

		// Cursors
		r.Route("/cursors", func(cursors chi.Router) {
			cursors.Get("/0/rooms/{room_id}/users/{user_id}", s.notImplemented)
			cursors.Put("/0/rooms/{room_id}/users/{user_id}", s.notImplemented)
			cursors.Get("/0/rooms/{room_id}", s.notImplemented)
			cursors.Get("/0/users/{user_id}", s.notImplemented)
			cursors.MethodFunc("SUBSCRIBE", "/0/users/{user_id}", s.notImplemented)
			cursors.MethodFunc("SUBSCRIBE", "/0/rooms/{room_id}", s.notImplemented)
		})

		// Token
		r.Post("/token", s.notImplemented)
	})

	s.serv = &http.Server{
		Addr:    s.cfg.Listen,
		Handler: api,
	}
}

func (s *Server) Run(ctx context.Context) error {
	if err := s.serv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Close() error {
	return s.serv.Close()
}

func (s *Server) notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
