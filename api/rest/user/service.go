package api

import (
	"log/slog"
	"net/http"

	// "strconv"
	// "time"

	db "github.com/bhushan217/go-rnd/db/sqlc"
	"github.com/bhushan217/go-rnd/server"
)

// NewService is a constructor of a interface { func RegisterHandlers(*http.ServeMux) } implementation.
// Use this function to customize the server by adding middlewares to it.
func NewService(querier *db.Queries) *UserService {
	return &UserService{querier: querier}
}

type UserService struct {
	querier *db.Queries
}

type UserResponse struct {
	UserID   int64  `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Email    string `json:"email,omitempty"`
	Version  int64  `json:"version,omitempty"`
}

func (s *UserService) handleRegistration() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		newUser, err := server.Decode[db.CreateUserParams](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		result, err := s.querier.CreateUser(r.Context(), newUser)
		if err != nil {
			slog.Error("sql call failed", "error", err, "method", "User")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res := UserResponse{
			UserID:   result.ID,
			Username: result.Username,
			FullName: result.FullName,
			Email:    result.Email,
			Version:  result.Version.Int64,
		}
		server.Encode(w, r, http.StatusOK, res)
	}
}

func (s *UserService) handleGetUser() http.HandlerFunc {
	// type request struct {
	// 	UserID int32 `json:"user_id"`
	// }

	return func(w http.ResponseWriter, r *http.Request) {
		username := r.PathValue("username")

		result, err := s.querier.GetUser(r.Context(), username)
		if err != nil {
			slog.Error("sql call failed", "error", err, "method", "GetUser")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var res UserResponse
		res.UserID = result.ID
		res.Username = result.Username
		res.FullName = result.FullName
		res.Email = result.Email
		res.Version = result.Version.Int64
		server.Encode(w, r, http.StatusOK, res)
	}
}

func (s *UserService) handleGetAllUser() http.HandlerFunc {
	// type request struct {
	// 	UserID int32 `json:"user_id"`
	// }

	return func(w http.ResponseWriter, r *http.Request) {

		searchParams, err := server.Decode[db.ListUsersParams](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		result, err := s.querier.ListUsers(r.Context(), searchParams)
		if err != nil {
			slog.Error("sql call failed", "error", err, "method", "GetAllUser")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		list := make([]UserResponse,0)
		for _, user := range result {
			var res UserResponse
			res.UserID = user.ID
			res.Username = user.Username
			res.FullName = user.FullName
			res.Email = user.Email
			res.Version = user.Version.Int64
			list = append(list, res)
		}
		server.Encode(w, r, http.StatusOK, list)
	}
}

/*
func (s *UserService) handleDeleteBook() http.HandlerFunc {
	type request struct {
		BookID int32 `json:"book_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		if str := r.PathValue("book_id"); str != "" {
			if v, err := strconv.ParseInt(str, 10, 32); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			} else {
				req.BookID = int32(v)
			}
		}
		bookID := req.BookID

		err := s.querier.DeleteBook(r.Context(), bookID)
		if err != nil {
			slog.Error("sql call failed", "error", err, "method", "DeleteBook")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}



func (s *UserService) handleUpdateBook() http.HandlerFunc {
	type request struct {
		Title    string   `json:"title"`
		Tags     []string `json:"tags"`
		BookType string   `json:"book_type"`
		BookID   int32    `json:"book_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req, err := server.Decode[request](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		var arg UpdateBookParams
		arg.Title = req.Title
		arg.Tags = req.Tags
		arg.BookType = BookType(req.BookType)
		arg.BookID = req.BookID

		err = s.querier.UpdateBook(r.Context(), arg)
		if err != nil {
			slog.Error("sql call failed", "error", err, "method", "UpdateBook")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
*/
