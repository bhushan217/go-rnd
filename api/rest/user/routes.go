package api

import "net/http"

func (s *UserService) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("POST /register", s.handleRegistration())
	mux.HandleFunc("GET /list", s.handleGetAllUser())
	mux.HandleFunc("GET /find/{username}", s.handleGetUser())
}